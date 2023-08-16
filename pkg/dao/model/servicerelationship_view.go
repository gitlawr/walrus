// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "seal". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/seal-io/walrus/pkg/dao/model/servicerelationship"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// ServiceRelationshipCreateInput holds the creation input of the ServiceRelationship entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type ServiceRelationshipCreateInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Dependency specifies full inserting the new Service entity of the ServiceRelationship entity.
	Dependency *ServiceQueryInput `uri:"-" query:"-" json:"dependency"`
}

// Model returns the ServiceRelationship entity for creating,
// after validating.
func (srci *ServiceRelationshipCreateInput) Model() *ServiceRelationship {
	if srci == nil {
		return nil
	}

	_sr := &ServiceRelationship{}

	if srci.Dependency != nil {
		_sr.DependencyID = srci.Dependency.ID
	}
	return _sr
}

// Validate checks the ServiceRelationshipCreateInput entity.
func (srci *ServiceRelationshipCreateInput) Validate() error {
	if srci == nil {
		return errors.New("nil receiver")
	}

	return srci.ValidateWith(srci.inputConfig.Context, srci.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceRelationshipCreateInput entity with the given context and client set.
func (srci *ServiceRelationshipCreateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	if srci.Dependency != nil {
		if err := srci.Dependency.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srci.Dependency = nil
			}
		}
	}

	return nil
}

// ServiceRelationshipCreateInputs holds the creation input item of the ServiceRelationship entities.
type ServiceRelationshipCreateInputsItem struct {

	// Dependency specifies full inserting the new Service entity.
	Dependency *ServiceQueryInput `uri:"-" query:"-" json:"dependency"`
}

// ValidateWith checks the ServiceRelationshipCreateInputsItem entity with the given context and client set.
func (srci *ServiceRelationshipCreateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	if srci.Dependency != nil {
		if err := srci.Dependency.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srci.Dependency = nil
			}
		}
	}

	return nil
}

// ServiceRelationshipCreateInputs holds the creation input of the ServiceRelationship entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ServiceRelationshipCreateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ServiceRelationshipCreateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the ServiceRelationship entities for creating,
// after validating.
func (srci *ServiceRelationshipCreateInputs) Model() []*ServiceRelationship {
	if srci == nil || len(srci.Items) == 0 {
		return nil
	}

	_srs := make([]*ServiceRelationship, len(srci.Items))

	for i := range srci.Items {
		_sr := &ServiceRelationship{}

		if srci.Items[i].Dependency != nil {
			_sr.DependencyID = srci.Items[i].Dependency.ID
		}

		_srs[i] = _sr
	}

	return _srs
}

