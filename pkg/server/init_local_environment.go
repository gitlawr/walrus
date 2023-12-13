package server

import (
	"context"

	"github.com/seal-io/walrus/pkg/dao"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/connector"
	"github.com/seal-io/walrus/pkg/dao/model/environment"
	"github.com/seal-io/walrus/pkg/dao/model/project"
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
