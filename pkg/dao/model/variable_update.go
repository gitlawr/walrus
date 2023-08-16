// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "seal". DO NOT EDIT.

package model

import (
	"context"
	stdsql "database/sql"
	"errors"
	"fmt"
	"reflect"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/variable"
	"github.com/seal-io/walrus/pkg/dao/types/crypto"
)

// VariableUpdate is the builder for updating Variable entities.
type VariableUpdate struct {
	config
	hooks     []Hook
	mutation  *VariableMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *Variable
}

// Where appends a list predicates to the VariableUpdate builder.
func (vu *VariableUpdate) Where(ps ...predicate.Variable) *VariableUpdate {
	vu.mutation.Where(ps...)
	return vu
}

// SetUpdateTime sets the "update_time" field.
func (vu *VariableUpdate) SetUpdateTime(t time.Time) *VariableUpdate {
	vu.mutation.SetUpdateTime(t)
	return vu
}

// SetValue sets the "value" field.
func (vu *VariableUpdate) SetValue(c crypto.String) *VariableUpdate {
	vu.mutation.SetValue(c)
	return vu
}

// SetSensitive sets the "sensitive" field.
func (vu *VariableUpdate) SetSensitive(b bool) *VariableUpdate {
	vu.mutation.SetSensitive(b)
	return vu
}

// SetNillableSensitive sets the "sensitive" field if the given value is not nil.
func (vu *VariableUpdate) SetNillableSensitive(b *bool) *VariableUpdate {
	if b != nil {
		vu.SetSensitive(*b)
	}
	return vu
}

// SetDescription sets the "description" field.
func (vu *VariableUpdate) SetDescription(s string) *VariableUpdate {
	vu.mutation.SetDescription(s)
	return vu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (vu *VariableUpdate) SetNillableDescription(s *string) *VariableUpdate {
	if s != nil {
		vu.SetDescription(*s)
	}
	return vu
}

// ClearDescription clears the value of the "description" field.
func (vu *VariableUpdate) ClearDescription() *VariableUpdate {
	vu.mutation.ClearDescription()
	return vu
}

// Mutation returns the VariableMutation object of the builder.
func (vu *VariableUpdate) Mutation() *VariableMutation {
	return vu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vu *VariableUpdate) Save(ctx context.Context) (int, error) {
	if err := vu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, vu.sqlSave, vu.mutation, vu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vu *VariableUpdate) SaveX(ctx context.Context) int {
	affected, err := vu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vu *VariableUpdate) Exec(ctx context.Context) error {
	_, err := vu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vu *VariableUpdate) ExecX(ctx context.Context) {
	if err := vu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vu *VariableUpdate) defaults() error {
	if _, ok := vu.mutation.UpdateTime(); !ok {
		if variable.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized variable.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := variable.UpdateDefaultUpdateTime()
		vu.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (vu *VariableUpdate) check() error {
	if v, ok := vu.mutation.Value(); ok {
		if err := variable.ValueValidator(string(v)); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`model: validator failed for field "Variable.value": %w`, err)}
		}
	}
	return nil
}

