// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import "time"

// ClusterCostQueryInput is the input for the ClusterCost query.
type ClusterCostQueryInput struct {
	// ID holds the value of the "id" field.
	ID int `uri:"id,omitempty" json:"id,omitempty"`
}

// Model converts the ClusterCostQueryInput to ClusterCost.
func (in ClusterCostQueryInput) Model() *ClusterCost {
	return &ClusterCost{
		ID: in.ID,
	}
}

// ClusterCostCreateInput is the input for the ClusterCost creation.
type ClusterCostCreateInput struct {
	// Usage start time for current cost
	StartTime time.Time `json:"startTime,omitempty"`
	// Usage end time for current cost
	EndTime time.Time `json:"endTime,omitempty"`
	// Usage minutes from start time to end time
	Minutes float64 `json:"minutes,omitempty"`
	// Cluster name for current cost
	ClusterName string `json:"clusterName"`
	// Cost number
	TotalCost float64 `json:"totalCost,omitempty"`
	// Cost currency
	Currency int `json:"currency,omitempty"`
	// CPU cost for current cost
	CpuCost float64 `json:"cpuCost,omitempty"`
	// GPU cost for current cost
	GpuCost float64 `json:"gpuCost,omitempty"`
	// Ram cost for current cost
	RamCost float64 `json:"ramCost,omitempty"`
	// Storage cost for current cost
	StorageCost float64 `json:"storageCost,omitempty"`
	// Allocation cost for current cost
	AllocationCost float64 `json:"allocationCost,omitempty"`
	// Idle cost for current cost
	IdleCost float64 `json:"idleCost,omitempty"`
	// Storage cost for current cost
	ManagementCost float64 `json:"managementCost,omitempty"`
	// Connector current cost linked
	Connector ConnectorQueryInput `json:"connector"`
}

// Model converts the ClusterCostCreateInput to ClusterCost.
func (in ClusterCostCreateInput) Model() *ClusterCost {
	var entity = &ClusterCost{
		StartTime:      in.StartTime,
		EndTime:        in.EndTime,
		Minutes:        in.Minutes,
		ClusterName:    in.ClusterName,
		TotalCost:      in.TotalCost,
		Currency:       in.Currency,
		CpuCost:        in.CpuCost,
		GpuCost:        in.GpuCost,
		RamCost:        in.RamCost,
		StorageCost:    in.StorageCost,
		AllocationCost: in.AllocationCost,
		IdleCost:       in.IdleCost,
		ManagementCost: in.ManagementCost,
	}
	entity.ConnectorID = in.Connector.ID
	return entity
}

// ClusterCostUpdateInput is the input for the ClusterCost modification.
type ClusterCostUpdateInput struct {
	// ID holds the value of the "id" field.
	ID int `uri:"id" json:"-"`
	// Cost number
	TotalCost float64 `json:"totalCost,omitempty"`
	// Cost currency
	Currency int `json:"currency,omitempty"`
	// CPU cost for current cost
	CpuCost float64 `json:"cpuCost,omitempty"`
	// GPU cost for current cost
	GpuCost float64 `json:"gpuCost,omitempty"`
	// Ram cost for current cost
	RamCost float64 `json:"ramCost,omitempty"`
	// Storage cost for current cost
	StorageCost float64 `json:"storageCost,omitempty"`
	// Allocation cost for current cost
	AllocationCost float64 `json:"allocationCost,omitempty"`
	// Idle cost for current cost
	IdleCost float64 `json:"idleCost,omitempty"`
	// Storage cost for current cost
	ManagementCost float64 `json:"managementCost,omitempty"`
}

// Model converts the ClusterCostUpdateInput to ClusterCost.
func (in ClusterCostUpdateInput) Model() *ClusterCost {
	var entity = &ClusterCost{
		ID:             in.ID,
		TotalCost:      in.TotalCost,
		Currency:       in.Currency,
		CpuCost:        in.CpuCost,
		GpuCost:        in.GpuCost,
		RamCost:        in.RamCost,
		StorageCost:    in.StorageCost,
		AllocationCost: in.AllocationCost,
		IdleCost:       in.IdleCost,
		ManagementCost: in.ManagementCost,
	}
	return entity
}

// ClusterCostOutput is the output for the ClusterCost.
type ClusterCostOutput struct {
	// ID holds the value of the "id" field.
	ID int `json:"id,omitempty"`
	// Usage start time for current cost
	StartTime time.Time `json:"startTime,omitempty"`
	// Usage end time for current cost
	EndTime time.Time `json:"endTime,omitempty"`
	// Usage minutes from start time to end time
	Minutes float64 `json:"minutes,omitempty"`
	// Cluster name for current cost
	ClusterName string `json:"clusterName,omitempty"`
	// Cost number
	TotalCost float64 `json:"totalCost,omitempty"`
	// Cost currency
	Currency int `json:"currency,omitempty"`
	// CPU cost for current cost
	CpuCost float64 `json:"cpuCost,omitempty"`
	// GPU cost for current cost
	GpuCost float64 `json:"gpuCost,omitempty"`
	// Ram cost for current cost
	RamCost float64 `json:"ramCost,omitempty"`
	// Storage cost for current cost
	StorageCost float64 `json:"storageCost,omitempty"`
	// Allocation cost for current cost
	AllocationCost float64 `json:"allocationCost,omitempty"`
	// Idle cost for current cost
	IdleCost float64 `json:"idleCost,omitempty"`
	// Storage cost for current cost
	ManagementCost float64 `json:"managementCost,omitempty"`
	// Connector current cost linked
	Connector *ConnectorOutput `json:"connector,omitempty"`
}

// ExposeClusterCost converts the ClusterCost to ClusterCostOutput.
func ExposeClusterCost(in *ClusterCost) *ClusterCostOutput {
	if in == nil {
		return nil
	}
	var entity = &ClusterCostOutput{
		ID:             in.ID,
		StartTime:      in.StartTime,
		EndTime:        in.EndTime,
		Minutes:        in.Minutes,
		ClusterName:    in.ClusterName,
		TotalCost:      in.TotalCost,
		Currency:       in.Currency,
		CpuCost:        in.CpuCost,
		GpuCost:        in.GpuCost,
		RamCost:        in.RamCost,
		StorageCost:    in.StorageCost,
		AllocationCost: in.AllocationCost,
		IdleCost:       in.IdleCost,
		ManagementCost: in.ManagementCost,
		Connector:      ExposeConnector(in.Edges.Connector),
	}
	if entity.Connector == nil {
		entity.Connector = &ConnectorOutput{}
	}
	entity.Connector.ID = in.ConnectorID
	return entity
}

// ExposeClusterCosts converts the ClusterCost slice to ClusterCostOutput pointer slice.
func ExposeClusterCosts(in []*ClusterCost) []*ClusterCostOutput {
	var out = make([]*ClusterCostOutput, 0, len(in))
	for i := 0; i < len(in); i++ {
		var o = ExposeClusterCost(in[i])
		if o == nil {
			continue
		}
		out = append(out, o)
	}
	if len(out) == 0 {
		return nil
	}
	return out
}