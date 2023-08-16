// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "seal". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/subject"
	"github.com/seal-io/walrus/pkg/dao/model/subjectrolerelationship"
	"github.com/seal-io/walrus/pkg/dao/model/token"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// SubjectCreate is the builder for creating a Subject entity.
type SubjectCreate struct {
	config
	mutation   *SubjectMutation
	hooks      []Hook
	conflict   []sql.ConflictOption
	object     *Subject
	fromUpsert bool
}

// SetCreateTime sets the "create_time" field.
func (sc *SubjectCreate) SetCreateTime(t time.Time) *SubjectCreate {
	sc.mutation.SetCreateTime(t)
	return sc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (sc *SubjectCreate) SetNillableCreateTime(t *time.Time) *SubjectCreate {
	if t != nil {
		sc.SetCreateTime(*t)
	}
	return sc
}

// SetUpdateTime sets the "update_time" field.
func (sc *SubjectCreate) SetUpdateTime(t time.Time) *SubjectCreate {
	sc.mutation.SetUpdateTime(t)
	return sc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (sc *SubjectCreate) SetNillableUpdateTime(t *time.Time) *SubjectCreate {
	if t != nil {
		sc.SetUpdateTime(*t)
	}
	return sc
}

// SetKind sets the "kind" field.
func (sc *SubjectCreate) SetKind(s string) *SubjectCreate {
	sc.mutation.SetKind(s)
	return sc
}

// SetNillableKind sets the "kind" field if the given value is not nil.
func (sc *SubjectCreate) SetNillableKind(s *string) *SubjectCreate {
	if s != nil {
		sc.SetKind(*s)
	}
	return sc
}

// SetDomain sets the "domain" field.
func (sc *SubjectCreate) SetDomain(s string) *SubjectCreate {
	sc.mutation.SetDomain(s)
	return sc
}

// SetNillableDomain sets the "domain" field if the given value is not nil.
func (sc *SubjectCreate) SetNillableDomain(s *string) *SubjectCreate {
	if s != nil {
		sc.SetDomain(*s)
	}
	return sc
}

// SetName sets the "name" field.
func (sc *SubjectCreate) SetName(s string) *SubjectCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetDescription sets the "description" field.
func (sc *SubjectCreate) SetDescription(s string) *SubjectCreate {
	sc.mutation.SetDescription(s)
	return sc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (sc *SubjectCreate) SetNillableDescription(s *string) *SubjectCreate {
	if s != nil {
		sc.SetDescription(*s)
	}
	return sc
}

// SetBuiltin sets the "builtin" field.
func (sc *SubjectCreate) SetBuiltin(b bool) *SubjectCreate {
	sc.mutation.SetBuiltin(b)
	return sc
}

// SetNillableBuiltin sets the "builtin" field if the given value is not nil.
func (sc *SubjectCreate) SetNillableBuiltin(b *bool) *SubjectCreate {
	if b != nil {
		sc.SetBuiltin(*b)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *SubjectCreate) SetID(o object.ID) *SubjectCreate {
	sc.mutation.SetID(o)
	return sc
}

// AddTokenIDs adds the "tokens" edge to the Token entity by IDs.
func (sc *SubjectCreate) AddTokenIDs(ids ...object.ID) *SubjectCreate {
	sc.mutation.AddTokenIDs(ids...)
	return sc
}

// AddTokens adds the "tokens" edges to the Token entity.
func (sc *SubjectCreate) AddTokens(t ...*Token) *SubjectCreate {
	ids := make([]object.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return sc.AddTokenIDs(ids...)
}

// AddRoleIDs adds the "roles" edge to the SubjectRoleRelationship entity by IDs.
func (sc *SubjectCreate) AddRoleIDs(ids ...object.ID) *SubjectCreate {
	sc.mutation.AddRoleIDs(ids...)
	return sc
}

// AddRoles adds the "roles" edges to the SubjectRoleRelationship entity.
func (sc *SubjectCreate) AddRoles(s ...*SubjectRoleRelationship) *SubjectCreate {
	ids := make([]object.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sc.AddRoleIDs(ids...)
}

// Mutation returns the SubjectMutation object of the builder.
func (sc *SubjectCreate) Mutation() *SubjectMutation {
	return sc.mutation
}

// Save creates the Subject in the database.
func (sc *SubjectCreate) Save(ctx context.Context) (*Subject, error) {
	if err := sc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SubjectCreate) SaveX(ctx context.Context) *Subject {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SubjectCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SubjectCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SubjectCreate) defaults() error {
	if _, ok := sc.mutation.CreateTime(); !ok {
		if subject.DefaultCreateTime == nil {
			return fmt.Errorf("model: uninitialized subject.DefaultCreateTime (forgotten import model/runtime?)")
		}
		v := subject.DefaultCreateTime()
		sc.mutation.SetCreateTime(v)
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		if subject.DefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized subject.DefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := subject.DefaultUpdateTime()
		sc.mutation.SetUpdateTime(v)
	}
	if _, ok := sc.mutation.Kind(); !ok {
		v := subject.DefaultKind
		sc.mutation.SetKind(v)
	}
	if _, ok := sc.mutation.Domain(); !ok {
		v := subject.DefaultDomain
		sc.mutation.SetDomain(v)
	}
	if _, ok := sc.mutation.Builtin(); !ok {
		v := subject.DefaultBuiltin
		sc.mutation.SetBuiltin(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (sc *SubjectCreate) check() error {
	if _, ok := sc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "Subject.create_time"`)}
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`model: missing required field "Subject.update_time"`)}
	}
	if _, ok := sc.mutation.Kind(); !ok {
		return &ValidationError{Name: "kind", err: errors.New(`model: missing required field "Subject.kind"`)}
	}
	if _, ok := sc.mutation.Domain(); !ok {
		return &ValidationError{Name: "domain", err: errors.New(`model: missing required field "Subject.domain"`)}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`model: missing required field "Subject.name"`)}
	}
	if v, ok := sc.mutation.Name(); ok {
		if err := subject.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`model: validator failed for field "Subject.name": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Builtin(); !ok {
		return &ValidationError{Name: "builtin", err: errors.New(`model: missing required field "Subject.builtin"`)}
	}
	return nil
}

func (sc *SubjectCreate) sqlSave(ctx context.Context) (*Subject, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*object.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SubjectCreate) createSpec() (*Subject, *sqlgraph.CreateSpec) {
	var (
		_node = &Subject{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(subject.Table, sqlgraph.NewFieldSpec(subject.FieldID, field.TypeString))
	)
	_spec.Schema = sc.schemaConfig.Subject
	_spec.OnConflict = sc.conflict
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.CreateTime(); ok {
		_spec.SetField(subject.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = &value
	}
	if value, ok := sc.mutation.UpdateTime(); ok {
		_spec.SetField(subject.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = &value
	}
	if value, ok := sc.mutation.Kind(); ok {
		_spec.SetField(subject.FieldKind, field.TypeString, value)
		_node.Kind = value
	}
	if value, ok := sc.mutation.Domain(); ok {
		_spec.SetField(subject.FieldDomain, field.TypeString, value)
		_node.Domain = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(subject.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Description(); ok {
		_spec.SetField(subject.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := sc.mutation.Builtin(); ok {
		_spec.SetField(subject.FieldBuiltin, field.TypeBool, value)
		_node.Builtin = value
	}
	if nodes := sc.mutation.TokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.TokensTable,
			Columns: []string{subject.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeString),
			},
		}
		edge.Schema = sc.schemaConfig.Token
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   subject.RolesTable,
			Columns: []string{subject.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subjectrolerelationship.FieldID, field.TypeString),
			},
		}
		edge.Schema = sc.schemaConfig.SubjectRoleRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// Set is different from other Set* methods,
// it sets the value by judging the definition of each field within the entire object.
//
// For required fields, Set calls directly.
//
// For optional fields, Set calls if the value is not zero.
//
// For example:
//
//	## Required
//
//	db.SetX(obj.X)
//
//	## Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	}
func (sc *SubjectCreate) Set(obj *Subject) *SubjectCreate {
	// Required.
	sc.SetKind(obj.Kind)
	sc.SetDomain(obj.Domain)
	sc.SetName(obj.Name)
	sc.SetBuiltin(obj.Builtin)

	// Optional.
	if obj.CreateTime != nil {
		sc.SetCreateTime(*obj.CreateTime)
	}
	if obj.UpdateTime != nil {
		sc.SetUpdateTime(*obj.UpdateTime)
	}
	if obj.Description != "" {
		sc.SetDescription(obj.Description)
	}

	// Record the given object.
	sc.object = obj

	return sc
}

// getClientSet returns the ClientSet for the given builder.
func (sc *SubjectCreate) getClientSet() (mc ClientSet) {
	if _, ok := sc.config.driver.(*txDriver); ok {
		tx := &Tx{config: sc.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: sc.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after created the Subject entity,
// which is always good for cascading create operations.
func (sc *SubjectCreate) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Subject) error) (*Subject, error) {
	obj, err := sc.Save(ctx)
	if err != nil {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := sc.getClientSet()
	if sc.fromUpsert {
		q := mc.Subjects().Query().
			Where(
				subject.Kind(obj.Kind),
				subject.Domain(obj.Domain),
				subject.Name(obj.Name),
			)
		obj.ID, err = q.OnlyID(ctx)
		if err != nil {
			return nil, fmt.Errorf("model: failed to query id of Subject entity: %w", err)
		}
	}

	if x := sc.object; x != nil {
		if _, set := sc.mutation.Field(subject.FieldName); set {
			obj.Name = x.Name
		}
		if _, set := sc.mutation.Field(subject.FieldDescription); set {
			obj.Description = x.Description
		}
		obj.Edges = x.Edges
	}

	for i := range cbs {
		if err = cbs[i](ctx, mc, obj); err != nil {
			return nil, err
		}
	}

	return obj, nil
}

// SaveEX is like SaveE, but panics if an error occurs.
func (sc *SubjectCreate) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Subject) error) *Subject {
	obj, err := sc.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (sc *SubjectCreate) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Subject) error) error {
	_, err := sc.SaveE(ctx, cbs...)
	return err
}

// ExecEX is like ExecE, but panics if an error occurs.
func (sc *SubjectCreate) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Subject) error) {
	if err := sc.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Set leverages the SubjectCreate Set method,
// it sets the value by judging the definition of each field within the entire item of the given list.
//
// For required fields, Set calls directly.
//
// For optional fields, Set calls if the value is not zero.
//
// For example:
//
//	## Required
//
//	db.SetX(obj.X)
//
//	## Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	}
func (scb *SubjectCreateBulk) Set(objs ...*Subject) *SubjectCreateBulk {
	if len(objs) != 0 {
		client := NewSubjectClient(scb.config)

		scb.builders = make([]*SubjectCreate, len(objs))
		for i := range objs {
			scb.builders[i] = client.Create().Set(objs[i])
		}

		// Record the given objects.
		scb.objects = objs
	}

	return scb
}

// getClientSet returns the ClientSet for the given builder.
func (scb *SubjectCreateBulk) getClientSet() (mc ClientSet) {
	if _, ok := scb.config.driver.(*txDriver); ok {
		tx := &Tx{config: scb.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: scb.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after created the Subject entities,
// which is always good for cascading create operations.
func (scb *SubjectCreateBulk) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Subject) error) ([]*Subject, error) {
	objs, err := scb.Save(ctx)
	if err != nil {
		return nil, err
	}

	if len(cbs) == 0 {
		return objs, err
	}

	mc := scb.getClientSet()
	if scb.fromUpsert {
		for i := range objs {
			obj := objs[i]
			q := mc.Subjects().Query().
				Where(
					subject.Kind(obj.Kind),
					subject.Domain(obj.Domain),
					subject.Name(obj.Name),
				)
			objs[i].ID, err = q.OnlyID(ctx)
			if err != nil {
				return nil, fmt.Errorf("model: failed to query id of Subject entity: %w", err)
			}
		}
	}

	if x := scb.objects; x != nil {
		for i := range x {
			if _, set := scb.builders[i].mutation.Field(subject.FieldName); set {
				objs[i].Name = x[i].Name
			}
			if _, set := scb.builders[i].mutation.Field(subject.FieldDescription); set {
				objs[i].Description = x[i].Description
			}
			objs[i].Edges = x[i].Edges
		}
	}

	for i := range objs {
		for j := range cbs {
			if err = cbs[j](ctx, mc, objs[i]); err != nil {
				return nil, err
			}
		}
	}

	return objs, nil
}

// SaveEX is like SaveE, but panics if an error occurs.
func (scb *SubjectCreateBulk) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Subject) error) []*Subject {
	objs, err := scb.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return objs
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (scb *SubjectCreateBulk) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Subject) error) error {
	_, err := scb.SaveE(ctx, cbs...)
	return err
}

// ExecEX is like ExecE, but panics if an error occurs.
func (scb *SubjectCreateBulk) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Subject) error) {
	if err := scb.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (u *SubjectUpsertOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Subject) error) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for SubjectUpsertOne.OnConflict")
	}
	u.create.fromUpsert = true
	return u.create.ExecE(ctx, cbs...)
}

// ExecEX is like ExecE, but panics if an error occurs.
func (u *SubjectUpsertOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Subject) error) {
	if err := u.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (u *SubjectUpsertBulk) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Subject) error) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the SubjectUpsertBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for SubjectUpsertBulk.OnConflict")
	}
	u.create.fromUpsert = true
	return u.create.ExecE(ctx, cbs...)
}

// ExecEX is like ExecE, but panics if an error occurs.
func (u *SubjectUpsertBulk) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Subject) error) {
	if err := u.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Subject.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SubjectUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (sc *SubjectCreate) OnConflict(opts ...sql.ConflictOption) *SubjectUpsertOne {
	sc.conflict = opts
	return &SubjectUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Subject.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *SubjectCreate) OnConflictColumns(columns ...string) *SubjectUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &SubjectUpsertOne{
		create: sc,
	}
}

