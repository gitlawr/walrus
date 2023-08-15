// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "seal". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/seal-io/seal/pkg/dao/model/servicerevision"
	"github.com/seal-io/seal/pkg/dao/schema/intercept"
	"github.com/seal-io/seal/pkg/dao/types"
	"github.com/seal-io/seal/pkg/dao/types/crypto"
	"github.com/seal-io/seal/pkg/dao/types/object"
	"github.com/seal-io/seal/pkg/dao/types/property"
)

// ServiceRevisionCreateInput holds the creation input of the ServiceRevision entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type ServiceRevisionCreateInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to create ServiceRevision entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// Environment indicates to create ServiceRevision entity MUST under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"-"`
	// Service indicates to create ServiceRevision entity MUST under the Service route.
	Service *ServiceQueryInput `path:",inline" query:"-" json:"-"`

	// Output of the revision.
	Output string `path:"-" query:"-" json:"output"`
	// Input plan of the revision.
	InputPlan string `path:"-" query:"-" json:"inputPlan"`
	// Version of the template.
	TemplateVersion string `path:"-" query:"-" json:"templateVersion"`
	// Name of the template.
	TemplateName string `path:"-" query:"-" json:"templateName"`
	// Attributes to configure the template.
	Attributes property.Values `path:"-" query:"-" json:"attributes,omitempty"`
	// Variables of the revision.
	Variables crypto.Map[string, string] `path:"-" query:"-" json:"variables,omitempty"`
	// Type of deployer.
	DeployerType string `path:"-" query:"-" json:"deployerType,omitempty"`
	// Duration in seconds of the revision deploying.
	Duration int `path:"-" query:"-" json:"duration,omitempty"`
	// Previous provider requirement of the revision.
	PreviousRequiredProviders []types.ProviderRequirement `path:"-" query:"-" json:"previousRequiredProviders,omitempty"`
	// Tags of the revision.
	Tags []string `path:"-" query:"-" json:"tags,omitempty"`
}

// Model returns the ServiceRevision entity for creating,
// after validating.
func (srci *ServiceRevisionCreateInput) Model() *ServiceRevision {
	if srci == nil {
		return nil
	}

	_sr := &ServiceRevision{
		Output:                    srci.Output,
		InputPlan:                 srci.InputPlan,
		TemplateVersion:           srci.TemplateVersion,
		TemplateName:              srci.TemplateName,
		Attributes:                srci.Attributes,
		Variables:                 srci.Variables,
		DeployerType:              srci.DeployerType,
		Duration:                  srci.Duration,
		PreviousRequiredProviders: srci.PreviousRequiredProviders,
		Tags:                      srci.Tags,
	}

	if srci.Project != nil {
		_sr.ProjectID = srci.Project.ID
	}
	if srci.Environment != nil {
		_sr.EnvironmentID = srci.Environment.ID
	}
	if srci.Service != nil {
		_sr.ServiceID = srci.Service.ID
	}

	return _sr
}

