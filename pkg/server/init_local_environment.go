package server

import (
	"context"
	"time"

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
		// Simple hack to wait template version ready.
		time.Sleep(time.Second * 10)
	} else if err != nil {
		return err
	}

	ktv, err := opts.ModelClient.TemplateVersions().Query().
		Where(
			templateversion.Name("kubernetes-containerservice"),
			templateversion.Version("v0.1.3"),
		).
		Only(ctx)
	if err != nil {
		return err
	}

	dtv, err := opts.ModelClient.TemplateVersions().Query().
		Where(
			templateversion.Name("docker-containerservice"),
			templateversion.Version("main"),
		).
		Only(ctx)
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
