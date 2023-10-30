// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/servicerevision"
)

// ServiceRevisionDelete is the builder for deleting a ServiceRevision entity.
type ServiceRevisionDelete struct {
	config
	hooks    []Hook
	mutation *ServiceRevisionMutation
}

// Where appends a list predicates to the ServiceRevisionDelete builder.
func (srd *ServiceRevisionDelete) Where(ps ...predicate.ServiceRevision) *ServiceRevisionDelete {
	srd.mutation.Where(ps...)
	return srd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (srd *ServiceRevisionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, srd.sqlExec, srd.mutation, srd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (srd *ServiceRevisionDelete) ExecX(ctx context.Context) int {
	n, err := srd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (srd *ServiceRevisionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(servicerevision.Table, sqlgraph.NewFieldSpec(servicerevision.FieldID, field.TypeString))
	_spec.Node.Schema = srd.schemaConfig.ServiceRevision
	ctx = internal.NewSchemaConfigContext(ctx, srd.schemaConfig)
	if ps := srd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, srd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	srd.mutation.done = true
	return affected, err
}

// ServiceRevisionDeleteOne is the builder for deleting a single ServiceRevision entity.
type ServiceRevisionDeleteOne struct {
	srd *ServiceRevisionDelete
}

// Where appends a list predicates to the ServiceRevisionDelete builder.
func (srdo *ServiceRevisionDeleteOne) Where(ps ...predicate.ServiceRevision) *ServiceRevisionDeleteOne {
	srdo.srd.mutation.Where(ps...)
	return srdo
}

// Exec executes the deletion query.
func (srdo *ServiceRevisionDeleteOne) Exec(ctx context.Context) error {
	n, err := srdo.srd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{servicerevision.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (srdo *ServiceRevisionDeleteOne) ExecX(ctx context.Context) {
	if err := srdo.Exec(ctx); err != nil {
		panic(err)
	}
}