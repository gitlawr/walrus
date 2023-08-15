// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "seal". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/seal-io/seal/pkg/dao/model/subjectrolerelationship"
	"github.com/seal-io/seal/pkg/dao/schema/intercept"
	"github.com/seal-io/seal/pkg/dao/types/object"
)

// SubjectRoleRelationshipCreateInput holds the creation input of the SubjectRoleRelationship entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type SubjectRoleRelationshipCreateInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to create SubjectRoleRelationship entity CAN under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`

	// Subject specifies full inserting the new Subject entity of the SubjectRoleRelationship entity.
	Subject *SubjectQueryInput `uri:"-" query:"-" json:"subject"`
	// Role specifies full inserting the new Role entity of the SubjectRoleRelationship entity.
	Role *RoleQueryInput `uri:"-" query:"-" json:"role"`
}

// Model returns the SubjectRoleRelationship entity for creating,
// after validating.
func (srrci *SubjectRoleRelationshipCreateInput) Model() *SubjectRoleRelationship {
	if srrci == nil {
		return nil
	}

	_srr := &SubjectRoleRelationship{}

	if srrci.Project != nil {
		_srr.ProjectID = srrci.Project.ID
	}

	if srrci.Subject != nil {
		_srr.SubjectID = srrci.Subject.ID
	}
	if srrci.Role != nil {
		_srr.RoleID = srrci.Role.ID
	}
	return _srr
}

// Validate checks the SubjectRoleRelationshipCreateInput entity.
func (srrci *SubjectRoleRelationshipCreateInput) Validate() error {
	if srrci == nil {
		return errors.New("nil receiver")
	}

	return srrci.ValidateWith(srrci.inputConfig.Context, srrci.inputConfig.Client, nil)
}

// ValidateWith checks the SubjectRoleRelationshipCreateInput entity with the given context and client set.
func (srrci *SubjectRoleRelationshipCreateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srrci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when creating under the Project route.
	if srrci.Project != nil {
		if err := srrci.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srrci.Project = nil
			}
		}
	}

	if srrci.Subject != nil {
		if err := srrci.Subject.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srrci.Subject = nil
			}
		}
	}

	if srrci.Role != nil {
		if err := srrci.Role.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srrci.Role = nil
			}
		}
	}

	return nil
}

// SubjectRoleRelationshipCreateInputs holds the creation input item of the SubjectRoleRelationship entities.
type SubjectRoleRelationshipCreateInputsItem struct {

	// Subject specifies full inserting the new Subject entity.
	Subject *SubjectQueryInput `uri:"-" query:"-" json:"subject"`
	// Role specifies full inserting the new Role entity.
	Role *RoleQueryInput `uri:"-" query:"-" json:"role"`
}

// ValidateWith checks the SubjectRoleRelationshipCreateInputsItem entity with the given context and client set.
func (srrci *SubjectRoleRelationshipCreateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srrci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	if srrci.Subject != nil {
		if err := srrci.Subject.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srrci.Subject = nil
			}
		}
	}

	if srrci.Role != nil {
		if err := srrci.Role.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srrci.Role = nil
			}
		}
	}

	return nil
}

// SubjectRoleRelationshipCreateInputs holds the creation input of the SubjectRoleRelationship entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type SubjectRoleRelationshipCreateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to create SubjectRoleRelationship entity CAN under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*SubjectRoleRelationshipCreateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the SubjectRoleRelationship entities for creating,
// after validating.
func (srrci *SubjectRoleRelationshipCreateInputs) Model() []*SubjectRoleRelationship {
	if srrci == nil || len(srrci.Items) == 0 {
		return nil
	}

	_srrs := make([]*SubjectRoleRelationship, len(srrci.Items))

	for i := range srrci.Items {
		_srr := &SubjectRoleRelationship{}

		if srrci.Project != nil {
			_srr.ProjectID = srrci.Project.ID
		}

		if srrci.Items[i].Subject != nil {
			_srr.SubjectID = srrci.Items[i].Subject.ID
		}
		if srrci.Items[i].Role != nil {
			_srr.RoleID = srrci.Items[i].Role.ID
		}

		_srrs[i] = _srr
	}

	return _srrs
}