// Validate checks the ServiceRelationshipCreateInputs entity .
func (srci *ServiceRelationshipCreateInputs) Validate() error {
	if srci == nil {
		return errors.New("nil receiver")
	}

	return srci.ValidateWith(srci.inputConfig.Context, srci.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceRelationshipCreateInputs entity with the given context and client set.
func (srci *ServiceRelationshipCreateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srci == nil {
		return errors.New("nil receiver")
	}

	if len(srci.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	for i := range srci.Items {
		if srci.Items[i] == nil {
			continue
		}

		if err := srci.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// ServiceRelationshipDeleteInput holds the deletion input of the ServiceRelationship entity,
// please tags with `path:",inline"` if embedding.
type ServiceRelationshipDeleteInput struct {
	ServiceRelationshipQueryInput `path:",inline"`
}

// ServiceRelationshipDeleteInputs holds the deletion input item of the ServiceRelationship entities.
type ServiceRelationshipDeleteInputsItem struct {
	// ID of the ServiceRelationship entity.
	ID object.ID `path:"-" query:"-" json:"id"`
}

// ServiceRelationshipDeleteInputs holds the deletion input of the ServiceRelationship entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ServiceRelationshipDeleteInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ServiceRelationshipDeleteInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the ServiceRelationship entities for deleting,
// after validating.
func (srdi *ServiceRelationshipDeleteInputs) Model() []*ServiceRelationship {
	if srdi == nil || len(srdi.Items) == 0 {
		return nil
	}

	_srs := make([]*ServiceRelationship, len(srdi.Items))
	for i := range srdi.Items {
		_srs[i] = &ServiceRelationship{
			ID: srdi.Items[i].ID,
		}
	}
	return _srs
}

// IDs returns the ID list of the ServiceRelationship entities for deleting,
// after validating.
func (srdi *ServiceRelationshipDeleteInputs) IDs() []object.ID {
	if srdi == nil || len(srdi.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(srdi.Items))
	for i := range srdi.Items {
		ids[i] = srdi.Items[i].ID
	}
	return ids
}

// Validate checks the ServiceRelationshipDeleteInputs entity.
func (srdi *ServiceRelationshipDeleteInputs) Validate() error {
	if srdi == nil {
		return errors.New("nil receiver")
	}

	return srdi.ValidateWith(srdi.inputConfig.Context, srdi.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceRelationshipDeleteInputs entity with the given context and client set.
func (srdi *ServiceRelationshipDeleteInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srdi == nil {
		return errors.New("nil receiver")
	}

	if len(srdi.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.ServiceRelationships().Query()

	ids := make([]object.ID, 0, len(srdi.Items))

	for i := range srdi.Items {
		if srdi.Items[i] == nil {
			return errors.New("nil item")
		}

		if srdi.Items[i].ID != "" {
			ids = append(ids, srdi.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	if len(ids) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	idsCnt, err := q.Where(servicerelationship.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != cap(ids) {
		return errors.New("found unrecognized item")
	}

	return nil
}

// ServiceRelationshipQueryInput holds the query input of the ServiceRelationship entity,
// please tags with `path:",inline"` if embedding.
type ServiceRelationshipQueryInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Refer holds the route path reference of the ServiceRelationship entity.
	Refer *object.Refer `path:"servicerelationship,default=" query:"-" json:"-"`
	// ID of the ServiceRelationship entity.
	ID object.ID `path:"-" query:"-" json:"id"`
}

// Model returns the ServiceRelationship entity for querying,
// after validating.
func (srqi *ServiceRelationshipQueryInput) Model() *ServiceRelationship {
	if srqi == nil {
		return nil
	}

	return &ServiceRelationship{
		ID: srqi.ID,
	}
}

// Validate checks the ServiceRelationshipQueryInput entity.
func (srqi *ServiceRelationshipQueryInput) Validate() error {
	if srqi == nil {
		return errors.New("nil receiver")
	}

	return srqi.ValidateWith(srqi.inputConfig.Context, srqi.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceRelationshipQueryInput entity with the given context and client set.
func (srqi *ServiceRelationshipQueryInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srqi == nil {
		return errors.New("nil receiver")
	}

	if srqi.Refer != nil && *srqi.Refer == "" {
		return fmt.Errorf("model: %s : %w", servicerelationship.Label, ErrBlankResourceRefer)
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.ServiceRelationships().Query()

	if srqi.Refer != nil {
		if srqi.Refer.IsID() {
			q.Where(
				servicerelationship.ID(srqi.Refer.ID()))
		} else {
			return errors.New("invalid identify refer of servicerelationship")
		}
	} else if srqi.ID != "" {
		q.Where(
			servicerelationship.ID(srqi.ID))
	} else {
		return errors.New("invalid identify of servicerelationship")
	}

	q.Select(
		servicerelationship.FieldID,
	)

	var e *ServiceRelationship
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
			e = cv.(*ServiceRelationship)
		}
	}

	srqi.ID = e.ID
	return nil
}

// ServiceRelationshipQueryInputs holds the query input of the ServiceRelationship entities,
// please tags with `path:",inline" query:",inline"` if embedding.
type ServiceRelationshipQueryInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`
}

// Validate checks the ServiceRelationshipQueryInputs entity.
func (srqi *ServiceRelationshipQueryInputs) Validate() error {
	if srqi == nil {
		return errors.New("nil receiver")
	}

	return srqi.ValidateWith(srqi.inputConfig.Context, srqi.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceRelationshipQueryInputs entity with the given context and client set.
func (srqi *ServiceRelationshipQueryInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srqi == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// ServiceRelationshipUpdateInput holds the modification input of the ServiceRelationship entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type ServiceRelationshipUpdateInput struct {
	ServiceRelationshipQueryInput `path:",inline" query:"-" json:"-"`

	// Dependency indicates replacing the stale Service entity.
	Dependency *ServiceQueryInput `uri:"-" query:"-" json:"dependency"`
}

// Model returns the ServiceRelationship entity for modifying,
// after validating.
func (srui *ServiceRelationshipUpdateInput) Model() *ServiceRelationship {
	if srui == nil {
		return nil
	}

	_sr := &ServiceRelationship{
		ID: srui.ID,
	}

	if srui.Dependency != nil {
		_sr.DependencyID = srui.Dependency.ID
	}
	return _sr
}

// Validate checks the ServiceRelationshipUpdateInput entity.
func (srui *ServiceRelationshipUpdateInput) Validate() error {
	if srui == nil {
		return errors.New("nil receiver")
	}

	return srui.ValidateWith(srui.inputConfig.Context, srui.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceRelationshipUpdateInput entity with the given context and client set.
func (srui *ServiceRelationshipUpdateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := srui.ServiceRelationshipQueryInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	if srui.Dependency != nil {
		if err := srui.Dependency.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srui.Dependency = nil
			}
		}
	}

	return nil
}

// ServiceRelationshipUpdateInputs holds the modification input item of the ServiceRelationship entities.
type ServiceRelationshipUpdateInputsItem struct {
	// ID of the ServiceRelationship entity.
	ID object.ID `path:"-" query:"-" json:"id"`

	// Dependency indicates replacing the stale Service entity.
	Dependency *ServiceQueryInput `uri:"-" query:"-" json:"dependency"`
}

// ValidateWith checks the ServiceRelationshipUpdateInputsItem entity with the given context and client set.
func (srui *ServiceRelationshipUpdateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srui == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	if srui.Dependency != nil {
		if err := srui.Dependency.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srui.Dependency = nil
			}
		}
	}

	return nil
}

// ServiceRelationshipUpdateInputs holds the modification input of the ServiceRelationship entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ServiceRelationshipUpdateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ServiceRelationshipUpdateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the ServiceRelationship entities for modifying,
// after validating.
func (srui *ServiceRelationshipUpdateInputs) Model() []*ServiceRelationship {
	if srui == nil || len(srui.Items) == 0 {
		return nil
	}

	_srs := make([]*ServiceRelationship, len(srui.Items))

	for i := range srui.Items {
		_sr := &ServiceRelationship{
			ID: srui.Items[i].ID,
		}

		if srui.Items[i].Dependency != nil {
			_sr.DependencyID = srui.Items[i].Dependency.ID
		}

		_srs[i] = _sr
	}

	return _srs
}

// IDs returns the ID list of the ServiceRelationship entities for modifying,
// after validating.
func (srui *ServiceRelationshipUpdateInputs) IDs() []object.ID {
	if srui == nil || len(srui.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(srui.Items))
	for i := range srui.Items {
		ids[i] = srui.Items[i].ID
	}
	return ids
}

// Validate checks the ServiceRelationshipUpdateInputs entity.
func (srui *ServiceRelationshipUpdateInputs) Validate() error {
	if srui == nil {
		return errors.New("nil receiver")
	}

	return srui.ValidateWith(srui.inputConfig.Context, srui.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceRelationshipUpdateInputs entity with the given context and client set.
func (srui *ServiceRelationshipUpdateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srui == nil {
		return errors.New("nil receiver")
	}

	if len(srui.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.ServiceRelationships().Query()

	ids := make([]object.ID, 0, len(srui.Items))

	for i := range srui.Items {
		if srui.Items[i] == nil {
			return errors.New("nil item")
		}

		if srui.Items[i].ID != "" {
			ids = append(ids, srui.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	if len(ids) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	idsCnt, err := q.Where(servicerelationship.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range srui.Items {
		if srui.Items[i] == nil {
			continue
		}

		if err := srui.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// ServiceRelationshipOutput holds the output of the ServiceRelationship entity.
type ServiceRelationshipOutput struct {
	ID         object.ID   `json:"id,omitempty"`
	CreateTime *time.Time  `json:"createTime,omitempty"`
	Path       []object.ID `json:"path,omitempty"`
	Type       string      `json:"type,omitempty"`

	Dependency *ServiceOutput `json:"dependency,omitempty"`
}

// View returns the output of ServiceRelationship entity.
func (_sr *ServiceRelationship) View() *ServiceRelationshipOutput {
	return ExposeServiceRelationship(_sr)
}

// View returns the output of ServiceRelationship entities.
func (_srs ServiceRelationships) View() []*ServiceRelationshipOutput {
	return ExposeServiceRelationships(_srs)
}

// ExposeServiceRelationship converts the ServiceRelationship to ServiceRelationshipOutput.
func ExposeServiceRelationship(_sr *ServiceRelationship) *ServiceRelationshipOutput {
	if _sr == nil {
		return nil
	}

	sro := &ServiceRelationshipOutput{
		ID:         _sr.ID,
		CreateTime: _sr.CreateTime,
		Path:       _sr.Path,
		Type:       _sr.Type,
	}

	if _sr.Edges.Dependency != nil {
		sro.Dependency = ExposeService(_sr.Edges.Dependency)
	} else if _sr.DependencyID != "" {
		sro.Dependency = &ServiceOutput{
			ID: _sr.DependencyID,
		}
	}
	return sro
}

// ExposeServiceRelationships converts the ServiceRelationship slice to ServiceRelationshipOutput pointer slice.
func ExposeServiceRelationships(_srs []*ServiceRelationship) []*ServiceRelationshipOutput {
	if len(_srs) == 0 {
		return nil
	}

	sros := make([]*ServiceRelationshipOutput, len(_srs))
	for i := range _srs {
		sros[i] = ExposeServiceRelationship(_srs[i])
	}
	return sros
}
