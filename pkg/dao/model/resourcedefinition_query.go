// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/resource"
	"github.com/seal-io/walrus/pkg/dao/model/resourcedefinition"
	"github.com/seal-io/walrus/pkg/dao/model/resourcedefinitionmatchingrule"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// ResourceDefinitionQuery is the builder for querying ResourceDefinition entities.
type ResourceDefinitionQuery struct {
	config
	ctx               *QueryContext
	order             []resourcedefinition.OrderOption
	inters            []Interceptor
	predicates        []predicate.ResourceDefinition
	withMatchingRules *ResourceDefinitionMatchingRuleQuery
	withResources     *ResourceQuery
	modifiers         []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ResourceDefinitionQuery builder.
func (rdq *ResourceDefinitionQuery) Where(ps ...predicate.ResourceDefinition) *ResourceDefinitionQuery {
	rdq.predicates = append(rdq.predicates, ps...)
	return rdq
}

// Limit the number of records to be returned by this query.
func (rdq *ResourceDefinitionQuery) Limit(limit int) *ResourceDefinitionQuery {
	rdq.ctx.Limit = &limit
	return rdq
}

// Offset to start from.
func (rdq *ResourceDefinitionQuery) Offset(offset int) *ResourceDefinitionQuery {
	rdq.ctx.Offset = &offset
	return rdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rdq *ResourceDefinitionQuery) Unique(unique bool) *ResourceDefinitionQuery {
	rdq.ctx.Unique = &unique
	return rdq
}

// Order specifies how the records should be ordered.
func (rdq *ResourceDefinitionQuery) Order(o ...resourcedefinition.OrderOption) *ResourceDefinitionQuery {
	rdq.order = append(rdq.order, o...)
	return rdq
}

// QueryMatchingRules chains the current query on the "matching_rules" edge.
func (rdq *ResourceDefinitionQuery) QueryMatchingRules() *ResourceDefinitionMatchingRuleQuery {
	query := (&ResourceDefinitionMatchingRuleClient{config: rdq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(resourcedefinition.Table, resourcedefinition.FieldID, selector),
			sqlgraph.To(resourcedefinitionmatchingrule.Table, resourcedefinitionmatchingrule.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, resourcedefinition.MatchingRulesTable, resourcedefinition.MatchingRulesColumn),
		)
		schemaConfig := rdq.schemaConfig
		step.To.Schema = schemaConfig.ResourceDefinitionMatchingRule
		step.Edge.Schema = schemaConfig.ResourceDefinitionMatchingRule
		fromU = sqlgraph.SetNeighbors(rdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryResources chains the current query on the "resources" edge.
func (rdq *ResourceDefinitionQuery) QueryResources() *ResourceQuery {
	query := (&ResourceClient{config: rdq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(resourcedefinition.Table, resourcedefinition.FieldID, selector),
			sqlgraph.To(resource.Table, resource.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, resourcedefinition.ResourcesTable, resourcedefinition.ResourcesColumn),
		)
		schemaConfig := rdq.schemaConfig
		step.To.Schema = schemaConfig.Resource
		step.Edge.Schema = schemaConfig.Resource
		fromU = sqlgraph.SetNeighbors(rdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ResourceDefinition entity from the query.
// Returns a *NotFoundError when no ResourceDefinition was found.
func (rdq *ResourceDefinitionQuery) First(ctx context.Context) (*ResourceDefinition, error) {
	nodes, err := rdq.Limit(1).All(setContextOp(ctx, rdq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{resourcedefinition.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rdq *ResourceDefinitionQuery) FirstX(ctx context.Context) *ResourceDefinition {
	node, err := rdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ResourceDefinition ID from the query.
// Returns a *NotFoundError when no ResourceDefinition ID was found.
func (rdq *ResourceDefinitionQuery) FirstID(ctx context.Context) (id object.ID, err error) {
	var ids []object.ID
	if ids, err = rdq.Limit(1).IDs(setContextOp(ctx, rdq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{resourcedefinition.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rdq *ResourceDefinitionQuery) FirstIDX(ctx context.Context) object.ID {
	id, err := rdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ResourceDefinition entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ResourceDefinition entity is found.
// Returns a *NotFoundError when no ResourceDefinition entities are found.
func (rdq *ResourceDefinitionQuery) Only(ctx context.Context) (*ResourceDefinition, error) {
	nodes, err := rdq.Limit(2).All(setContextOp(ctx, rdq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{resourcedefinition.Label}
	default:
		return nil, &NotSingularError{resourcedefinition.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rdq *ResourceDefinitionQuery) OnlyX(ctx context.Context) *ResourceDefinition {
	node, err := rdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ResourceDefinition ID in the query.
// Returns a *NotSingularError when more than one ResourceDefinition ID is found.
// Returns a *NotFoundError when no entities are found.
func (rdq *ResourceDefinitionQuery) OnlyID(ctx context.Context) (id object.ID, err error) {
	var ids []object.ID
	if ids, err = rdq.Limit(2).IDs(setContextOp(ctx, rdq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{resourcedefinition.Label}
	default:
		err = &NotSingularError{resourcedefinition.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rdq *ResourceDefinitionQuery) OnlyIDX(ctx context.Context) object.ID {
	id, err := rdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ResourceDefinitions.
func (rdq *ResourceDefinitionQuery) All(ctx context.Context) ([]*ResourceDefinition, error) {
	ctx = setContextOp(ctx, rdq.ctx, "All")
	if err := rdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ResourceDefinition, *ResourceDefinitionQuery]()
	return withInterceptors[[]*ResourceDefinition](ctx, rdq, qr, rdq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rdq *ResourceDefinitionQuery) AllX(ctx context.Context) []*ResourceDefinition {
	nodes, err := rdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ResourceDefinition IDs.
func (rdq *ResourceDefinitionQuery) IDs(ctx context.Context) (ids []object.ID, err error) {
	if rdq.ctx.Unique == nil && rdq.path != nil {
		rdq.Unique(true)
	}
	ctx = setContextOp(ctx, rdq.ctx, "IDs")
	if err = rdq.Select(resourcedefinition.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rdq *ResourceDefinitionQuery) IDsX(ctx context.Context) []object.ID {
	ids, err := rdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rdq *ResourceDefinitionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rdq.ctx, "Count")
	if err := rdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rdq, querierCount[*ResourceDefinitionQuery](), rdq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rdq *ResourceDefinitionQuery) CountX(ctx context.Context) int {
	count, err := rdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rdq *ResourceDefinitionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rdq.ctx, "Exist")
	switch _, err := rdq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("model: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rdq *ResourceDefinitionQuery) ExistX(ctx context.Context) bool {
	exist, err := rdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ResourceDefinitionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rdq *ResourceDefinitionQuery) Clone() *ResourceDefinitionQuery {
	if rdq == nil {
		return nil
	}
	return &ResourceDefinitionQuery{
		config:            rdq.config,
		ctx:               rdq.ctx.Clone(),
		order:             append([]resourcedefinition.OrderOption{}, rdq.order...),
		inters:            append([]Interceptor{}, rdq.inters...),
		predicates:        append([]predicate.ResourceDefinition{}, rdq.predicates...),
		withMatchingRules: rdq.withMatchingRules.Clone(),
		withResources:     rdq.withResources.Clone(),
		// clone intermediate query.
		sql:  rdq.sql.Clone(),
		path: rdq.path,
	}
}

// WithMatchingRules tells the query-builder to eager-load the nodes that are connected to
// the "matching_rules" edge. The optional arguments are used to configure the query builder of the edge.
func (rdq *ResourceDefinitionQuery) WithMatchingRules(opts ...func(*ResourceDefinitionMatchingRuleQuery)) *ResourceDefinitionQuery {
	query := (&ResourceDefinitionMatchingRuleClient{config: rdq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rdq.withMatchingRules = query
	return rdq
}

// WithResources tells the query-builder to eager-load the nodes that are connected to
// the "resources" edge. The optional arguments are used to configure the query builder of the edge.
func (rdq *ResourceDefinitionQuery) WithResources(opts ...func(*ResourceQuery)) *ResourceDefinitionQuery {
	query := (&ResourceClient{config: rdq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rdq.withResources = query
	return rdq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ResourceDefinition.Query().
//		GroupBy(resourcedefinition.FieldName).
//		Aggregate(model.Count()).
//		Scan(ctx, &v)
func (rdq *ResourceDefinitionQuery) GroupBy(field string, fields ...string) *ResourceDefinitionGroupBy {
	rdq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ResourceDefinitionGroupBy{build: rdq}
	grbuild.flds = &rdq.ctx.Fields
	grbuild.label = resourcedefinition.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.ResourceDefinition.Query().
//		Select(resourcedefinition.FieldName).
//		Scan(ctx, &v)
func (rdq *ResourceDefinitionQuery) Select(fields ...string) *ResourceDefinitionSelect {
	rdq.ctx.Fields = append(rdq.ctx.Fields, fields...)
	sbuild := &ResourceDefinitionSelect{ResourceDefinitionQuery: rdq}
	sbuild.label = resourcedefinition.Label
	sbuild.flds, sbuild.scan = &rdq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ResourceDefinitionSelect configured with the given aggregations.
func (rdq *ResourceDefinitionQuery) Aggregate(fns ...AggregateFunc) *ResourceDefinitionSelect {
	return rdq.Select().Aggregate(fns...)
}

func (rdq *ResourceDefinitionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rdq.inters {
		if inter == nil {
			return fmt.Errorf("model: uninitialized interceptor (forgotten import model/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rdq); err != nil {
				return err
			}
		}
	}
	for _, f := range rdq.ctx.Fields {
		if !resourcedefinition.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
		}
	}
	if rdq.path != nil {
		prev, err := rdq.path(ctx)
		if err != nil {
			return err
		}
		rdq.sql = prev
	}
	return nil
}

func (rdq *ResourceDefinitionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ResourceDefinition, error) {
	var (
		nodes       = []*ResourceDefinition{}
		_spec       = rdq.querySpec()
		loadedTypes = [2]bool{
			rdq.withMatchingRules != nil,
			rdq.withResources != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ResourceDefinition).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ResourceDefinition{config: rdq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = rdq.schemaConfig.ResourceDefinition
	ctx = internal.NewSchemaConfigContext(ctx, rdq.schemaConfig)
	if len(rdq.modifiers) > 0 {
		_spec.Modifiers = rdq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rdq.withMatchingRules; query != nil {
		if err := rdq.loadMatchingRules(ctx, query, nodes,
			func(n *ResourceDefinition) { n.Edges.MatchingRules = []*ResourceDefinitionMatchingRule{} },
			func(n *ResourceDefinition, e *ResourceDefinitionMatchingRule) {
				n.Edges.MatchingRules = append(n.Edges.MatchingRules, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := rdq.withResources; query != nil {
		if err := rdq.loadResources(ctx, query, nodes,
			func(n *ResourceDefinition) { n.Edges.Resources = []*Resource{} },
			func(n *ResourceDefinition, e *Resource) { n.Edges.Resources = append(n.Edges.Resources, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rdq *ResourceDefinitionQuery) loadMatchingRules(ctx context.Context, query *ResourceDefinitionMatchingRuleQuery, nodes []*ResourceDefinition, init func(*ResourceDefinition), assign func(*ResourceDefinition, *ResourceDefinitionMatchingRule)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[object.ID]*ResourceDefinition)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(resourcedefinitionmatchingrule.FieldResourceDefinitionID)
	}
	query.Where(predicate.ResourceDefinitionMatchingRule(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(resourcedefinition.MatchingRulesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ResourceDefinitionID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "resource_definition_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (rdq *ResourceDefinitionQuery) loadResources(ctx context.Context, query *ResourceQuery, nodes []*ResourceDefinition, init func(*ResourceDefinition), assign func(*ResourceDefinition, *Resource)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[object.ID]*ResourceDefinition)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(resource.FieldResourceDefinitionID)
	}
	query.Where(predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(resourcedefinition.ResourcesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ResourceDefinitionID
		if fk == nil {
			return fmt.Errorf(`foreign-key "resource_definition_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "resource_definition_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (rdq *ResourceDefinitionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rdq.querySpec()
	_spec.Node.Schema = rdq.schemaConfig.ResourceDefinition
	ctx = internal.NewSchemaConfigContext(ctx, rdq.schemaConfig)
	if len(rdq.modifiers) > 0 {
		_spec.Modifiers = rdq.modifiers
	}
	_spec.Node.Columns = rdq.ctx.Fields
	if len(rdq.ctx.Fields) > 0 {
		_spec.Unique = rdq.ctx.Unique != nil && *rdq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rdq.driver, _spec)
}

func (rdq *ResourceDefinitionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(resourcedefinition.Table, resourcedefinition.Columns, sqlgraph.NewFieldSpec(resourcedefinition.FieldID, field.TypeString))
	_spec.From = rdq.sql
	if unique := rdq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rdq.path != nil {
		_spec.Unique = true
	}
	if fields := rdq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, resourcedefinition.FieldID)
		for i := range fields {
			if fields[i] != resourcedefinition.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rdq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rdq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rdq *ResourceDefinitionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rdq.driver.Dialect())
	t1 := builder.Table(resourcedefinition.Table)
	columns := rdq.ctx.Fields
	if len(columns) == 0 {
		columns = resourcedefinition.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rdq.sql != nil {
		selector = rdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rdq.ctx.Unique != nil && *rdq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(rdq.schemaConfig.ResourceDefinition)
	ctx = internal.NewSchemaConfigContext(ctx, rdq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range rdq.modifiers {
		m(selector)
	}
	for _, p := range rdq.predicates {
		p(selector)
	}
	for _, p := range rdq.order {
		p(selector)
	}
	if offset := rdq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rdq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (rdq *ResourceDefinitionQuery) ForUpdate(opts ...sql.LockOption) *ResourceDefinitionQuery {
	if rdq.driver.Dialect() == dialect.Postgres {
		rdq.Unique(false)
	}
	rdq.modifiers = append(rdq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return rdq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (rdq *ResourceDefinitionQuery) ForShare(opts ...sql.LockOption) *ResourceDefinitionQuery {
	if rdq.driver.Dialect() == dialect.Postgres {
		rdq.Unique(false)
	}
	rdq.modifiers = append(rdq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return rdq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (rdq *ResourceDefinitionQuery) Modify(modifiers ...func(s *sql.Selector)) *ResourceDefinitionSelect {
	rdq.modifiers = append(rdq.modifiers, modifiers...)
	return rdq.Select()
}

// WhereP appends storage-level predicates to the ResourceDefinitionQuery builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (rdq *ResourceDefinitionQuery) WhereP(ps ...func(*sql.Selector)) {
	var wps = make([]predicate.ResourceDefinition, 0, len(ps))
	for i := 0; i < len(ps); i++ {
		wps = append(wps, predicate.ResourceDefinition(ps[i]))
	}
	rdq.predicates = append(rdq.predicates, wps...)
}

// ResourceDefinitionGroupBy is the group-by builder for ResourceDefinition entities.
type ResourceDefinitionGroupBy struct {
	selector
	build *ResourceDefinitionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rdgb *ResourceDefinitionGroupBy) Aggregate(fns ...AggregateFunc) *ResourceDefinitionGroupBy {
	rdgb.fns = append(rdgb.fns, fns...)
	return rdgb
}

// Scan applies the selector query and scans the result into the given value.
func (rdgb *ResourceDefinitionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rdgb.build.ctx, "GroupBy")
	if err := rdgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ResourceDefinitionQuery, *ResourceDefinitionGroupBy](ctx, rdgb.build, rdgb, rdgb.build.inters, v)
}

func (rdgb *ResourceDefinitionGroupBy) sqlScan(ctx context.Context, root *ResourceDefinitionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rdgb.fns))
	for _, fn := range rdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rdgb.flds)+len(rdgb.fns))
		for _, f := range *rdgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rdgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rdgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ResourceDefinitionSelect is the builder for selecting fields of ResourceDefinition entities.
type ResourceDefinitionSelect struct {
	*ResourceDefinitionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rds *ResourceDefinitionSelect) Aggregate(fns ...AggregateFunc) *ResourceDefinitionSelect {
	rds.fns = append(rds.fns, fns...)
	return rds
}

// Scan applies the selector query and scans the result into the given value.
func (rds *ResourceDefinitionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rds.ctx, "Select")
	if err := rds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ResourceDefinitionQuery, *ResourceDefinitionSelect](ctx, rds.ResourceDefinitionQuery, rds, rds.inters, v)
}

func (rds *ResourceDefinitionSelect) sqlScan(ctx context.Context, root *ResourceDefinitionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rds.fns))
	for _, fn := range rds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (rds *ResourceDefinitionSelect) Modify(modifiers ...func(s *sql.Selector)) *ResourceDefinitionSelect {
	rds.modifiers = append(rds.modifiers, modifiers...)
	return rds
}
