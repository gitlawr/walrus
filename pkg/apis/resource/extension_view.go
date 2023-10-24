package resource

import (
	"errors"
	"fmt"

	"github.com/seal-io/walrus/pkg/apis/runtime"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/resource"
	"github.com/seal-io/walrus/pkg/dao/model/resourcerevision"
	"github.com/seal-io/walrus/pkg/dao/model/templateversion"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
)

type AccessEndpoint struct {
	// Name is identifier for the endpoint.
	Name string `json:"name,omitempty"`
	// Endpoint is access endpoint.
	Endpoints []string `json:"endpoints,omitempty"`
}

type RouteUpgradeRequest struct {
	_ struct{} `route:"PUT=/upgrade"`

	model.ResourceUpdateInput `path:",inline" json:",inline"`
}

func (r *RouteUpgradeRequest) Validate() error {
	if err := r.ResourceUpdateInput.Validate(); err != nil {
		return err
	}

	entity, err := r.Client.Resources().Query().
		Where(resource.ID(r.ID)).
		Select(resource.FieldTemplateID).
		WithTemplate(func(tvq *model.TemplateVersionQuery) {
			tvq.Select(
				templateversion.FieldName,
				templateversion.FieldVersion)
		}).
		Only(r.Context)
	if err != nil {
		return fmt.Errorf("failed to get service: %w", err)
	}

	if r.Template.Name != entity.Edges.Template.Name {
		return errors.New("invalid template name: immutable")
	}

	tv, err := r.Client.TemplateVersions().Query().
		Where(templateversion.ID(r.Template.ID)).
		Select(
			templateversion.FieldSchema,
			templateversion.FieldUiSchema,
		).
		Only(r.Context)
	if err != nil {
		return fmt.Errorf("failed to get template version: %w", err)
	}

	// Verify attributes with variables schema of the template version.
	if !tv.Schema.IsEmpty() {
		if err = r.Attributes.ValidateWith(tv.Schema.VariableSchemas()); err != nil {
			return fmt.Errorf("invalid variables: %w", err)
		}
	}

	// Verify that variables in attributes are valid.
	err = validateVariable(r.Context, r.Client, r.Attributes, r.Name, r.Project.ID, r.Environment.ID)
	if err != nil {
		return err
	}

	if err = validateRevisionsStatus(r.Context, r.Client, r.ID); err != nil {
		return err
	}

	return nil
}

type RouteRollbackRequest struct {
	_ struct{} `route:"POST=/rollback"`

	model.ResourceQueryInput `path:",inline"`

	RevisionID object.ID `query:"revisionID"`
}

func (r *RouteRollbackRequest) Validate() error {
	if err := r.ResourceQueryInput.Validate(); err != nil {
		return err
	}

	latestRevision, err := r.Client.ResourceRevisions().Query().
		Where(resourcerevision.ResourceID(r.ID)).
		Order(model.Desc(resourcerevision.FieldCreateTime)).
		Select(resourcerevision.FieldStatus).
		First(r.Context)
	if err != nil && !model.IsNotFound(err) {
		return fmt.Errorf("failed to get the latest revision: %w", err)
	}

	if status.ResourceRevisionStatusReady.IsUnknown(latestRevision) {
		return errors.New("latest revision is running")
	}

	return nil
}

type (
	RouteGetAccessEndpointsRequest struct {
		_ struct{} `route:"GET=/access-endpoints"`

		model.ResourceQueryInput `path:",inline"`

		Stream *runtime.RequestUnidiStream
	}

	RouteGetAccessEndpointsResponse = []AccessEndpoint
)

func (r *RouteGetAccessEndpointsRequest) SetStream(stream runtime.RequestUnidiStream) {
	r.Stream = &stream
}

type (
	RouteGetOutputsRequest struct {
		_ struct{} `route:"GET=/outputs"`

		model.ResourceQueryInput `path:",inline"`

		Stream *runtime.RequestUnidiStream
	}

	RouteGetOutputsResponse = []types.OutputValue
)

func (r *RouteGetOutputsRequest) SetStream(stream runtime.RequestUnidiStream) {
	r.Stream = &stream
}

type (
	// GraphVertexID defines the identifier of the vertex,
	// which uniquely represents an API resource.
	GraphVertexID = types.GraphVertexID
	// GraphVertex defines the vertex of graph.
	GraphVertex = types.GraphVertex
	// GraphEdge defines the edge of graph.
	GraphEdge = types.GraphEdge

	RouteGetGraphRequest struct {
		_ struct{} `route:"GET=/graph"`

		model.ResourceQueryInput `path:",inline"`

		WithoutKeys bool `query:"withoutKeys,omitempty"`
	}

	RouteGetGraphResponse struct {
		Vertices []GraphVertex `json:"vertices"`
		Edges    []GraphEdge   `json:"edges"`
	}
)