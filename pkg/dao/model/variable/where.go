// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package variable

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"

	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/types/crypto"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// ID filters vertices based on their ID field.
func ID(id object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Variable {
	return predicate.Variable(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Variable {
	return predicate.Variable(sql.FieldEQ(FieldUpdateTime, v))
}

// ProjectID applies equality check predicate on the "project_id" field. It's identical to ProjectIDEQ.
func ProjectID(v object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldEQ(FieldProjectID, v))
}

// EnvironmentID applies equality check predicate on the "environment_id" field. It's identical to EnvironmentIDEQ.
func EnvironmentID(v object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldEQ(FieldEnvironmentID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Variable {
	return predicate.Variable(sql.FieldEQ(FieldName, v))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v crypto.String) predicate.Variable {
	return predicate.Variable(sql.FieldEQ(FieldValue, v))
}

// Sensitive applies equality check predicate on the "sensitive" field. It's identical to SensitiveEQ.
func Sensitive(v bool) predicate.Variable {
	return predicate.Variable(sql.FieldEQ(FieldSensitive, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Variable {
	return predicate.Variable(sql.FieldEQ(FieldDescription, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Variable {
	return predicate.Variable(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Variable {
	return predicate.Variable(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Variable {
	return predicate.Variable(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Variable {
	return predicate.Variable(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Variable {
	return predicate.Variable(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Variable {
	return predicate.Variable(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Variable {
	return predicate.Variable(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Variable {
	return predicate.Variable(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Variable {
	return predicate.Variable(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Variable {
	return predicate.Variable(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Variable {
	return predicate.Variable(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Variable {
	return predicate.Variable(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Variable {
	return predicate.Variable(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Variable {
	return predicate.Variable(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Variable {
	return predicate.Variable(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Variable {
	return predicate.Variable(sql.FieldLTE(FieldUpdateTime, v))
}

// ProjectIDEQ applies the EQ predicate on the "project_id" field.
func ProjectIDEQ(v object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldEQ(FieldProjectID, v))
}

// ProjectIDNEQ applies the NEQ predicate on the "project_id" field.
func ProjectIDNEQ(v object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldNEQ(FieldProjectID, v))
}

// ProjectIDIn applies the In predicate on the "project_id" field.
func ProjectIDIn(vs ...object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldIn(FieldProjectID, vs...))
}

// ProjectIDNotIn applies the NotIn predicate on the "project_id" field.
func ProjectIDNotIn(vs ...object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldNotIn(FieldProjectID, vs...))
}

// ProjectIDGT applies the GT predicate on the "project_id" field.
func ProjectIDGT(v object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldGT(FieldProjectID, v))
}

// ProjectIDGTE applies the GTE predicate on the "project_id" field.
func ProjectIDGTE(v object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldGTE(FieldProjectID, v))
}

// ProjectIDLT applies the LT predicate on the "project_id" field.
func ProjectIDLT(v object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldLT(FieldProjectID, v))
}

// ProjectIDLTE applies the LTE predicate on the "project_id" field.
func ProjectIDLTE(v object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldLTE(FieldProjectID, v))
}

// ProjectIDContains applies the Contains predicate on the "project_id" field.
func ProjectIDContains(v object.ID) predicate.Variable {
	vc := string(v)
	return predicate.Variable(sql.FieldContains(FieldProjectID, vc))
}

// ProjectIDHasPrefix applies the HasPrefix predicate on the "project_id" field.
func ProjectIDHasPrefix(v object.ID) predicate.Variable {
	vc := string(v)
	return predicate.Variable(sql.FieldHasPrefix(FieldProjectID, vc))
}

// ProjectIDHasSuffix applies the HasSuffix predicate on the "project_id" field.
func ProjectIDHasSuffix(v object.ID) predicate.Variable {
	vc := string(v)
	return predicate.Variable(sql.FieldHasSuffix(FieldProjectID, vc))
}

// ProjectIDIsNil applies the IsNil predicate on the "project_id" field.
func ProjectIDIsNil() predicate.Variable {
	return predicate.Variable(sql.FieldIsNull(FieldProjectID))
}

// ProjectIDNotNil applies the NotNil predicate on the "project_id" field.
func ProjectIDNotNil() predicate.Variable {
	return predicate.Variable(sql.FieldNotNull(FieldProjectID))
}

// ProjectIDEqualFold applies the EqualFold predicate on the "project_id" field.
func ProjectIDEqualFold(v object.ID) predicate.Variable {
	vc := string(v)
	return predicate.Variable(sql.FieldEqualFold(FieldProjectID, vc))
}

// ProjectIDContainsFold applies the ContainsFold predicate on the "project_id" field.
func ProjectIDContainsFold(v object.ID) predicate.Variable {
	vc := string(v)
	return predicate.Variable(sql.FieldContainsFold(FieldProjectID, vc))
}

// EnvironmentIDEQ applies the EQ predicate on the "environment_id" field.
func EnvironmentIDEQ(v object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldEQ(FieldEnvironmentID, v))
}

// EnvironmentIDNEQ applies the NEQ predicate on the "environment_id" field.
func EnvironmentIDNEQ(v object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldNEQ(FieldEnvironmentID, v))
}

// EnvironmentIDIn applies the In predicate on the "environment_id" field.
func EnvironmentIDIn(vs ...object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldIn(FieldEnvironmentID, vs...))
}

// EnvironmentIDNotIn applies the NotIn predicate on the "environment_id" field.
func EnvironmentIDNotIn(vs ...object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldNotIn(FieldEnvironmentID, vs...))
}

// EnvironmentIDGT applies the GT predicate on the "environment_id" field.
func EnvironmentIDGT(v object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldGT(FieldEnvironmentID, v))
}

// EnvironmentIDGTE applies the GTE predicate on the "environment_id" field.
func EnvironmentIDGTE(v object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldGTE(FieldEnvironmentID, v))
}

// EnvironmentIDLT applies the LT predicate on the "environment_id" field.
func EnvironmentIDLT(v object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldLT(FieldEnvironmentID, v))
}

// EnvironmentIDLTE applies the LTE predicate on the "environment_id" field.
func EnvironmentIDLTE(v object.ID) predicate.Variable {
	return predicate.Variable(sql.FieldLTE(FieldEnvironmentID, v))
}

// EnvironmentIDContains applies the Contains predicate on the "environment_id" field.
func EnvironmentIDContains(v object.ID) predicate.Variable {
	vc := string(v)
	return predicate.Variable(sql.FieldContains(FieldEnvironmentID, vc))
}

// EnvironmentIDHasPrefix applies the HasPrefix predicate on the "environment_id" field.
func EnvironmentIDHasPrefix(v object.ID) predicate.Variable {
	vc := string(v)
	return predicate.Variable(sql.FieldHasPrefix(FieldEnvironmentID, vc))
}

// EnvironmentIDHasSuffix applies the HasSuffix predicate on the "environment_id" field.
func EnvironmentIDHasSuffix(v object.ID) predicate.Variable {
	vc := string(v)
	return predicate.Variable(sql.FieldHasSuffix(FieldEnvironmentID, vc))
}

// EnvironmentIDIsNil applies the IsNil predicate on the "environment_id" field.
func EnvironmentIDIsNil() predicate.Variable {
	return predicate.Variable(sql.FieldIsNull(FieldEnvironmentID))
}

// EnvironmentIDNotNil applies the NotNil predicate on the "environment_id" field.
func EnvironmentIDNotNil() predicate.Variable {
	return predicate.Variable(sql.FieldNotNull(FieldEnvironmentID))
}

// EnvironmentIDEqualFold applies the EqualFold predicate on the "environment_id" field.
func EnvironmentIDEqualFold(v object.ID) predicate.Variable {
	vc := string(v)
	return predicate.Variable(sql.FieldEqualFold(FieldEnvironmentID, vc))
}

// EnvironmentIDContainsFold applies the ContainsFold predicate on the "environment_id" field.
func EnvironmentIDContainsFold(v object.ID) predicate.Variable {
	vc := string(v)
	return predicate.Variable(sql.FieldContainsFold(FieldEnvironmentID, vc))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Variable {
	return predicate.Variable(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Variable {
	return predicate.Variable(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Variable {
	return predicate.Variable(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Variable {
	return predicate.Variable(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Variable {
	return predicate.Variable(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Variable {
	return predicate.Variable(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Variable {
	return predicate.Variable(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Variable {
	return predicate.Variable(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Variable {
	return predicate.Variable(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Variable {
	return predicate.Variable(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Variable {
	return predicate.Variable(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Variable {
	return predicate.Variable(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Variable {
	return predicate.Variable(sql.FieldContainsFold(FieldName, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v crypto.String) predicate.Variable {
	return predicate.Variable(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v crypto.String) predicate.Variable {
	return predicate.Variable(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...crypto.String) predicate.Variable {
	return predicate.Variable(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...crypto.String) predicate.Variable {
	return predicate.Variable(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v crypto.String) predicate.Variable {
	return predicate.Variable(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v crypto.String) predicate.Variable {
	return predicate.Variable(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v crypto.String) predicate.Variable {
	return predicate.Variable(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v crypto.String) predicate.Variable {
	return predicate.Variable(sql.FieldLTE(FieldValue, v))
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v crypto.String) predicate.Variable {
	vc := string(v)
	return predicate.Variable(sql.FieldContains(FieldValue, vc))
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v crypto.String) predicate.Variable {
	vc := string(v)
	return predicate.Variable(sql.FieldHasPrefix(FieldValue, vc))
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v crypto.String) predicate.Variable {
	vc := string(v)
	return predicate.Variable(sql.FieldHasSuffix(FieldValue, vc))
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v crypto.String) predicate.Variable {
	vc := string(v)
	return predicate.Variable(sql.FieldEqualFold(FieldValue, vc))
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v crypto.String) predicate.Variable {
	vc := string(v)
	return predicate.Variable(sql.FieldContainsFold(FieldValue, vc))
}

// SensitiveEQ applies the EQ predicate on the "sensitive" field.
func SensitiveEQ(v bool) predicate.Variable {
	return predicate.Variable(sql.FieldEQ(FieldSensitive, v))
}

// SensitiveNEQ applies the NEQ predicate on the "sensitive" field.
func SensitiveNEQ(v bool) predicate.Variable {
	return predicate.Variable(sql.FieldNEQ(FieldSensitive, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Variable {
	return predicate.Variable(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Variable {
	return predicate.Variable(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Variable {
	return predicate.Variable(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Variable {
	return predicate.Variable(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Variable {
	return predicate.Variable(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Variable {
	return predicate.Variable(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Variable {
	return predicate.Variable(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Variable {
	return predicate.Variable(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Variable {
	return predicate.Variable(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Variable {
	return predicate.Variable(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Variable {
	return predicate.Variable(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Variable {
	return predicate.Variable(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Variable {
	return predicate.Variable(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Variable {
	return predicate.Variable(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Variable {
	return predicate.Variable(sql.FieldContainsFold(FieldDescription, v))
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.Variable {
	return predicate.Variable(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Project
		step.Edge.Schema = schemaConfig.Variable
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.Variable {
	return predicate.Variable(func(s *sql.Selector) {
		step := newProjectStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Project
		step.Edge.Schema = schemaConfig.Variable
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEnvironment applies the HasEdge predicate on the "environment" edge.
func HasEnvironment() predicate.Variable {
	return predicate.Variable(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EnvironmentTable, EnvironmentColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Environment
		step.Edge.Schema = schemaConfig.Variable
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEnvironmentWith applies the HasEdge predicate on the "environment" edge with a given conditions (other predicates).
func HasEnvironmentWith(preds ...predicate.Environment) predicate.Variable {
	return predicate.Variable(func(s *sql.Selector) {
		step := newEnvironmentStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Environment
		step.Edge.Schema = schemaConfig.Variable
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Variable) predicate.Variable {
	return predicate.Variable(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Variable) predicate.Variable {
	return predicate.Variable(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Variable) predicate.Variable {
	return predicate.Variable(sql.NotPredicates(p))
}
