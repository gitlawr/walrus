package server

import (
	"context"
	modbus "github.com/seal-io/walrus/pkg/bus/template"
	"github.com/seal-io/walrus/pkg/dao/types/property"
	"k8s.io/apimachinery/pkg/util/wait"
	"time"

	types2 "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"github.com/seal-io/walrus/pkg/dao"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/connector"
	"github.com/seal-io/walrus/pkg/dao/model/environment"
	"github.com/seal-io/walrus/pkg/dao/model/project"
	"github.com/seal-io/walrus/pkg/dao/model/resourcedefinition"
	"github.com/seal-io/walrus/pkg/dao/model/template"
	"github.com/seal-io/walrus/pkg/dao/model/templateversion"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/crypto"
	"github.com/seal-io/walrus/pkg/settings"
	"github.com/seal-io/walrus/utils/log"
)

func (r *Server) createLocalEnvironment(ctx context.Context, opts initOptions) error {
	return opts.ModelClient.WithTx(ctx, func(tx *model.Tx) error {
		enableLocalMode, err := settings.EnableLocalMode.ValueBool(ctx, tx)
		if err != nil {
			return err
		}

		if !enableLocalMode {
			return nil
		}

		defaultProject, err := opts.ModelClient.Projects().Query().
			Where(project.Name("default")).
			Only(ctx)
		if err != nil {
			return err
		}

		conn, err := tx.Connectors().Query().
			Where(connector.Name("docker")).
			Only(ctx)
		if model.IsNotFound(err) {
			conn = &model.Connector{
				Name:                      "docker",
				Type:                      "docker",
				ProjectID:                 defaultProject.ID,
				ApplicableEnvironmentType: types.EnvironmentDevelopment,
				// TODO rename ConnectorCategoryKubernetes to ConnectorCategoryContainer.
				Category:      types.ConnectorCategoryKubernetes,
				ConfigVersion: "v1",
				ConfigData:    crypto.Properties{},
			}

			conn, err = tx.Connectors().Create().
				Set(conn).
				Save(ctx)
			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		}

		env := &model.Environment{
			Name:      "local",
			ProjectID: defaultProject.ID,
			Type:      types.EnvironmentDevelopment,
			Edges: model.EnvironmentEdges{
				Connectors: []*model.EnvironmentConnectorRelationship{
					{
						ConnectorID: conn.ID,
					},
				},
			},
		}

		err = tx.Environments().Create().
			Set(env).
			OnConflictColumns(environment.FieldProjectID, environment.FieldName).
			Ignore().
			ExecE(ctx, dao.EnvironmentConnectorsEdgeSave)

		return err
	})
}

func (r *Server) createBuiltinDefinitions(ctx context.Context, opts initOptions) error {
	go func() {
		err := r.doCreateBuiltinDefinitions(ctx, opts)
		if err != nil {
			log.Error("failed to create builtin definitions, error: %v", err)
		}
	}()
	return nil
}

func (r *Server) doCreateBuiltinDefinitions(ctx context.Context, opts initOptions) error {
	enableLocalMode, err := settings.EnableLocalMode.ValueBool(ctx, opts.ModelClient)
	if err != nil {
		return err
	}

	if !enableLocalMode {
		return nil
	}

	rd, err := opts.ModelClient.ResourceDefinitions().Query().
		Where(resourcedefinition.Name("containerservice")).
		Only(ctx)
	if err != nil && !model.IsNotFound(err) {
		return err
	} else if err == nil {
		return nil
	}

	// create container service resource definition.

	tmpl, err := opts.ModelClient.Templates().Query().
		Where(template.Name("docker-containerservice")).
		Only(ctx)
	if model.IsNotFound(err) {
		tmpl = &model.Template{
			Name:   "docker-containerservice",
			Source: "https://github.com/walrus-catalog-sandbox/terraform-docker-containerservice?ref=main",
		}
		tmpl, err = opts.ModelClient.Templates().Create().
			Set(tmpl).
			Save(ctx)
		if err != nil {
			return err
		}
		if err = modbus.Notify(ctx, tmpl); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	var ktv, dtv *model.TemplateVersion
	err = wait.PollUntilContextTimeout(ctx, 15*time.Second, 5*time.Minute, false,
		func(ctx context.Context) (bool, error) {
			ktv, err = opts.ModelClient.TemplateVersions().Query().
				Where(
					templateversion.Name("kubernetes-containerservice"),
					templateversion.Version("v0.1.3"),
				).
				Only(ctx)
			if model.IsNotFound(err) {
				return false, nil
			} else if err != nil {
				return false, err
			}

			dtv, err = opts.ModelClient.TemplateVersions().Query().
				Where(
					templateversion.Name("docker-containerservice"),
					templateversion.Version("main"),
				).
				Only(ctx)
			if model.IsNotFound(err) {
				return false, nil
			} else if err != nil {
				return false, err
			}

			return true, nil
		})
	if err != nil {
		return err
	}

	rd = &model.ResourceDefinition{
		Name: "containerservice",
		Type: "containerservice",
		Edges: model.ResourceDefinitionEdges{
			MatchingRules: []*model.ResourceDefinitionMatchingRule{
				{
					Name:       "local",
					TemplateID: dtv.ID,
					Selector: types.Selector{
						EnvironmentName: "local",
					},
					Attributes: property.Values{
						"infrastructure": []byte(`{"network_id": "walrus-local"}`),
					},
				}, {
					Name:       "development",
					TemplateID: ktv.ID,
					Selector: types.Selector{
						EnvironmentType: types.EnvironmentDevelopment,
					},
				},
			},
		},
	}
	_, err = opts.ModelClient.ResourceDefinitions().Create().
		Set(rd).
		SaveE(ctx, dao.ResourceDefinitionMatchingRulesEdgeSave)
	return err
}

func (r *Server) createLocalDockerNetwork(ctx context.Context, opts initOptions) error {
	enableLocalMode, err := settings.EnableLocalMode.ValueBool(ctx, opts.ModelClient)
	if err != nil {
		return err
	}

	if !enableLocalMode {
		return nil
	}

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	networkName := "walrus-local"
	networks, err := cli.NetworkList(ctx, types2.NetworkListOptions{})
	if err != nil {
		return err
	}

	exists := false
	for _, n := range networks {
		if n.Name == networkName {
			exists = true
			break
		}
	}
	if !exists {
		_, err = cli.NetworkCreate(ctx, networkName, types2.NetworkCreate{
			Driver: "bridge",
		})
		if err != nil {
			return err
		}
		log.Debug("walrus-local docker network created")
	}
	return nil
}