// Validate checks the SubjectRoleRelationshipCreateInputs entity .
func (srrci *SubjectRoleRelationshipCreateInputs) Validate() error {
	if srrci == nil {
		return errors.New("nil receiver")
	}

	return srrci.ValidateWith(srrci.inputConfig.Context, srrci.inputConfig.Client, nil)
}

// ValidateWith checks the SubjectRoleRelationshipCreateInputs entity with the given context and client set.
func (srrci *SubjectRoleRelationshipCreateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srrci == nil {
		return errors.New("nil receiver")
	}

	if len(srrci.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when creating under the Project route.
	if srrci.Project != nil {
		if err := srrci.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srrci.Project = nil
			}
		}
	}

	for i := range srrci.Items {
		if srrci.Items[i] == nil {
			continue
		}

		if err := srrci.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// SubjectRoleRelationshipDeleteInput holds the deletion input of the SubjectRoleRelationship entity,
// please tags with `path:",inline"` if embedding.
type SubjectRoleRelationshipDeleteInput struct {
	SubjectRoleRelationshipQueryInput `path:",inline"`
}

// SubjectRoleRelationshipDeleteInputs holds the deletion input item of the SubjectRoleRelationship entities.
type SubjectRoleRelationshipDeleteInputsItem struct {
	// ID of the SubjectRoleRelationship entity.
	ID object.ID `path:"-" query:"-" json:"id"`
}

// SubjectRoleRelationshipDeleteInputs holds the deletion input of the SubjectRoleRelationship entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type SubjectRoleRelationshipDeleteInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to delete SubjectRoleRelationship entity CAN under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*SubjectRoleRelationshipDeleteInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the SubjectRoleRelationship entities for deleting,
// after validating.
func (srrdi *SubjectRoleRelationshipDeleteInputs) Model() []*SubjectRoleRelationship {
	if srrdi == nil || len(srrdi.Items) == 0 {
		return nil
	}

	_srrs := make([]*SubjectRoleRelationship, len(srrdi.Items))
	for i := range srrdi.Items {
		_srrs[i] = &SubjectRoleRelationship{
			ID: srrdi.Items[i].ID,
		}
	}
	return _srrs
}

// IDs returns the ID list of the SubjectRoleRelationship entities for deleting,
// after validating.
func (srrdi *SubjectRoleRelationshipDeleteInputs) IDs() []object.ID {
	if srrdi == nil || len(srrdi.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(srrdi.Items))
	for i := range srrdi.Items {
		ids[i] = srrdi.Items[i].ID
	}
	return ids
}

// Validate checks the SubjectRoleRelationshipDeleteInputs entity.
func (srrdi *SubjectRoleRelationshipDeleteInputs) Validate() error {
	if srrdi == nil {
		return errors.New("nil receiver")
	}

	return srrdi.ValidateWith(srrdi.inputConfig.Context, srrdi.inputConfig.Client, nil)
}

// ValidateWith checks the SubjectRoleRelationshipDeleteInputs entity with the given context and client set.
func (srrdi *SubjectRoleRelationshipDeleteInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srrdi == nil {
		return errors.New("nil receiver")
	}

	if len(srrdi.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.SubjectRoleRelationships().Query()

	// Validate when deleting under the Project route.
	if srrdi.Project != nil {
		if err := srrdi.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srrdi.Project = nil
				q.Where(
					subjectrolerelationship.ProjectIDIsNil())
			}
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				subjectrolerelationship.ProjectID(srrdi.Project.ID))
		}
	} else {
		q.Where(
			subjectrolerelationship.ProjectIDIsNil())
	}

	ids := make([]object.ID, 0, len(srrdi.Items))

	for i := range srrdi.Items {
		if srrdi.Items[i] == nil {
			return errors.New("nil item")
		}

		if srrdi.Items[i].ID != "" {
			ids = append(ids, srrdi.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	if len(ids) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	idsCnt, err := q.Where(subjectrolerelationship.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != cap(ids) {
		return errors.New("found unrecognized item")
	}

	return nil
}

// SubjectRoleRelationshipQueryInput holds the query input of the SubjectRoleRelationship entity,
// please tags with `path:",inline"` if embedding.
type SubjectRoleRelationshipQueryInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to query SubjectRoleRelationship entity CAN under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"project,omitempty"`

	// Refer holds the route path reference of the SubjectRoleRelationship entity.
	Refer *object.Refer `path:"subjectrolerelationship,default=" query:"-" json:"-"`
	// ID of the SubjectRoleRelationship entity.
	ID object.ID `path:"-" query:"-" json:"id"`
}

// Model returns the SubjectRoleRelationship entity for querying,
// after validating.
func (srrqi *SubjectRoleRelationshipQueryInput) Model() *SubjectRoleRelationship {
	if srrqi == nil {
		return nil
	}

	return &SubjectRoleRelationship{
		ID: srrqi.ID,
	}
}

// Validate checks the SubjectRoleRelationshipQueryInput entity.
func (srrqi *SubjectRoleRelationshipQueryInput) Validate() error {
	if srrqi == nil {
		return errors.New("nil receiver")
	}

	return srrqi.ValidateWith(srrqi.inputConfig.Context, srrqi.inputConfig.Client, nil)
}

// ValidateWith checks the SubjectRoleRelationshipQueryInput entity with the given context and client set.
func (srrqi *SubjectRoleRelationshipQueryInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srrqi == nil {
		return errors.New("nil receiver")
	}

	if srrqi.Refer != nil && *srrqi.Refer == "" {
		return fmt.Errorf("model: %s : %w", subjectrolerelationship.Label, ErrBlankResourceRefer)
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.SubjectRoleRelationships().Query()

	// Validate when querying under the Project route.
	if srrqi.Project != nil {
		if err := srrqi.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srrqi.Project = nil
				q.Where(
					subjectrolerelationship.ProjectIDIsNil())
			}
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				subjectrolerelationship.ProjectID(srrqi.Project.ID))
		}
	} else {
		q.Where(
			subjectrolerelationship.ProjectIDIsNil())
	}

	if srrqi.Refer != nil {
		if srrqi.Refer.IsID() {
			q.Where(
				subjectrolerelationship.ID(srrqi.Refer.ID()))
		} else {
			return errors.New("invalid identify refer of subjectrolerelationship")
		}
	} else if srrqi.ID != "" {
		q.Where(
			subjectrolerelationship.ID(srrqi.ID))
	} else {
		return errors.New("invalid identify of subjectrolerelationship")
	}

	q.Select(
		subjectrolerelationship.FieldID,
	)

	var e *SubjectRoleRelationship
	{
		// Get cache from previous validation.
		queryStmt, queryArgs := q.sqlQuery(setContextOp(ctx, q.ctx, "cache")).Query()
		ck := fmt.Sprintf("stmt=%v, args=%v", queryStmt, queryArgs)
		if cv, existed := cache[ck]; !existed {
			var err error
			e, err = q.Only(ctx)
			if err != nil {
				return err
			}

			// Set cache for other validation.
			cache[ck] = e
		} else {
			e = cv.(*SubjectRoleRelationship)
		}
	}

	srrqi.ID = e.ID
	return nil
}

// SubjectRoleRelationshipQueryInputs holds the query input of the SubjectRoleRelationship entities,
// please tags with `path:",inline" query:",inline"` if embedding.
type SubjectRoleRelationshipQueryInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to query SubjectRoleRelationship entity CAN under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
}

// Validate checks the SubjectRoleRelationshipQueryInputs entity.
func (srrqi *SubjectRoleRelationshipQueryInputs) Validate() error {
	if srrqi == nil {
		return errors.New("nil receiver")
	}

	return srrqi.ValidateWith(srrqi.inputConfig.Context, srrqi.inputConfig.Client, nil)
}

// ValidateWith checks the SubjectRoleRelationshipQueryInputs entity with the given context and client set.
func (srrqi *SubjectRoleRelationshipQueryInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srrqi == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when querying under the Project route.
	if srrqi.Project != nil {
		if err := srrqi.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srrqi.Project = nil
			}
		}
	}

	return nil
}

// SubjectRoleRelationshipUpdateInput holds the modification input of the SubjectRoleRelationship entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type SubjectRoleRelationshipUpdateInput struct {
	SubjectRoleRelationshipQueryInput `path:",inline" query:"-" json:"-"`

	// Subject indicates replacing the stale Subject entity.
	Subject *SubjectQueryInput `uri:"-" query:"-" json:"subject"`
	// Role indicates replacing the stale Role entity.
	Role *RoleQueryInput `uri:"-" query:"-" json:"role"`
}

// Model returns the SubjectRoleRelationship entity for modifying,
// after validating.
func (srrui *SubjectRoleRelationshipUpdateInput) Model() *SubjectRoleRelationship {
	if srrui == nil {
		return nil
	}

	_srr := &SubjectRoleRelationship{
		ID: srrui.ID,
	}

	if srrui.Subject != nil {
		_srr.SubjectID = srrui.Subject.ID
	}
	if srrui.Role != nil {
		_srr.RoleID = srrui.Role.ID
	}
	return _srr
}

// Validate checks the SubjectRoleRelationshipUpdateInput entity.
func (srrui *SubjectRoleRelationshipUpdateInput) Validate() error {
	if srrui == nil {
		return errors.New("nil receiver")
	}

	return srrui.ValidateWith(srrui.inputConfig.Context, srrui.inputConfig.Client, nil)
}

// ValidateWith checks the SubjectRoleRelationshipUpdateInput entity with the given context and client set.
func (srrui *SubjectRoleRelationshipUpdateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := srrui.SubjectRoleRelationshipQueryInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	if srrui.Subject != nil {
		if err := srrui.Subject.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srrui.Subject = nil
			}
		}
	}

	if srrui.Role != nil {
		if err := srrui.Role.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srrui.Role = nil
			}
		}
	}

	return nil
}

// SubjectRoleRelationshipUpdateInputs holds the modification input item of the SubjectRoleRelationship entities.
type SubjectRoleRelationshipUpdateInputsItem struct {
	// ID of the SubjectRoleRelationship entity.
	ID object.ID `path:"-" query:"-" json:"id"`

	// Subject indicates replacing the stale Subject entity.
	Subject *SubjectQueryInput `uri:"-" query:"-" json:"subject"`
	// Role indicates replacing the stale Role entity.
	Role *RoleQueryInput `uri:"-" query:"-" json:"role"`
}

// ValidateWith checks the SubjectRoleRelationshipUpdateInputsItem entity with the given context and client set.
func (srrui *SubjectRoleRelationshipUpdateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srrui == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	if srrui.Subject != nil {
		if err := srrui.Subject.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srrui.Subject = nil
			}
		}
	}

	if srrui.Role != nil {
		if err := srrui.Role.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srrui.Role = nil
			}
		}
	}

	return nil
}

// SubjectRoleRelationshipUpdateInputs holds the modification input of the SubjectRoleRelationship entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type SubjectRoleRelationshipUpdateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to update SubjectRoleRelationship entity CAN under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*SubjectRoleRelationshipUpdateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the SubjectRoleRelationship entities for modifying,
// after validating.
func (srrui *SubjectRoleRelationshipUpdateInputs) Model() []*SubjectRoleRelationship {
	if srrui == nil || len(srrui.Items) == 0 {
		return nil
	}

	_srrs := make([]*SubjectRoleRelationship, len(srrui.Items))

	for i := range srrui.Items {
		_srr := &SubjectRoleRelationship{
			ID: srrui.Items[i].ID,
		}

		if srrui.Items[i].Subject != nil {
			_srr.SubjectID = srrui.Items[i].Subject.ID
		}
		if srrui.Items[i].Role != nil {
			_srr.RoleID = srrui.Items[i].Role.ID
		}

		_srrs[i] = _srr
	}

	return _srrs
}

// IDs returns the ID list of the SubjectRoleRelationship entities for modifying,
// after validating.
func (srrui *SubjectRoleRelationshipUpdateInputs) IDs() []object.ID {
	if srrui == nil || len(srrui.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(srrui.Items))
	for i := range srrui.Items {
		ids[i] = srrui.Items[i].ID
	}
	return ids
}

// Validate checks the SubjectRoleRelationshipUpdateInputs entity.
func (srrui *SubjectRoleRelationshipUpdateInputs) Validate() error {
	if srrui == nil {
		return errors.New("nil receiver")
	}

	return srrui.ValidateWith(srrui.inputConfig.Context, srrui.inputConfig.Client, nil)
}

// ValidateWith checks the SubjectRoleRelationshipUpdateInputs entity with the given context and client set.
func (srrui *SubjectRoleRelationshipUpdateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srrui == nil {
		return errors.New("nil receiver")
	}

	if len(srrui.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.SubjectRoleRelationships().Query()

	// Validate when updating under the Project route.
	if srrui.Project != nil {
		if err := srrui.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srrui.Project = nil
				q.Where(
					subjectrolerelationship.ProjectIDIsNil())
			}
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				subjectrolerelationship.ProjectID(srrui.Project.ID))
		}
	} else {
		q.Where(
			subjectrolerelationship.ProjectIDIsNil())
	}

	ids := make([]object.ID, 0, len(srrui.Items))

	for i := range srrui.Items {
		if srrui.Items[i] == nil {
			return errors.New("nil item")
		}

		if srrui.Items[i].ID != "" {
			ids = append(ids, srrui.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	if len(ids) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	idsCnt, err := q.Where(subjectrolerelationship.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range srrui.Items {
		if srrui.Items[i] == nil {
			continue
		}

		if err := srrui.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// SubjectRoleRelationshipOutput holds the output of the SubjectRoleRelationship entity.
type SubjectRoleRelationshipOutput struct {
	ID         object.ID  `json:"id,omitempty"`
	CreateTime *time.Time `json:"createTime,omitempty"`

	Project *ProjectOutput `json:"project,omitempty"`
	Subject *SubjectOutput `json:"subject,omitempty"`
	Role    *RoleOutput    `json:"role,omitempty"`
}

// View returns the output of SubjectRoleRelationship entity.
func (_srr *SubjectRoleRelationship) View() *SubjectRoleRelationshipOutput {
	return ExposeSubjectRoleRelationship(_srr)
}

// View returns the output of SubjectRoleRelationship entities.
func (_srrs SubjectRoleRelationships) View() []*SubjectRoleRelationshipOutput {
	return ExposeSubjectRoleRelationships(_srrs)
}

// ExposeSubjectRoleRelationship converts the SubjectRoleRelationship to SubjectRoleRelationshipOutput.
func ExposeSubjectRoleRelationship(_srr *SubjectRoleRelationship) *SubjectRoleRelationshipOutput {
	if _srr == nil {
		return nil
	}

	srro := &SubjectRoleRelationshipOutput{
		ID:         _srr.ID,
		CreateTime: _srr.CreateTime,
	}

	if _srr.Edges.Project != nil {
		srro.Project = ExposeProject(_srr.Edges.Project)
	} else if _srr.ProjectID != "" {
		srro.Project = &ProjectOutput{
			ID: _srr.ProjectID,
		}
	}
	if _srr.Edges.Subject != nil {
		srro.Subject = ExposeSubject(_srr.Edges.Subject)
	} else if _srr.SubjectID != "" {
		srro.Subject = &SubjectOutput{
			ID: _srr.SubjectID,
		}
	}
	if _srr.Edges.Role != nil {
		srro.Role = ExposeRole(_srr.Edges.Role)
	} else if _srr.RoleID != "" {
		srro.Role = &RoleOutput{
			ID: _srr.RoleID,
		}
	}
	return srro
}

// ExposeSubjectRoleRelationships converts the SubjectRoleRelationship slice to SubjectRoleRelationshipOutput pointer slice.
func ExposeSubjectRoleRelationships(_srrs []*SubjectRoleRelationship) []*SubjectRoleRelationshipOutput {
	if len(_srrs) == 0 {
		return nil
	}

	srros := make([]*SubjectRoleRelationshipOutput, len(_srrs))
	for i := range _srrs {
		srros[i] = ExposeSubjectRoleRelationship(_srrs[i])
	}
	return srros
}
