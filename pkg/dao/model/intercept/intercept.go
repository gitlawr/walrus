// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "seal". DO NOT EDIT.

package intercept

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/catalog"
	"github.com/seal-io/walrus/pkg/dao/model/connector"
	"github.com/seal-io/walrus/pkg/dao/model/costreport"
	"github.com/seal-io/walrus/pkg/dao/model/distributelock"
	"github.com/seal-io/walrus/pkg/dao/model/environment"
	"github.com/seal-io/walrus/pkg/dao/model/environmentconnectorrelationship"
	"github.com/seal-io/walrus/pkg/dao/model/perspective"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/project"
	"github.com/seal-io/walrus/pkg/dao/model/role"
	"github.com/seal-io/walrus/pkg/dao/model/service"
	"github.com/seal-io/walrus/pkg/dao/model/servicerelationship"
	"github.com/seal-io/walrus/pkg/dao/model/serviceresource"
	"github.com/seal-io/walrus/pkg/dao/model/serviceresourcerelationship"
	"github.com/seal-io/walrus/pkg/dao/model/servicerevision"
	"github.com/seal-io/walrus/pkg/dao/model/setting"
	"github.com/seal-io/walrus/pkg/dao/model/subject"
	"github.com/seal-io/walrus/pkg/dao/model/subjectrolerelationship"
	"github.com/seal-io/walrus/pkg/dao/model/template"
	"github.com/seal-io/walrus/pkg/dao/model/templateversion"
	"github.com/seal-io/walrus/pkg/dao/model/token"
	"github.com/seal-io/walrus/pkg/dao/model/variable"
)

// The Query interface represents an operation that queries a graph.
// By using this interface, users can write generic code that manipulates
// query builders of different types.
type Query interface {
	// Type returns the string representation of the query type.
	Type() string
	// Limit the number of records to be returned by this query.
	Limit(int)
	// Offset to start from.
	Offset(int)
	// Unique configures the query builder to filter duplicate records.
	Unique(bool)
	// Order specifies how the records should be ordered.
	Order(...func(*sql.Selector))
	// WhereP appends storage-level predicates to the query builder. Using this method, users
	// can use type-assertion to append predicates that do not depend on any generated package.
	WhereP(...func(*sql.Selector))
}

// The Func type is an adapter that allows ordinary functions to be used as interceptors.
// Unlike traversal functions, interceptors are skipped during graph traversals. Note that the
// implementation of Func is different from the one defined in entgo.io/ent.InterceptFunc.
type Func func(context.Context, Query) error

// Intercept calls f(ctx, q) and then applied the next Querier.
func (f Func) Intercept(next model.Querier) model.Querier {
	return model.QuerierFunc(func(ctx context.Context, q model.Query) (model.Value, error) {
		query, err := NewQuery(q)
		if err != nil {
			return nil, err
		}
		if err := f(ctx, query); err != nil {
			return nil, err
		}
		return next.Query(ctx, q)
	})
}

// The TraverseFunc type is an adapter to allow the use of ordinary function as Traverser.
// If f is a function with the appropriate signature, TraverseFunc(f) is a Traverser that calls f.
type TraverseFunc func(context.Context, Query) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFunc) Intercept(next model.Querier) model.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFunc) Traverse(ctx context.Context, q model.Query) error {
	query, err := NewQuery(q)
	if err != nil {
		return err
	}
	return f(ctx, query)
}

// The CatalogFunc type is an adapter to allow the use of ordinary function as a Querier.
type CatalogFunc func(context.Context, *model.CatalogQuery) (model.Value, error)