// Validate checks the ServiceRevisionCreateInput entity.
func (srci *ServiceRevisionCreateInput) Validate() error {
	if srci == nil {
		return errors.New("nil receiver")
	}

	return srci.ValidateWith(srci.inputConfig.Context, srci.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceRevisionCreateInput entity with the given context and client set.
func (srci *ServiceRevisionCreateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when creating under the Project route.
	if srci.Project != nil {
		if err := srci.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}
	// Validate when creating under the Environment route.
	if srci.Environment != nil {
		if err := srci.Environment.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}
	// Validate when creating under the Service route.
	if srci.Service != nil {
		if err := srci.Service.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// ServiceRevisionCreateInputs holds the creation input item of the ServiceRevision entities.
type ServiceRevisionCreateInputsItem struct {
	// Output of the revision.
	Output string `path:"-" query:"-" json:"output"`
	// Input plan of the revision.
	InputPlan string `path:"-" query:"-" json:"inputPlan"`
	// Version of the template.
	TemplateVersion string `path:"-" query:"-" json:"templateVersion"`
	// Name of the template.
	TemplateName string `path:"-" query:"-" json:"templateName"`
	// Attributes to configure the template.
	Attributes property.Values `path:"-" query:"-" json:"attributes,omitempty"`
	// Variables of the revision.
	Variables crypto.Map[string, string] `path:"-" query:"-" json:"variables,omitempty"`
	// Type of deployer.
	DeployerType string `path:"-" query:"-" json:"deployerType,omitempty"`
	// Duration in seconds of the revision deploying.
	Duration int `path:"-" query:"-" json:"duration,omitempty"`
	// Previous provider requirement of the revision.
	PreviousRequiredProviders []types.ProviderRequirement `path:"-" query:"-" json:"previousRequiredProviders,omitempty"`
	// Tags of the revision.
	Tags []string `path:"-" query:"-" json:"tags,omitempty"`
}

// ValidateWith checks the ServiceRevisionCreateInputsItem entity with the given context and client set.
func (srci *ServiceRevisionCreateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// ServiceRevisionCreateInputs holds the creation input of the ServiceRevision entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ServiceRevisionCreateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to create ServiceRevision entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// Environment indicates to create ServiceRevision entity MUST under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"-"`
	// Service indicates to create ServiceRevision entity MUST under the Service route.
	Service *ServiceQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ServiceRevisionCreateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the ServiceRevision entities for creating,
// after validating.
func (srci *ServiceRevisionCreateInputs) Model() []*ServiceRevision {
	if srci == nil || len(srci.Items) == 0 {
		return nil
	}

	_srs := make([]*ServiceRevision, len(srci.Items))

	for i := range srci.Items {
		_sr := &ServiceRevision{
			Output:                    srci.Items[i].Output,
			InputPlan:                 srci.Items[i].InputPlan,
			TemplateVersion:           srci.Items[i].TemplateVersion,
			TemplateName:              srci.Items[i].TemplateName,
			Attributes:                srci.Items[i].Attributes,
			Variables:                 srci.Items[i].Variables,
			DeployerType:              srci.Items[i].DeployerType,
			Duration:                  srci.Items[i].Duration,
			PreviousRequiredProviders: srci.Items[i].PreviousRequiredProviders,
			Tags:                      srci.Items[i].Tags,
		}

		if srci.Project != nil {
			_sr.ProjectID = srci.Project.ID
		}
		if srci.Environment != nil {
			_sr.EnvironmentID = srci.Environment.ID
		}
		if srci.Service != nil {
			_sr.ServiceID = srci.Service.ID
		}

		_srs[i] = _sr
	}

	return _srs
}

// Validate checks the ServiceRevisionCreateInputs entity .
func (srci *ServiceRevisionCreateInputs) Validate() error {
	if srci == nil {
		return errors.New("nil receiver")
	}

	return srci.ValidateWith(srci.inputConfig.Context, srci.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceRevisionCreateInputs entity with the given context and client set.
func (srci *ServiceRevisionCreateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srci == nil {
		return errors.New("nil receiver")
	}

	if len(srci.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when creating under the Project route.
	if srci.Project != nil {
		if err := srci.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srci.Project = nil
			}
		}
	}
	// Validate when creating under the Environment route.
	if srci.Environment != nil {
		if err := srci.Environment.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srci.Environment = nil
			}
		}
	}
	// Validate when creating under the Service route.
	if srci.Service != nil {
		if err := srci.Service.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				srci.Service = nil
			}
		}
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

// ServiceRevisionDeleteInput holds the deletion input of the ServiceRevision entity,
// please tags with `path:",inline"` if embedding.
type ServiceRevisionDeleteInput struct {
	ServiceRevisionQueryInput `path:",inline"`
}

// ServiceRevisionDeleteInputs holds the deletion input item of the ServiceRevision entities.
type ServiceRevisionDeleteInputsItem struct {
	// ID of the ServiceRevision entity.
	ID object.ID `path:"-" query:"-" json:"id"`
}

// ServiceRevisionDeleteInputs holds the deletion input of the ServiceRevision entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ServiceRevisionDeleteInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to delete ServiceRevision entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// Environment indicates to delete ServiceRevision entity MUST under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"-"`
	// Service indicates to delete ServiceRevision entity MUST under the Service route.
	Service *ServiceQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ServiceRevisionDeleteInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the ServiceRevision entities for deleting,
// after validating.
func (srdi *ServiceRevisionDeleteInputs) Model() []*ServiceRevision {
	if srdi == nil || len(srdi.Items) == 0 {
		return nil
	}

	_srs := make([]*ServiceRevision, len(srdi.Items))
	for i := range srdi.Items {
		_srs[i] = &ServiceRevision{
			ID: srdi.Items[i].ID,
		}
	}
	return _srs
}

// IDs returns the ID list of the ServiceRevision entities for deleting,
// after validating.
func (srdi *ServiceRevisionDeleteInputs) IDs() []object.ID {
	if srdi == nil || len(srdi.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(srdi.Items))
	for i := range srdi.Items {
		ids[i] = srdi.Items[i].ID
	}
	return ids
}

// Validate checks the ServiceRevisionDeleteInputs entity.
func (srdi *ServiceRevisionDeleteInputs) Validate() error {
	if srdi == nil {
		return errors.New("nil receiver")
	}

	return srdi.ValidateWith(srdi.inputConfig.Context, srdi.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceRevisionDeleteInputs entity with the given context and client set.
func (srdi *ServiceRevisionDeleteInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srdi == nil {
		return errors.New("nil receiver")
	}

	if len(srdi.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.ServiceRevisions().Query()

	// Validate when deleting under the Project route.
	if srdi.Project != nil {
		if err := srdi.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				servicerevision.ProjectID(srdi.Project.ID))
		}
	}

	// Validate when deleting under the Environment route.
	if srdi.Environment != nil {
		if err := srdi.Environment.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				servicerevision.EnvironmentID(srdi.Environment.ID))
		}
	}

	// Validate when deleting under the Service route.
	if srdi.Service != nil {
		if err := srdi.Service.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				servicerevision.ServiceID(srdi.Service.ID))
		}
	}

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

	idsCnt, err := q.Where(servicerevision.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != cap(ids) {
		return errors.New("found unrecognized item")
	}

	return nil
}

// ServiceRevisionQueryInput holds the query input of the ServiceRevision entity,
// please tags with `path:",inline"` if embedding.
type ServiceRevisionQueryInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to query ServiceRevision entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"project"`
	// Environment indicates to query ServiceRevision entity MUST under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"environment"`
	// Service indicates to query ServiceRevision entity MUST under the Service route.
	Service *ServiceQueryInput `path:",inline" query:"-" json:"service"`

	// Refer holds the route path reference of the ServiceRevision entity.
	Refer *object.Refer `path:"servicerevision,default=" query:"-" json:"-"`
	// ID of the ServiceRevision entity.
	ID object.ID `path:"-" query:"-" json:"id"`
}

// Model returns the ServiceRevision entity for querying,
// after validating.
func (srqi *ServiceRevisionQueryInput) Model() *ServiceRevision {
	if srqi == nil {
		return nil
	}

	return &ServiceRevision{
		ID: srqi.ID,
	}
}

// Validate checks the ServiceRevisionQueryInput entity.
func (srqi *ServiceRevisionQueryInput) Validate() error {
	if srqi == nil {
		return errors.New("nil receiver")
	}

	return srqi.ValidateWith(srqi.inputConfig.Context, srqi.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceRevisionQueryInput entity with the given context and client set.
func (srqi *ServiceRevisionQueryInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srqi == nil {
		return errors.New("nil receiver")
	}

	if srqi.Refer != nil && *srqi.Refer == "" {
		return fmt.Errorf("model: %s : %w", servicerevision.Label, ErrBlankResourceRefer)
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.ServiceRevisions().Query()

	// Validate when querying under the Project route.
	if srqi.Project != nil {
		if err := srqi.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				servicerevision.ProjectID(srqi.Project.ID))
		}
	}

	// Validate when querying under the Environment route.
	if srqi.Environment != nil {
		if err := srqi.Environment.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				servicerevision.EnvironmentID(srqi.Environment.ID))
		}
	}

	// Validate when querying under the Service route.
	if srqi.Service != nil {
		if err := srqi.Service.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				servicerevision.ServiceID(srqi.Service.ID))
		}
	}

	if srqi.Refer != nil {
		if srqi.Refer.IsID() {
			q.Where(
				servicerevision.ID(srqi.Refer.ID()))
		} else {
			return errors.New("invalid identify refer of servicerevision")
		}
	} else if srqi.ID != "" {
		q.Where(
			servicerevision.ID(srqi.ID))
	} else {
		return errors.New("invalid identify of servicerevision")
	}

	q.Select(
		servicerevision.FieldID,
	)

	var e *ServiceRevision
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
			e = cv.(*ServiceRevision)
		}
	}

	srqi.ID = e.ID
	return nil
}

