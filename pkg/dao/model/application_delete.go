// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/seal/pkg/dao/model/application"
	"github.com/seal-io/seal/pkg/dao/model/internal"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
)

// ApplicationDelete is the builder for deleting a Application entity.
type ApplicationDelete struct {
	config
	hooks    []Hook
	mutation *ApplicationMutation
}

// Where appends a list predicates to the ApplicationDelete builder.
func (ad *ApplicationDelete) Where(ps ...predicate.Application) *ApplicationDelete {
	ad.mutation.Where(ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *ApplicationDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ApplicationMutation](ctx, ad.sqlExec, ad.mutation, ad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *ApplicationDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *ApplicationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: application.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeOther,
				Column: application.FieldID,
			},
		},
	}
	_spec.Node.Schema = ad.schemaConfig.Application
	ctx = internal.NewSchemaConfigContext(ctx, ad.schemaConfig)
	if ps := ad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ad.mutation.done = true
	return affected, err
}

// ApplicationDeleteOne is the builder for deleting a single Application entity.
type ApplicationDeleteOne struct {
	ad *ApplicationDelete
}

// Where appends a list predicates to the ApplicationDelete builder.
func (ado *ApplicationDeleteOne) Where(ps ...predicate.Application) *ApplicationDeleteOne {
	ado.ad.mutation.Where(ps...)
	return ado
}

// Exec executes the deletion query.
func (ado *ApplicationDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{application.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *ApplicationDeleteOne) ExecX(ctx context.Context) {
	if err := ado.Exec(ctx); err != nil {
		panic(err)
	}
}