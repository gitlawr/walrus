// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "seal". DO NOT EDIT.

package environment

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"golang.org/x/exp/slices"
)

const (
	// Label holds the string label denoting the environment type in the database.
	Label = "environment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldLabels holds the string denoting the labels field in the database.
	FieldLabels = "labels"
	// FieldAnnotations holds the string denoting the annotations field in the database.
	FieldAnnotations = "annotations"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldProjectID holds the string denoting the project_id field in the database.
	FieldProjectID = "project_id"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// EdgeConnectors holds the string denoting the connectors edge name in mutations.
	EdgeConnectors = "connectors"
	// EdgeServices holds the string denoting the services edge name in mutations.
	EdgeServices = "services"
	// EdgeServiceRevisions holds the string denoting the service_revisions edge name in mutations.
	EdgeServiceRevisions = "service_revisions"
	// EdgeServiceResources holds the string denoting the service_resources edge name in mutations.
	EdgeServiceResources = "service_resources"
	// EdgeVariables holds the string denoting the variables edge name in mutations.
	EdgeVariables = "variables"
	// Table holds the table name of the environment in the database.
	Table = "environments"
	// ProjectTable is the table that holds the project relation/edge.
	ProjectTable = "environments"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "project_id"
	// ConnectorsTable is the table that holds the connectors relation/edge.
	ConnectorsTable = "environment_connector_relationships"
	// ConnectorsInverseTable is the table name for the EnvironmentConnectorRelationship entity.
	// It exists in this package in order to avoid circular dependency with the "environmentconnectorrelationship" package.
	ConnectorsInverseTable = "environment_connector_relationships"
	// ConnectorsColumn is the table column denoting the connectors relation/edge.
	ConnectorsColumn = "environment_id"
	// ServicesTable is the table that holds the services relation/edge.
	ServicesTable = "services"
	// ServicesInverseTable is the table name for the Service entity.
	// It exists in this package in order to avoid circular dependency with the "service" package.
	ServicesInverseTable = "services"
	// ServicesColumn is the table column denoting the services relation/edge.
	ServicesColumn = "environment_id"
	// ServiceRevisionsTable is the table that holds the service_revisions relation/edge.
	ServiceRevisionsTable = "service_revisions"
	// ServiceRevisionsInverseTable is the table name for the ServiceRevision entity.
	// It exists in this package in order to avoid circular dependency with the "servicerevision" package.
	ServiceRevisionsInverseTable = "service_revisions"
	// ServiceRevisionsColumn is the table column denoting the service_revisions relation/edge.
	ServiceRevisionsColumn = "environment_id"
	// ServiceResourcesTable is the table that holds the service_resources relation/edge.
	ServiceResourcesTable = "service_resources"
	// ServiceResourcesInverseTable is the table name for the ServiceResource entity.
	// It exists in this package in order to avoid circular dependency with the "serviceresource" package.
	ServiceResourcesInverseTable = "service_resources"
	// ServiceResourcesColumn is the table column denoting the service_resources relation/edge.
	ServiceResourcesColumn = "environment_id"
	// VariablesTable is the table that holds the variables relation/edge.
	VariablesTable = "variables"
	// VariablesInverseTable is the table name for the Variable entity.
	// It exists in this package in order to avoid circular dependency with the "variable" package.
	VariablesInverseTable = "variables"
	// VariablesColumn is the table column denoting the variables relation/edge.
	VariablesColumn = "environment_id"
)

// Columns holds all SQL columns for environment fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldLabels,
	FieldAnnotations,
	FieldCreateTime,
	FieldUpdateTime,
	FieldProjectID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/seal-io/walrus/pkg/dao/model/runtime"
var (
	Hooks        [1]ent.Hook
	Interceptors [1]ent.Interceptor
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultLabels holds the default value on creation for the "labels" field.
	DefaultLabels map[string]string
	// DefaultAnnotations holds the default value on creation for the "annotations" field.
	DefaultAnnotations map[string]string
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// ProjectIDValidator is a validator for the "project_id" field. It is called by the builders before save.
	ProjectIDValidator func(string) error
)

// OrderOption defines the ordering options for the Environment queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByProjectID orders the results by the project_id field.
func ByProjectID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProjectID, opts...).ToFunc()
}

// ByProjectField orders the results by project field.
func ByProjectField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProjectStep(), sql.OrderByField(field, opts...))
	}
}

// ByConnectorsCount orders the results by connectors count.
func ByConnectorsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newConnectorsStep(), opts...)
	}
}

// ByConnectors orders the results by connectors terms.
func ByConnectors(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newConnectorsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByServicesCount orders the results by services count.
func ByServicesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newServicesStep(), opts...)
	}
}

// ByServices orders the results by services terms.
func ByServices(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newServicesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByServiceRevisionsCount orders the results by service_revisions count.
func ByServiceRevisionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newServiceRevisionsStep(), opts...)
	}
}

// ByServiceRevisions orders the results by service_revisions terms.
func ByServiceRevisions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newServiceRevisionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByServiceResourcesCount orders the results by service_resources count.
func ByServiceResourcesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newServiceResourcesStep(), opts...)
	}
}

// ByServiceResources orders the results by service_resources terms.
func ByServiceResources(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newServiceResourcesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByVariablesCount orders the results by variables count.
func ByVariablesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newVariablesStep(), opts...)
	}
}

// ByVariables orders the results by variables terms.
func ByVariables(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVariablesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProjectStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProjectInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
	)
}
func newConnectorsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ConnectorsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, ConnectorsTable, ConnectorsColumn),
	)
}
func newServicesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ServicesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ServicesTable, ServicesColumn),
	)
}
func newServiceRevisionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ServiceRevisionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ServiceRevisionsTable, ServiceRevisionsColumn),
	)
}
func newServiceResourcesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ServiceResourcesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ServiceResourcesTable, ServiceResourcesColumn),
	)
}
func newVariablesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VariablesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, VariablesTable, VariablesColumn),
	)
}

// WithoutFields returns the fields ignored the given list.
func WithoutFields(ignores ...string) []string {
	if len(ignores) == 0 {
		return slices.Clone(Columns)
	}

	var s = make(map[string]bool, len(ignores))
	for i := range ignores {
		s[ignores[i]] = true
	}

	var r = make([]string, 0, len(Columns)-len(s))
	for i := range Columns {
		if s[Columns[i]] {
			continue
		}
		r = append(r, Columns[i])
	}
	return r
}