// ServiceRevisionQueryInputs holds the query input of the ServiceRevision entities,
// please tags with `path:",inline" query:",inline"` if embedding.
type ServiceRevisionQueryInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to query ServiceRevision entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// Environment indicates to query ServiceRevision entity MUST under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"-"`
	// Service indicates to query ServiceRevision entity MUST under the Service route.
	Service *ServiceQueryInput `path:",inline" query:"-" json:"-"`
}

// Validate checks the ServiceRevisionQueryInputs entity.
func (srqi *ServiceRevisionQueryInputs) Validate() error {
	if srqi == nil {
		return errors.New("nil receiver")
	}

	return srqi.ValidateWith(srqi.inputConfig.Context, srqi.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceRevisionQueryInputs entity with the given context and client set.
func (srqi *ServiceRevisionQueryInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srqi == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when querying under the Project route.
	if srqi.Project != nil {
		if err := srqi.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	// Validate when querying under the Environment route.
	if srqi.Environment != nil {
		if err := srqi.Environment.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	// Validate when querying under the Service route.
	if srqi.Service != nil {
		if err := srqi.Service.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// ServiceRevisionUpdateInput holds the modification input of the ServiceRevision entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type ServiceRevisionUpdateInput struct {
	ServiceRevisionQueryInput `path:",inline" query:"-" json:"-"`

	// Version of the template.
	TemplateVersion string `path:"-" query:"-" json:"templateVersion,omitempty"`
	// Attributes to configure the template.
	Attributes property.Values `path:"-" query:"-" json:"attributes,omitempty"`
	// Variables of the revision.
	Variables crypto.Map[string, string] `path:"-" query:"-" json:"variables,omitempty"`
	// Input plan of the revision.
	InputPlan string `path:"-" query:"-" json:"inputPlan,omitempty"`
	// Output of the revision.
	Output string `path:"-" query:"-" json:"output,omitempty"`
	// Type of deployer.
	DeployerType string `path:"-" query:"-" json:"deployerType,omitempty"`
	// Duration in seconds of the revision deploying.
	Duration int `path:"-" query:"-" json:"duration,omitempty"`
	// Previous provider requirement of the revision.
	PreviousRequiredProviders []types.ProviderRequirement `path:"-" query:"-" json:"previousRequiredProviders,omitempty"`
	// Tags of the revision.
	Tags []string `path:"-" query:"-" json:"tags,omitempty"`
}

// Model returns the ServiceRevision entity for modifying,
// after validating.
func (srui *ServiceRevisionUpdateInput) Model() *ServiceRevision {
	if srui == nil {
		return nil
	}

	_sr := &ServiceRevision{
		ID:                        srui.ID,
		TemplateVersion:           srui.TemplateVersion,
		Attributes:                srui.Attributes,
		Variables:                 srui.Variables,
		InputPlan:                 srui.InputPlan,
		Output:                    srui.Output,
		DeployerType:              srui.DeployerType,
		Duration:                  srui.Duration,
		PreviousRequiredProviders: srui.PreviousRequiredProviders,
		Tags:                      srui.Tags,
	}

	return _sr
}

// Validate checks the ServiceRevisionUpdateInput entity.
func (srui *ServiceRevisionUpdateInput) Validate() error {
	if srui == nil {
		return errors.New("nil receiver")
	}

	return srui.ValidateWith(srui.inputConfig.Context, srui.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceRevisionUpdateInput entity with the given context and client set.
func (srui *ServiceRevisionUpdateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := srui.ServiceRevisionQueryInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	return nil
}

// ServiceRevisionUpdateInputs holds the modification input item of the ServiceRevision entities.
type ServiceRevisionUpdateInputsItem struct {
	// ID of the ServiceRevision entity.
	ID object.ID `path:"-" query:"-" json:"id"`

	// Version of the template.
	TemplateVersion string `path:"-" query:"-" json:"templateVersion"`
	// Attributes to configure the template.
	Attributes property.Values `path:"-" query:"-" json:"attributes,omitempty"`
	// Variables of the revision.
	Variables crypto.Map[string, string] `path:"-" query:"-" json:"variables"`
	// Input plan of the revision.
	InputPlan string `path:"-" query:"-" json:"inputPlan"`
	// Output of the revision.
	Output string `path:"-" query:"-" json:"output"`
	// Type of deployer.
	DeployerType string `path:"-" query:"-" json:"deployerType"`
	// Duration in seconds of the revision deploying.
	Duration int `path:"-" query:"-" json:"duration"`
	// Previous provider requirement of the revision.
	PreviousRequiredProviders []types.ProviderRequirement `path:"-" query:"-" json:"previousRequiredProviders"`
	// Tags of the revision.
	Tags []string `path:"-" query:"-" json:"tags"`
}

// ValidateWith checks the ServiceRevisionUpdateInputsItem entity with the given context and client set.
func (srui *ServiceRevisionUpdateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srui == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// ServiceRevisionUpdateInputs holds the modification input of the ServiceRevision entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ServiceRevisionUpdateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to update ServiceRevision entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// Environment indicates to update ServiceRevision entity MUST under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"-"`
	// Service indicates to update ServiceRevision entity MUST under the Service route.
	Service *ServiceQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ServiceRevisionUpdateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the ServiceRevision entities for modifying,
// after validating.
func (srui *ServiceRevisionUpdateInputs) Model() []*ServiceRevision {
	if srui == nil || len(srui.Items) == 0 {
		return nil
	}

	_srs := make([]*ServiceRevision, len(srui.Items))

	for i := range srui.Items {
		_sr := &ServiceRevision{
			ID:                        srui.Items[i].ID,
			TemplateVersion:           srui.Items[i].TemplateVersion,
			Attributes:                srui.Items[i].Attributes,
			Variables:                 srui.Items[i].Variables,
			InputPlan:                 srui.Items[i].InputPlan,
			Output:                    srui.Items[i].Output,
			DeployerType:              srui.Items[i].DeployerType,
			Duration:                  srui.Items[i].Duration,
			PreviousRequiredProviders: srui.Items[i].PreviousRequiredProviders,
			Tags:                      srui.Items[i].Tags,
		}

		_srs[i] = _sr
	}

	return _srs
}

// IDs returns the ID list of the ServiceRevision entities for modifying,
// after validating.
func (srui *ServiceRevisionUpdateInputs) IDs() []object.ID {
	if srui == nil || len(srui.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(srui.Items))
	for i := range srui.Items {
		ids[i] = srui.Items[i].ID
	}
	return ids
}

// Validate checks the ServiceRevisionUpdateInputs entity.
func (srui *ServiceRevisionUpdateInputs) Validate() error {
	if srui == nil {
		return errors.New("nil receiver")
	}

	return srui.ValidateWith(srui.inputConfig.Context, srui.inputConfig.Client, nil)
}

// ValidateWith checks the ServiceRevisionUpdateInputs entity with the given context and client set.
func (srui *ServiceRevisionUpdateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if srui == nil {
		return errors.New("nil receiver")
	}

	if len(srui.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.ServiceRevisions().Query()

	// Validate when updating under the Project route.
	if srui.Project != nil {
		if err := srui.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				servicerevision.ProjectID(srui.Project.ID))
		}
	}

	// Validate when updating under the Environment route.
	if srui.Environment != nil {
		if err := srui.Environment.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				servicerevision.EnvironmentID(srui.Environment.ID))
		}
	}

	// Validate when updating under the Service route.
	if srui.Service != nil {
		if err := srui.Service.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				servicerevision.ServiceID(srui.Service.ID))
		}
	}

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

	idsCnt, err := q.Where(servicerevision.IDIn(ids...)).
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

// ServiceRevisionOutput holds the output of the ServiceRevision entity.
type ServiceRevisionOutput struct {
	ID                        object.ID                   `json:"id,omitempty"`
	CreateTime                *time.Time                  `json:"createTime,omitempty"`
	Status                    string                      `json:"status,omitempty"`
	StatusMessage             string                      `json:"statusMessage,omitempty"`
	TemplateName              string                      `json:"templateName,omitempty"`
	TemplateVersion           string                      `json:"templateVersion,omitempty"`
	Attributes                property.Values             `json:"attributes,omitempty"`
	Variables                 crypto.Map[string, string]  `json:"variables,omitempty"`
	DeployerType              string                      `json:"deployerType,omitempty"`
	Duration                  int                         `json:"duration,omitempty"`
	PreviousRequiredProviders []types.ProviderRequirement `json:"previousRequiredProviders,omitempty"`
	Tags                      []string                    `json:"tags,omitempty"`

	Project     *ProjectOutput     `json:"project,omitempty"`
	Environment *EnvironmentOutput `json:"environment,omitempty"`
	Service     *ServiceOutput     `json:"service,omitempty"`
}

// View returns the output of ServiceRevision entity.
func (_sr *ServiceRevision) View() *ServiceRevisionOutput {
	return ExposeServiceRevision(_sr)
}

// View returns the output of ServiceRevision entities.
func (_srs ServiceRevisions) View() []*ServiceRevisionOutput {
	return ExposeServiceRevisions(_srs)
}

// ExposeServiceRevision converts the ServiceRevision to ServiceRevisionOutput.
func ExposeServiceRevision(_sr *ServiceRevision) *ServiceRevisionOutput {
	if _sr == nil {
		return nil
	}

	sro := &ServiceRevisionOutput{
		ID:                        _sr.ID,
		CreateTime:                _sr.CreateTime,
		Status:                    _sr.Status,
		StatusMessage:             _sr.StatusMessage,
		TemplateName:              _sr.TemplateName,
		TemplateVersion:           _sr.TemplateVersion,
		Attributes:                _sr.Attributes,
		Variables:                 _sr.Variables,
		DeployerType:              _sr.DeployerType,
		Duration:                  _sr.Duration,
		PreviousRequiredProviders: _sr.PreviousRequiredProviders,
		Tags:                      _sr.Tags,
	}

	if _sr.Edges.Project != nil {
		sro.Project = ExposeProject(_sr.Edges.Project)
	} else if _sr.ProjectID != "" {
		sro.Project = &ProjectOutput{
			ID: _sr.ProjectID,
		}
	}
	if _sr.Edges.Environment != nil {
		sro.Environment = ExposeEnvironment(_sr.Edges.Environment)
	} else if _sr.EnvironmentID != "" {
		sro.Environment = &EnvironmentOutput{
			ID: _sr.EnvironmentID,
		}
	}
	if _sr.Edges.Service != nil {
		sro.Service = ExposeService(_sr.Edges.Service)
	} else if _sr.ServiceID != "" {
		sro.Service = &ServiceOutput{
			ID: _sr.ServiceID,
		}
	}
	return sro
}

// ExposeServiceRevisions converts the ServiceRevision slice to ServiceRevisionOutput pointer slice.
func ExposeServiceRevisions(_srs []*ServiceRevision) []*ServiceRevisionOutput {
	if len(_srs) == 0 {
		return nil
	}

	sros := make([]*ServiceRevisionOutput, len(_srs))
	for i := range _srs {
		sros[i] = ExposeServiceRevision(_srs[i])
	}
	return sros
}