// Set is different from other Set* methods,
// it sets the value by judging the definition of each field within the entire object.
//
// For default fields, Set calls if the value is not zero.
//
// For no default but required fields, Set calls directly.
//
// For no default but optional fields, Set calls if the value is not zero,
// or clears if the value is zero.
//
// For example:
//
//	## Without Default
//
//	### Required
//
//	db.SetX(obj.X)
//
//	### Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	} else {
//	   db.ClearX()
//	}
//
//	## With Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	}
func (vu *VariableUpdate) Set(obj *Variable) *VariableUpdate {
	// Without Default.
	vu.SetValue(obj.Value)
	vu.SetSensitive(obj.Sensitive)
	if obj.Description != "" {
		vu.SetDescription(obj.Description)
	} else {
		vu.ClearDescription()
	}

	// With Default.
	if obj.UpdateTime != nil {
		vu.SetUpdateTime(*obj.UpdateTime)
	}

	// Record the given object.
	vu.object = obj

	return vu
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (vu *VariableUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *VariableUpdate {
	vu.modifiers = append(vu.modifiers, modifiers...)
	return vu
}

func (vu *VariableUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := vu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(variable.Table, variable.Columns, sqlgraph.NewFieldSpec(variable.FieldID, field.TypeString))
	if ps := vu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vu.mutation.UpdateTime(); ok {
		_spec.SetField(variable.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := vu.mutation.Value(); ok {
		_spec.SetField(variable.FieldValue, field.TypeString, value)
	}
	if value, ok := vu.mutation.Sensitive(); ok {
		_spec.SetField(variable.FieldSensitive, field.TypeBool, value)
	}
	if value, ok := vu.mutation.Description(); ok {
		_spec.SetField(variable.FieldDescription, field.TypeString, value)
	}
	if vu.mutation.DescriptionCleared() {
		_spec.ClearField(variable.FieldDescription, field.TypeString)
	}
	_spec.Node.Schema = vu.schemaConfig.Variable
	ctx = internal.NewSchemaConfigContext(ctx, vu.schemaConfig)
	_spec.AddModifiers(vu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, vu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{variable.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vu.mutation.done = true
	return n, nil
}

// VariableUpdateOne is the builder for updating a single Variable entity.
type VariableUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *VariableMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *Variable
}

// SetUpdateTime sets the "update_time" field.
func (vuo *VariableUpdateOne) SetUpdateTime(t time.Time) *VariableUpdateOne {
	vuo.mutation.SetUpdateTime(t)
	return vuo
}

// SetValue sets the "value" field.
func (vuo *VariableUpdateOne) SetValue(c crypto.String) *VariableUpdateOne {
	vuo.mutation.SetValue(c)
	return vuo
}

// SetSensitive sets the "sensitive" field.
func (vuo *VariableUpdateOne) SetSensitive(b bool) *VariableUpdateOne {
	vuo.mutation.SetSensitive(b)
	return vuo
}

// SetNillableSensitive sets the "sensitive" field if the given value is not nil.
func (vuo *VariableUpdateOne) SetNillableSensitive(b *bool) *VariableUpdateOne {
	if b != nil {
		vuo.SetSensitive(*b)
	}
	return vuo
}

// SetDescription sets the "description" field.
func (vuo *VariableUpdateOne) SetDescription(s string) *VariableUpdateOne {
	vuo.mutation.SetDescription(s)
	return vuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (vuo *VariableUpdateOne) SetNillableDescription(s *string) *VariableUpdateOne {
	if s != nil {
		vuo.SetDescription(*s)
	}
	return vuo
}

// ClearDescription clears the value of the "description" field.
func (vuo *VariableUpdateOne) ClearDescription() *VariableUpdateOne {
	vuo.mutation.ClearDescription()
	return vuo
}

// Mutation returns the VariableMutation object of the builder.
func (vuo *VariableUpdateOne) Mutation() *VariableMutation {
	return vuo.mutation
}

// Where appends a list predicates to the VariableUpdate builder.
func (vuo *VariableUpdateOne) Where(ps ...predicate.Variable) *VariableUpdateOne {
	vuo.mutation.Where(ps...)
	return vuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vuo *VariableUpdateOne) Select(field string, fields ...string) *VariableUpdateOne {
	vuo.fields = append([]string{field}, fields...)
	return vuo
}

// Save executes the query and returns the updated Variable entity.
func (vuo *VariableUpdateOne) Save(ctx context.Context) (*Variable, error) {
	if err := vuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, vuo.sqlSave, vuo.mutation, vuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vuo *VariableUpdateOne) SaveX(ctx context.Context) *Variable {
	node, err := vuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuo *VariableUpdateOne) Exec(ctx context.Context) error {
	_, err := vuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *VariableUpdateOne) ExecX(ctx context.Context) {
	if err := vuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vuo *VariableUpdateOne) defaults() error {
	if _, ok := vuo.mutation.UpdateTime(); !ok {
		if variable.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized variable.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := variable.UpdateDefaultUpdateTime()
		vuo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (vuo *VariableUpdateOne) check() error {
	if v, ok := vuo.mutation.Value(); ok {
		if err := variable.ValueValidator(string(v)); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`model: validator failed for field "Variable.value": %w`, err)}
		}
	}
	return nil
}

// Set is different from other Set* methods,
// it sets the value by judging the definition of each field within the entire object.
//
// For default fields, Set calls if the value changes from the original.
//
// For no default but required fields, Set calls if the value changes from the original.
//
// For no default but optional fields, Set calls if the value changes from the original,
// or clears if changes to zero.
//
// For example:
//
//	## Without Default
//
//	### Required
//
//	db.SetX(obj.X)
//
//	### Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   if _is_not_equal_(db.X, obj.X) {
//	      db.SetX(obj.X)
//	   }
//	} else {
//	   db.ClearX()
//	}
//
//	## With Default
//
//	if _is_zero_value_(obj.X) && _is_not_equal_(db.X, obj.X) {
//	   db.SetX(obj.X)
//	}
func (vuo *VariableUpdateOne) Set(obj *Variable) *VariableUpdateOne {
	h := func(n ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mt := m.(*VariableMutation)
			db, err := mt.Client().Variable.Get(ctx, *mt.id)
			if err != nil {
				return nil, fmt.Errorf("failed getting Variable with id: %v", *mt.id)
			}

			// Without Default.
			if db.Value != obj.Value {
				vuo.SetValue(obj.Value)
			}
			if db.Sensitive != obj.Sensitive {
				vuo.SetSensitive(obj.Sensitive)
			}
			if obj.Description != "" {
				if db.Description != obj.Description {
					vuo.SetDescription(obj.Description)
				}
			} else {
				vuo.ClearDescription()
			}

			// With Default.
			if (obj.UpdateTime != nil) && (!reflect.DeepEqual(db.UpdateTime, obj.UpdateTime)) {
				vuo.SetUpdateTime(*obj.UpdateTime)
			}

			// Record the given object.
			vuo.object = obj

			return n.Mutate(ctx, m)
		})
	}

	vuo.hooks = append(vuo.hooks, h)

	return vuo
}

// getClientSet returns the ClientSet for the given builder.
func (vuo *VariableUpdateOne) getClientSet() (mc ClientSet) {
	if _, ok := vuo.config.driver.(*txDriver); ok {
		tx := &Tx{config: vuo.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: vuo.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after updated the Variable entity,
// which is always good for cascading update operations.
func (vuo *VariableUpdateOne) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Variable) error) (*Variable, error) {
	obj, err := vuo.Save(ctx)
	if err != nil &&
		(vuo.object == nil || !errors.Is(err, stdsql.ErrNoRows)) {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := vuo.getClientSet()

	if obj == nil {
		obj = vuo.object
	} else if x := vuo.object; x != nil {
		if _, set := vuo.mutation.Field(variable.FieldValue); set {
			obj.Value = x.Value
		}
		if _, set := vuo.mutation.Field(variable.FieldSensitive); set {
			obj.Sensitive = x.Sensitive
		}
		if _, set := vuo.mutation.Field(variable.FieldDescription); set {
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
func (vuo *VariableUpdateOne) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Variable) error) *Variable {
	obj, err := vuo.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading update operations.
func (vuo *VariableUpdateOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Variable) error) error {
	_, err := vuo.SaveE(ctx, cbs...)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *VariableUpdateOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Variable) error) {
	if err := vuo.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (vuo *VariableUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *VariableUpdateOne {
	vuo.modifiers = append(vuo.modifiers, modifiers...)
	return vuo
}

func (vuo *VariableUpdateOne) sqlSave(ctx context.Context) (_node *Variable, err error) {
	if err := vuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(variable.Table, variable.Columns, sqlgraph.NewFieldSpec(variable.FieldID, field.TypeString))
	id, ok := vuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "Variable.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, variable.FieldID)
		for _, f := range fields {
			if !variable.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != variable.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vuo.mutation.UpdateTime(); ok {
		_spec.SetField(variable.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := vuo.mutation.Value(); ok {
		_spec.SetField(variable.FieldValue, field.TypeString, value)
	}
	if value, ok := vuo.mutation.Sensitive(); ok {
		_spec.SetField(variable.FieldSensitive, field.TypeBool, value)
	}
	if value, ok := vuo.mutation.Description(); ok {
		_spec.SetField(variable.FieldDescription, field.TypeString, value)
	}
	if vuo.mutation.DescriptionCleared() {
		_spec.ClearField(variable.FieldDescription, field.TypeString)
	}
	_spec.Node.Schema = vuo.schemaConfig.Variable
	ctx = internal.NewSchemaConfigContext(ctx, vuo.schemaConfig)
	_spec.AddModifiers(vuo.modifiers...)
	_node = &Variable{config: vuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{variable.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vuo.mutation.done = true
	return _node, nil
}