// Query calls f(ctx, q).
func (f CatalogFunc) Query(ctx context.Context, q model.Query) (model.Value, error) {
	if q, ok := q.(*model.CatalogQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *model.CatalogQuery", q)
}

// The TraverseCatalog type is an adapter to allow the use of ordinary function as Traverser.
type TraverseCatalog func(context.Context, *model.CatalogQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseCatalog) Intercept(next model.Querier) model.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseCatalog) Traverse(ctx context.Context, q model.Query) error {
	if q, ok := q.(*model.CatalogQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *model.CatalogQuery", q)
}

// The ConnectorFunc type is an adapter to allow the use of ordinary function as a Querier.
type ConnectorFunc func(context.Context, *model.ConnectorQuery) (model.Value, error)

// Query calls f(ctx, q).
func (f ConnectorFunc) Query(ctx context.Context, q model.Query) (model.Value, error) {
	if q, ok := q.(*model.ConnectorQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *model.ConnectorQuery", q)
}

// The TraverseConnector type is an adapter to allow the use of ordinary function as Traverser.
type TraverseConnector func(context.Context, *model.ConnectorQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseConnector) Intercept(next model.Querier) model.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseConnector) Traverse(ctx context.Context, q model.Query) error {
	if q, ok := q.(*model.ConnectorQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *model.ConnectorQuery", q)
}

// The CostReportFunc type is an adapter to allow the use of ordinary function as a Querier.
type CostReportFunc func(context.Context, *model.CostReportQuery) (model.Value, error)

// Query calls f(ctx, q).
func (f CostReportFunc) Query(ctx context.Context, q model.Query) (model.Value, error) {
	if q, ok := q.(*model.CostReportQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *model.CostReportQuery", q)
}

// The TraverseCostReport type is an adapter to allow the use of ordinary function as Traverser.
type TraverseCostReport func(context.Context, *model.CostReportQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseCostReport) Intercept(next model.Querier) model.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseCostReport) Traverse(ctx context.Context, q model.Query) error {
	if q, ok := q.(*model.CostReportQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *model.CostReportQuery", q)
}

// The DistributeLockFunc type is an adapter to allow the use of ordinary function as a Querier.
type DistributeLockFunc func(context.Context, *model.DistributeLockQuery) (model.Value, error)

// Query calls f(ctx, q).
func (f DistributeLockFunc) Query(ctx context.Context, q model.Query) (model.Value, error) {
	if q, ok := q.(*model.DistributeLockQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *model.DistributeLockQuery", q)
}

// The TraverseDistributeLock type is an adapter to allow the use of ordinary function as Traverser.
type TraverseDistributeLock func(context.Context, *model.DistributeLockQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseDistributeLock) Intercept(next model.Querier) model.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseDistributeLock) Traverse(ctx context.Context, q model.Query) error {
	if q, ok := q.(*model.DistributeLockQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *model.DistributeLockQuery", q)
}

// The EnvironmentFunc type is an adapter to allow the use of ordinary function as a Querier.
type EnvironmentFunc func(context.Context, *model.EnvironmentQuery) (model.Value, error)

// Query calls f(ctx, q).
func (f EnvironmentFunc) Query(ctx context.Context, q model.Query) (model.Value, error) {
	if q, ok := q.(*model.EnvironmentQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *model.EnvironmentQuery", q)
}

// The TraverseEnvironment type is an adapter to allow the use of ordinary function as Traverser.
type TraverseEnvironment func(context.Context, *model.EnvironmentQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseEnvironment) Intercept(next model.Querier) model.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseEnvironment) Traverse(ctx context.Context, q model.Query) error {
	if q, ok := q.(*model.EnvironmentQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *model.EnvironmentQuery", q)
}

// The EnvironmentConnectorRelationshipFunc type is an adapter to allow the use of ordinary function as a Querier.
type EnvironmentConnectorRelationshipFunc func(context.Context, *model.EnvironmentConnectorRelationshipQuery) (model.Value, error)

// Query calls f(ctx, q).
func (f EnvironmentConnectorRelationshipFunc) Query(ctx context.Context, q model.Query) (model.Value, error) {
	if q, ok := q.(*model.EnvironmentConnectorRelationshipQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *model.EnvironmentConnectorRelationshipQuery", q)
}

// The TraverseEnvironmentConnectorRelationship type is an adapter to allow the use of ordinary function as Traverser.
type TraverseEnvironmentConnectorRelationship func(context.Context, *model.EnvironmentConnectorRelationshipQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseEnvironmentConnectorRelationship) Intercept(next model.Querier) model.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseEnvironmentConnectorRelationship) Traverse(ctx context.Context, q model.Query) error {
	if q, ok := q.(*model.EnvironmentConnectorRelationshipQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *model.EnvironmentConnectorRelationshipQuery", q)
}

// The PerspectiveFunc type is an adapter to allow the use of ordinary function as a Querier.
type PerspectiveFunc func(context.Context, *model.PerspectiveQuery) (model.Value, error)

// Query calls f(ctx, q).
func (f PerspectiveFunc) Query(ctx context.Context, q model.Query) (model.Value, error) {
	if q, ok := q.(*model.PerspectiveQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *model.PerspectiveQuery", q)
}

// The TraversePerspective type is an adapter to allow the use of ordinary function as Traverser.
type TraversePerspective func(context.Context, *model.PerspectiveQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraversePerspective) Intercept(next model.Querier) model.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraversePerspective) Traverse(ctx context.Context, q model.Query) error {
	if q, ok := q.(*model.PerspectiveQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *model.PerspectiveQuery", q)
}

// The ProjectFunc type is an adapter to allow the use of ordinary function as a Querier.
type ProjectFunc func(context.Context, *model.ProjectQuery) (model.Value, error)

// Query calls f(ctx, q).
func (f ProjectFunc) Query(ctx context.Context, q model.Query) (model.Value, error) {
	if q, ok := q.(*model.ProjectQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *model.ProjectQuery", q)
}

// The TraverseProject type is an adapter to allow the use of ordinary function as Traverser.
type TraverseProject func(context.Context, *model.ProjectQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseProject) Intercept(next model.Querier) model.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseProject) Traverse(ctx context.Context, q model.Query) error {
	if q, ok := q.(*model.ProjectQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *model.ProjectQuery", q)
}

// The RoleFunc type is an adapter to allow the use of ordinary function as a Querier.
type RoleFunc func(context.Context, *model.RoleQuery) (model.Value, error)

// Query calls f(ctx, q).
func (f RoleFunc) Query(ctx context.Context, q model.Query) (model.Value, error) {
	if q, ok := q.(*model.RoleQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *model.RoleQuery", q)
}

// The TraverseRole type is an adapter to allow the use of ordinary function as Traverser.
type TraverseRole func(context.Context, *model.RoleQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseRole) Intercept(next model.Querier) model.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseRole) Traverse(ctx context.Context, q model.Query) error {
	if q, ok := q.(*model.RoleQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *model.RoleQuery", q)
}

// The ServiceFunc type is an adapter to allow the use of ordinary function as a Querier.
type ServiceFunc func(context.Context, *model.ServiceQuery) (model.Value, error)

// Query calls f(ctx, q).
func (f ServiceFunc) Query(ctx context.Context, q model.Query) (model.Value, error) {
	if q, ok := q.(*model.ServiceQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *model.ServiceQuery", q)
}

// The TraverseService type is an adapter to allow the use of ordinary function as Traverser.
type TraverseService func(context.Context, *model.ServiceQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseService) Intercept(next model.Querier) model.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseService) Traverse(ctx context.Context, q model.Query) error {
	if q, ok := q.(*model.ServiceQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *model.ServiceQuery", q)
}

// The ServiceRelationshipFunc type is an adapter to allow the use of ordinary function as a Querier.
type ServiceRelationshipFunc func(context.Context, *model.ServiceRelationshipQuery) (model.Value, error)

// Query calls f(ctx, q).
func (f ServiceRelationshipFunc) Query(ctx context.Context, q model.Query) (model.Value, error) {
	if q, ok := q.(*model.ServiceRelationshipQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *model.ServiceRelationshipQuery", q)
}

// The TraverseServiceRelationship type is an adapter to allow the use of ordinary function as Traverser.
type TraverseServiceRelationship func(context.Context, *model.ServiceRelationshipQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseServiceRelationship) Intercept(next model.Querier) model.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseServiceRelationship) Traverse(ctx context.Context, q model.Query) error {
	if q, ok := q.(*model.ServiceRelationshipQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *model.ServiceRelationshipQuery", q)
}

// The ServiceResourceFunc type is an adapter to allow the use of ordinary function as a Querier.
type ServiceResourceFunc func(context.Context, *model.ServiceResourceQuery) (model.Value, error)

// Query calls f(ctx, q).
func (f ServiceResourceFunc) Query(ctx context.Context, q model.Query) (model.Value, error) {
	if q, ok := q.(*model.ServiceResourceQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *model.ServiceResourceQuery", q)
}

// The TraverseServiceResource type is an adapter to allow the use of ordinary function as Traverser.
type TraverseServiceResource func(context.Context, *model.ServiceResourceQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseServiceResource) Intercept(next model.Querier) model.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseServiceResource) Traverse(ctx context.Context, q model.Query) error {
	if q, ok := q.(*model.ServiceResourceQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *model.ServiceResourceQuery", q)
}

// The ServiceResourceRelationshipFunc type is an adapter to allow the use of ordinary function as a Querier.
type ServiceResourceRelationshipFunc func(context.Context, *model.ServiceResourceRelationshipQuery) (model.Value, error)

// Query calls f(ctx, q).
func (f ServiceResourceRelationshipFunc) Query(ctx context.Context, q model.Query) (model.Value, error) {
	if q, ok := q.(*model.ServiceResourceRelationshipQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *model.ServiceResourceRelationshipQuery", q)
}

// The TraverseServiceResourceRelationship type is an adapter to allow the use of ordinary function as Traverser.
type TraverseServiceResourceRelationship func(context.Context, *model.ServiceResourceRelationshipQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseServiceResourceRelationship) Intercept(next model.Querier) model.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseServiceResourceRelationship) Traverse(ctx context.Context, q model.Query) error {
	if q, ok := q.(*model.ServiceResourceRelationshipQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *model.ServiceResourceRelationshipQuery", q)
}

// The ServiceRevisionFunc type is an adapter to allow the use of ordinary function as a Querier.
type ServiceRevisionFunc func(context.Context, *model.ServiceRevisionQuery) (model.Value, error)

// Query calls f(ctx, q).
func (f ServiceRevisionFunc) Query(ctx context.Context, q model.Query) (model.Value, error) {
	if q, ok := q.(*model.ServiceRevisionQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *model.ServiceRevisionQuery", q)
}

// The TraverseServiceRevision type is an adapter to allow the use of ordinary function as Traverser.
type TraverseServiceRevision func(context.Context, *model.ServiceRevisionQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseServiceRevision) Intercept(next model.Querier) model.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseServiceRevision) Traverse(ctx context.Context, q model.Query) error {
	if q, ok := q.(*model.ServiceRevisionQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *model.ServiceRevisionQuery", q)
}

// The SettingFunc type is an adapter to allow the use of ordinary function as a Querier.
type SettingFunc func(context.Context, *model.SettingQuery) (model.Value, error)

// Query calls f(ctx, q).
func (f SettingFunc) Query(ctx context.Context, q model.Query) (model.Value, error) {
	if q, ok := q.(*model.SettingQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *model.SettingQuery", q)
}

// The TraverseSetting type is an adapter to allow the use of ordinary function as Traverser.
type TraverseSetting func(context.Context, *model.SettingQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseSetting) Intercept(next model.Querier) model.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseSetting) Traverse(ctx context.Context, q model.Query) error {
	if q, ok := q.(*model.SettingQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *model.SettingQuery", q)
}

// The SubjectFunc type is an adapter to allow the use of ordinary function as a Querier.
type SubjectFunc func(context.Context, *model.SubjectQuery) (model.Value, error)

// Query calls f(ctx, q).
func (f SubjectFunc) Query(ctx context.Context, q model.Query) (model.Value, error) {
	if q, ok := q.(*model.SubjectQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *model.SubjectQuery", q)
}

// The TraverseSubject type is an adapter to allow the use of ordinary function as Traverser.
type TraverseSubject func(context.Context, *model.SubjectQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseSubject) Intercept(next model.Querier) model.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseSubject) Traverse(ctx context.Context, q model.Query) error {
	if q, ok := q.(*model.SubjectQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *model.SubjectQuery", q)
}

// The SubjectRoleRelationshipFunc type is an adapter to allow the use of ordinary function as a Querier.
type SubjectRoleRelationshipFunc func(context.Context, *model.SubjectRoleRelationshipQuery) (model.Value, error)

// Query calls f(ctx, q).
func (f SubjectRoleRelationshipFunc) Query(ctx context.Context, q model.Query) (model.Value, error) {
	if q, ok := q.(*model.SubjectRoleRelationshipQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *model.SubjectRoleRelationshipQuery", q)
}

// The TraverseSubjectRoleRelationship type is an adapter to allow the use of ordinary function as Traverser.
type TraverseSubjectRoleRelationship func(context.Context, *model.SubjectRoleRelationshipQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseSubjectRoleRelationship) Intercept(next model.Querier) model.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseSubjectRoleRelationship) Traverse(ctx context.Context, q model.Query) error {
	if q, ok := q.(*model.SubjectRoleRelationshipQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *model.SubjectRoleRelationshipQuery", q)
}

// The TemplateFunc type is an adapter to allow the use of ordinary function as a Querier.
type TemplateFunc func(context.Context, *model.TemplateQuery) (model.Value, error)

// Query calls f(ctx, q).
func (f TemplateFunc) Query(ctx context.Context, q model.Query) (model.Value, error) {
	if q, ok := q.(*model.TemplateQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *model.TemplateQuery", q)
}

// The TraverseTemplate type is an adapter to allow the use of ordinary function as Traverser.
type TraverseTemplate func(context.Context, *model.TemplateQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseTemplate) Intercept(next model.Querier) model.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseTemplate) Traverse(ctx context.Context, q model.Query) error {
	if q, ok := q.(*model.TemplateQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *model.TemplateQuery", q)
}

// The TemplateVersionFunc type is an adapter to allow the use of ordinary function as a Querier.
type TemplateVersionFunc func(context.Context, *model.TemplateVersionQuery) (model.Value, error)

// Query calls f(ctx, q).
func (f TemplateVersionFunc) Query(ctx context.Context, q model.Query) (model.Value, error) {
	if q, ok := q.(*model.TemplateVersionQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *model.TemplateVersionQuery", q)
}

// The TraverseTemplateVersion type is an adapter to allow the use of ordinary function as Traverser.
type TraverseTemplateVersion func(context.Context, *model.TemplateVersionQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseTemplateVersion) Intercept(next model.Querier) model.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseTemplateVersion) Traverse(ctx context.Context, q model.Query) error {
	if q, ok := q.(*model.TemplateVersionQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *model.TemplateVersionQuery", q)
}

// The TokenFunc type is an adapter to allow the use of ordinary function as a Querier.
type TokenFunc func(context.Context, *model.TokenQuery) (model.Value, error)

// Query calls f(ctx, q).
func (f TokenFunc) Query(ctx context.Context, q model.Query) (model.Value, error) {
	if q, ok := q.(*model.TokenQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *model.TokenQuery", q)
}

// The TraverseToken type is an adapter to allow the use of ordinary function as Traverser.
type TraverseToken func(context.Context, *model.TokenQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseToken) Intercept(next model.Querier) model.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseToken) Traverse(ctx context.Context, q model.Query) error {
	if q, ok := q.(*model.TokenQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *model.TokenQuery", q)
}

// The VariableFunc type is an adapter to allow the use of ordinary function as a Querier.
type VariableFunc func(context.Context, *model.VariableQuery) (model.Value, error)

// Query calls f(ctx, q).
func (f VariableFunc) Query(ctx context.Context, q model.Query) (model.Value, error) {
	if q, ok := q.(*model.VariableQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *model.VariableQuery", q)
}

// The TraverseVariable type is an adapter to allow the use of ordinary function as Traverser.
type TraverseVariable func(context.Context, *model.VariableQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseVariable) Intercept(next model.Querier) model.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseVariable) Traverse(ctx context.Context, q model.Query) error {
	if q, ok := q.(*model.VariableQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *model.VariableQuery", q)
}

// NewQuery returns the generic Query interface for the given typed query.
func NewQuery(q model.Query) (Query, error) {
	switch q := q.(type) {
	case *model.CatalogQuery:
		return &query[*model.CatalogQuery, predicate.Catalog, catalog.OrderOption]{typ: model.TypeCatalog, tq: q}, nil
	case *model.ConnectorQuery:
		return &query[*model.ConnectorQuery, predicate.Connector, connector.OrderOption]{typ: model.TypeConnector, tq: q}, nil
	case *model.CostReportQuery:
		return &query[*model.CostReportQuery, predicate.CostReport, costreport.OrderOption]{typ: model.TypeCostReport, tq: q}, nil
	case *model.DistributeLockQuery:
		return &query[*model.DistributeLockQuery, predicate.DistributeLock, distributelock.OrderOption]{typ: model.TypeDistributeLock, tq: q}, nil
	case *model.EnvironmentQuery:
		return &query[*model.EnvironmentQuery, predicate.Environment, environment.OrderOption]{typ: model.TypeEnvironment, tq: q}, nil
	case *model.EnvironmentConnectorRelationshipQuery:
		return &query[*model.EnvironmentConnectorRelationshipQuery, predicate.EnvironmentConnectorRelationship, environmentconnectorrelationship.OrderOption]{typ: model.TypeEnvironmentConnectorRelationship, tq: q}, nil
	case *model.PerspectiveQuery:
		return &query[*model.PerspectiveQuery, predicate.Perspective, perspective.OrderOption]{typ: model.TypePerspective, tq: q}, nil
	case *model.ProjectQuery:
		return &query[*model.ProjectQuery, predicate.Project, project.OrderOption]{typ: model.TypeProject, tq: q}, nil
	case *model.RoleQuery:
		return &query[*model.RoleQuery, predicate.Role, role.OrderOption]{typ: model.TypeRole, tq: q}, nil
	case *model.ServiceQuery:
		return &query[*model.ServiceQuery, predicate.Service, service.OrderOption]{typ: model.TypeService, tq: q}, nil
	case *model.ServiceRelationshipQuery:
		return &query[*model.ServiceRelationshipQuery, predicate.ServiceRelationship, servicerelationship.OrderOption]{typ: model.TypeServiceRelationship, tq: q}, nil
	case *model.ServiceResourceQuery:
		return &query[*model.ServiceResourceQuery, predicate.ServiceResource, serviceresource.OrderOption]{typ: model.TypeServiceResource, tq: q}, nil
	case *model.ServiceResourceRelationshipQuery:
		return &query[*model.ServiceResourceRelationshipQuery, predicate.ServiceResourceRelationship, serviceresourcerelationship.OrderOption]{typ: model.TypeServiceResourceRelationship, tq: q}, nil
	case *model.ServiceRevisionQuery:
		return &query[*model.ServiceRevisionQuery, predicate.ServiceRevision, servicerevision.OrderOption]{typ: model.TypeServiceRevision, tq: q}, nil
	case *model.SettingQuery:
		return &query[*model.SettingQuery, predicate.Setting, setting.OrderOption]{typ: model.TypeSetting, tq: q}, nil
	case *model.SubjectQuery:
		return &query[*model.SubjectQuery, predicate.Subject, subject.OrderOption]{typ: model.TypeSubject, tq: q}, nil
	case *model.SubjectRoleRelationshipQuery:
		return &query[*model.SubjectRoleRelationshipQuery, predicate.SubjectRoleRelationship, subjectrolerelationship.OrderOption]{typ: model.TypeSubjectRoleRelationship, tq: q}, nil
	case *model.TemplateQuery:
		return &query[*model.TemplateQuery, predicate.Template, template.OrderOption]{typ: model.TypeTemplate, tq: q}, nil
	case *model.TemplateVersionQuery:
		return &query[*model.TemplateVersionQuery, predicate.TemplateVersion, templateversion.OrderOption]{typ: model.TypeTemplateVersion, tq: q}, nil
	case *model.TokenQuery:
		return &query[*model.TokenQuery, predicate.Token, token.OrderOption]{typ: model.TypeToken, tq: q}, nil
	case *model.VariableQuery:
		return &query[*model.VariableQuery, predicate.Variable, variable.OrderOption]{typ: model.TypeVariable, tq: q}, nil
	default:
		return nil, fmt.Errorf("unknown query type %T", q)
	}
}

type query[T any, P ~func(*sql.Selector), R ~func(*sql.Selector)] struct {
	typ string
	tq  interface {
		Limit(int) T
		Offset(int) T
		Unique(bool) T
		Order(...R) T
		Where(...P) T
	}
}

func (q query[T, P, R]) Type() string {
	return q.typ
}

func (q query[T, P, R]) Limit(limit int) {
	q.tq.Limit(limit)
}

func (q query[T, P, R]) Offset(offset int) {
	q.tq.Offset(offset)
}

func (q query[T, P, R]) Unique(unique bool) {
	q.tq.Unique(unique)
}

func (q query[T, P, R]) Order(orders ...func(*sql.Selector)) {
	rs := make([]R, len(orders))
	for i := range orders {
		rs[i] = orders[i]
	}
	q.tq.Order(rs...)
}

func (q query[T, P, R]) WhereP(ps ...func(*sql.Selector)) {
	p := make([]P, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	q.tq.Where(p...)
}