type (
	// SubjectUpsertOne is the builder for "upsert"-ing
	//  one Subject node.
	SubjectUpsertOne struct {
		create *SubjectCreate
	}

	// SubjectUpsert is the "OnConflict" setter.
	SubjectUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *SubjectUpsert) SetUpdateTime(v time.Time) *SubjectUpsert {
	u.Set(subject.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *SubjectUpsert) UpdateUpdateTime() *SubjectUpsert {
	u.SetExcluded(subject.FieldUpdateTime)
	return u
}

// SetDescription sets the "description" field.
func (u *SubjectUpsert) SetDescription(v string) *SubjectUpsert {
	u.Set(subject.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *SubjectUpsert) UpdateDescription() *SubjectUpsert {
	u.SetExcluded(subject.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *SubjectUpsert) ClearDescription() *SubjectUpsert {
	u.SetNull(subject.FieldDescription)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Subject.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(subject.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SubjectUpsertOne) UpdateNewValues() *SubjectUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(subject.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(subject.FieldCreateTime)
		}
		if _, exists := u.create.mutation.Kind(); exists {
			s.SetIgnore(subject.FieldKind)
		}
		if _, exists := u.create.mutation.Domain(); exists {
			s.SetIgnore(subject.FieldDomain)
		}
		if _, exists := u.create.mutation.Name(); exists {
			s.SetIgnore(subject.FieldName)
		}
		if _, exists := u.create.mutation.Builtin(); exists {
			s.SetIgnore(subject.FieldBuiltin)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Subject.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SubjectUpsertOne) Ignore() *SubjectUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SubjectUpsertOne) DoNothing() *SubjectUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SubjectCreate.OnConflict
// documentation for more info.
func (u *SubjectUpsertOne) Update(set func(*SubjectUpsert)) *SubjectUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SubjectUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *SubjectUpsertOne) SetUpdateTime(v time.Time) *SubjectUpsertOne {
	return u.Update(func(s *SubjectUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *SubjectUpsertOne) UpdateUpdateTime() *SubjectUpsertOne {
	return u.Update(func(s *SubjectUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetDescription sets the "description" field.
func (u *SubjectUpsertOne) SetDescription(v string) *SubjectUpsertOne {
	return u.Update(func(s *SubjectUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *SubjectUpsertOne) UpdateDescription() *SubjectUpsertOne {
	return u.Update(func(s *SubjectUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *SubjectUpsertOne) ClearDescription() *SubjectUpsertOne {
	return u.Update(func(s *SubjectUpsert) {
		s.ClearDescription()
	})
}

// Exec executes the query.
func (u *SubjectUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for SubjectCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SubjectUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SubjectUpsertOne) ID(ctx context.Context) (id object.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("model: SubjectUpsertOne.ID is not supported by MySQL driver. Use SubjectUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SubjectUpsertOne) IDX(ctx context.Context) object.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SubjectCreateBulk is the builder for creating many Subject entities in bulk.
type SubjectCreateBulk struct {
	config
	builders   []*SubjectCreate
	conflict   []sql.ConflictOption
	objects    []*Subject
	fromUpsert bool
}

// Save creates the Subject entities in the database.
func (scb *SubjectCreateBulk) Save(ctx context.Context) ([]*Subject, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Subject, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubjectMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SubjectCreateBulk) SaveX(ctx context.Context) []*Subject {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SubjectCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SubjectCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Subject.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SubjectUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (scb *SubjectCreateBulk) OnConflict(opts ...sql.ConflictOption) *SubjectUpsertBulk {
	scb.conflict = opts
	return &SubjectUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Subject.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *SubjectCreateBulk) OnConflictColumns(columns ...string) *SubjectUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &SubjectUpsertBulk{
		create: scb,
	}
}

// SubjectUpsertBulk is the builder for "upsert"-ing
// a bulk of Subject nodes.
type SubjectUpsertBulk struct {
	create *SubjectCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Subject.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(subject.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SubjectUpsertBulk) UpdateNewValues() *SubjectUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(subject.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(subject.FieldCreateTime)
			}
			if _, exists := b.mutation.Kind(); exists {
				s.SetIgnore(subject.FieldKind)
			}
			if _, exists := b.mutation.Domain(); exists {
				s.SetIgnore(subject.FieldDomain)
			}
			if _, exists := b.mutation.Name(); exists {
				s.SetIgnore(subject.FieldName)
			}
			if _, exists := b.mutation.Builtin(); exists {
				s.SetIgnore(subject.FieldBuiltin)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Subject.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SubjectUpsertBulk) Ignore() *SubjectUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SubjectUpsertBulk) DoNothing() *SubjectUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SubjectCreateBulk.OnConflict
// documentation for more info.
func (u *SubjectUpsertBulk) Update(set func(*SubjectUpsert)) *SubjectUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SubjectUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *SubjectUpsertBulk) SetUpdateTime(v time.Time) *SubjectUpsertBulk {
	return u.Update(func(s *SubjectUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *SubjectUpsertBulk) UpdateUpdateTime() *SubjectUpsertBulk {
	return u.Update(func(s *SubjectUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetDescription sets the "description" field.
func (u *SubjectUpsertBulk) SetDescription(v string) *SubjectUpsertBulk {
	return u.Update(func(s *SubjectUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *SubjectUpsertBulk) UpdateDescription() *SubjectUpsertBulk {
	return u.Update(func(s *SubjectUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *SubjectUpsertBulk) ClearDescription() *SubjectUpsertBulk {
	return u.Update(func(s *SubjectUpsert) {
		s.ClearDescription()
	})
}

// Exec executes the query.
func (u *SubjectUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the SubjectCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for SubjectCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SubjectUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
