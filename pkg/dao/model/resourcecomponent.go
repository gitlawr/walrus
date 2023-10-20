// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"

	"github.com/seal-io/walrus/pkg/dao/model/connector"
	"github.com/seal-io/walrus/pkg/dao/model/environment"
	"github.com/seal-io/walrus/pkg/dao/model/project"
	"github.com/seal-io/walrus/pkg/dao/model/resource"
	"github.com/seal-io/walrus/pkg/dao/model/resourcecomponent"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/utils/json"
)

// ResourceComponent is the model entity for the ResourceComponent schema.
type ResourceComponent struct {
	config `json:"-"`
	// ID of the ent.
	ID object.ID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// ID of the project to belong.
	ProjectID object.ID `json:"project_id,omitempty"`
	// ID of the environment to which the component belongs.
	EnvironmentID object.ID `json:"environment_id,omitempty"`
	// ID of the resource to which the component belongs.
	ResourceID object.ID `json:"resource_id,omitempty"`
	// ID of the connector to which the component deploys.
	ConnectorID object.ID `json:"connector_id,omitempty"`
	// ID of the parent component.
	CompositionID object.ID `json:"composition_id,omitempty"`
	// ID of the parent class of the component realization.
	ClassID object.ID `json:"class_id,omitempty"`
	// Mode that manages the generated component, it is the management way of the deployer to the component, which provides by deployer.
	Mode string `json:"mode,omitempty"`
	// Type of the generated component, it is the type of the resource which the deployer observes, which provides by deployer.
	Type string `json:"type,omitempty"`
	// Name of the generated component, it is the real identifier of the component, which provides by deployer.
	Name string `json:"name,omitempty"`
	// Type of deployer.
	DeployerType string `json:"deployer_type,omitempty"`
	// Shape of the component, it can be class or instance shape.
	Shape string `json:"shape,omitempty"`
	// Status of the component.
	Status types.ResourceComponentStatus `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ResourceComponentQuery when eager-loading is set.
	Edges        ResourceComponentEdges `json:"edges,omitempty"`
	selectValues sql.SelectValues

	// Keys of the component.
	// Keys does not store in the database.
	Keys *types.ResourceComponentOperationKeys `json:"keys,omitempty"`
}

// ResourceComponentEdges holds the relations/edges for other nodes in the graph.
type ResourceComponentEdges struct {
	// Project to which the component belongs.
	Project *Project `json:"project,omitempty"`
	// Environment to which the component deploys.
	Environment *Environment `json:"environment,omitempty"`
	// Resource to which the component belongs.
	Resource *Resource `json:"resource,omitempty"`
	// Connector to which the component deploys.
	Connector *Connector `json:"connector,omitempty"`
	// Composition holds the value of the composition edge.
	Composition *ResourceComponent `json:"composition,omitempty"`
	// Components that makes up the resource component.
	Components []*ResourceComponent `json:"components,omitempty"`
	// Class holds the value of the class edge.
	Class *ResourceComponent `json:"class,omitempty"`
	// Instances that realizes the resource component.
	Instances []*ResourceComponent `json:"instances,omitempty"`
	// Dependencies holds the value of the dependencies edge.
	Dependencies []*ResourceComponentRelationship `json:"dependencies,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [9]bool
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResourceComponentEdges) ProjectOrErr() (*Project, error) {
	if e.loadedTypes[0] {
		if e.Project == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: project.Label}
		}
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// EnvironmentOrErr returns the Environment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResourceComponentEdges) EnvironmentOrErr() (*Environment, error) {
	if e.loadedTypes[1] {
		if e.Environment == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: environment.Label}
		}
		return e.Environment, nil
	}
	return nil, &NotLoadedError{edge: "environment"}
}

// ResourceOrErr returns the Resource value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResourceComponentEdges) ResourceOrErr() (*Resource, error) {
	if e.loadedTypes[2] {
		if e.Resource == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: resource.Label}
		}
		return e.Resource, nil
	}
	return nil, &NotLoadedError{edge: "resource"}
}

// ConnectorOrErr returns the Connector value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResourceComponentEdges) ConnectorOrErr() (*Connector, error) {
	if e.loadedTypes[3] {
		if e.Connector == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: connector.Label}
		}
		return e.Connector, nil
	}
	return nil, &NotLoadedError{edge: "connector"}
}

// CompositionOrErr returns the Composition value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResourceComponentEdges) CompositionOrErr() (*ResourceComponent, error) {
	if e.loadedTypes[4] {
		if e.Composition == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: resourcecomponent.Label}
		}
		return e.Composition, nil
	}
	return nil, &NotLoadedError{edge: "composition"}
}

// ComponentsOrErr returns the Components value or an error if the edge
// was not loaded in eager-loading.
func (e ResourceComponentEdges) ComponentsOrErr() ([]*ResourceComponent, error) {
	if e.loadedTypes[5] {
		return e.Components, nil
	}
	return nil, &NotLoadedError{edge: "components"}
}

// ClassOrErr returns the Class value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResourceComponentEdges) ClassOrErr() (*ResourceComponent, error) {
	if e.loadedTypes[6] {
		if e.Class == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: resourcecomponent.Label}
		}
		return e.Class, nil
	}
	return nil, &NotLoadedError{edge: "class"}
}

// InstancesOrErr returns the Instances value or an error if the edge
// was not loaded in eager-loading.
func (e ResourceComponentEdges) InstancesOrErr() ([]*ResourceComponent, error) {
	if e.loadedTypes[7] {
		return e.Instances, nil
	}
	return nil, &NotLoadedError{edge: "instances"}
}

// DependenciesOrErr returns the Dependencies value or an error if the edge
// was not loaded in eager-loading.
func (e ResourceComponentEdges) DependenciesOrErr() ([]*ResourceComponentRelationship, error) {
	if e.loadedTypes[8] {
		return e.Dependencies, nil
	}
	return nil, &NotLoadedError{edge: "dependencies"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ResourceComponent) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case resourcecomponent.FieldStatus:
			values[i] = new([]byte)
		case resourcecomponent.FieldID, resourcecomponent.FieldProjectID, resourcecomponent.FieldEnvironmentID, resourcecomponent.FieldResourceID, resourcecomponent.FieldConnectorID, resourcecomponent.FieldCompositionID, resourcecomponent.FieldClassID:
			values[i] = new(object.ID)
		case resourcecomponent.FieldMode, resourcecomponent.FieldType, resourcecomponent.FieldName, resourcecomponent.FieldDeployerType, resourcecomponent.FieldShape:
			values[i] = new(sql.NullString)
		case resourcecomponent.FieldCreateTime, resourcecomponent.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ResourceComponent fields.
func (rc *ResourceComponent) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case resourcecomponent.FieldID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				rc.ID = *value
			}
		case resourcecomponent.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				rc.CreateTime = new(time.Time)
				*rc.CreateTime = value.Time
			}
		case resourcecomponent.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				rc.UpdateTime = new(time.Time)
				*rc.UpdateTime = value.Time
			}
		case resourcecomponent.FieldProjectID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field project_id", values[i])
			} else if value != nil {
				rc.ProjectID = *value
			}
		case resourcecomponent.FieldEnvironmentID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field environment_id", values[i])
			} else if value != nil {
				rc.EnvironmentID = *value
			}
		case resourcecomponent.FieldResourceID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field resource_id", values[i])
			} else if value != nil {
				rc.ResourceID = *value
			}
		case resourcecomponent.FieldConnectorID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field connector_id", values[i])
			} else if value != nil {
				rc.ConnectorID = *value
			}
		case resourcecomponent.FieldCompositionID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field composition_id", values[i])
			} else if value != nil {
				rc.CompositionID = *value
			}
		case resourcecomponent.FieldClassID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field class_id", values[i])
			} else if value != nil {
				rc.ClassID = *value
			}
		case resourcecomponent.FieldMode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mode", values[i])
			} else if value.Valid {
				rc.Mode = value.String
			}
		case resourcecomponent.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				rc.Type = value.String
			}
		case resourcecomponent.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				rc.Name = value.String
			}
		case resourcecomponent.FieldDeployerType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deployer_type", values[i])
			} else if value.Valid {
				rc.DeployerType = value.String
			}
		case resourcecomponent.FieldShape:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field shape", values[i])
			} else if value.Valid {
				rc.Shape = value.String
			}
		case resourcecomponent.FieldStatus:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &rc.Status); err != nil {
					return fmt.Errorf("unmarshal field status: %w", err)
				}
			}
		default:
			rc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ResourceComponent.
// This includes values selected through modifiers, order, etc.
func (rc *ResourceComponent) Value(name string) (ent.Value, error) {
	return rc.selectValues.Get(name)
}

// QueryProject queries the "project" edge of the ResourceComponent entity.
func (rc *ResourceComponent) QueryProject() *ProjectQuery {
	return NewResourceComponentClient(rc.config).QueryProject(rc)
}

// QueryEnvironment queries the "environment" edge of the ResourceComponent entity.
func (rc *ResourceComponent) QueryEnvironment() *EnvironmentQuery {
	return NewResourceComponentClient(rc.config).QueryEnvironment(rc)
}

// QueryResource queries the "resource" edge of the ResourceComponent entity.
func (rc *ResourceComponent) QueryResource() *ResourceQuery {
	return NewResourceComponentClient(rc.config).QueryResource(rc)
}

// QueryConnector queries the "connector" edge of the ResourceComponent entity.
func (rc *ResourceComponent) QueryConnector() *ConnectorQuery {
	return NewResourceComponentClient(rc.config).QueryConnector(rc)
}

// QueryComposition queries the "composition" edge of the ResourceComponent entity.
func (rc *ResourceComponent) QueryComposition() *ResourceComponentQuery {
	return NewResourceComponentClient(rc.config).QueryComposition(rc)
}

// QueryComponents queries the "components" edge of the ResourceComponent entity.
func (rc *ResourceComponent) QueryComponents() *ResourceComponentQuery {
	return NewResourceComponentClient(rc.config).QueryComponents(rc)
}

// QueryClass queries the "class" edge of the ResourceComponent entity.
func (rc *ResourceComponent) QueryClass() *ResourceComponentQuery {
	return NewResourceComponentClient(rc.config).QueryClass(rc)
}

// QueryInstances queries the "instances" edge of the ResourceComponent entity.
func (rc *ResourceComponent) QueryInstances() *ResourceComponentQuery {
	return NewResourceComponentClient(rc.config).QueryInstances(rc)
}

// QueryDependencies queries the "dependencies" edge of the ResourceComponent entity.
func (rc *ResourceComponent) QueryDependencies() *ResourceComponentRelationshipQuery {
	return NewResourceComponentClient(rc.config).QueryDependencies(rc)
}

// Update returns a builder for updating this ResourceComponent.
// Note that you need to call ResourceComponent.Unwrap() before calling this method if this ResourceComponent
// was returned from a transaction, and the transaction was committed or rolled back.
func (rc *ResourceComponent) Update() *ResourceComponentUpdateOne {
	return NewResourceComponentClient(rc.config).UpdateOne(rc)
}

// Unwrap unwraps the ResourceComponent entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rc *ResourceComponent) Unwrap() *ResourceComponent {
	_tx, ok := rc.config.driver.(*txDriver)
	if !ok {
		panic("model: ResourceComponent is not a transactional entity")
	}
	rc.config.driver = _tx.drv
	return rc
}

// String implements the fmt.Stringer.
func (rc *ResourceComponent) String() string {
	var builder strings.Builder
	builder.WriteString("ResourceComponent(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rc.ID))
	if v := rc.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := rc.UpdateTime; v != nil {
		builder.WriteString("update_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("project_id=")
	builder.WriteString(fmt.Sprintf("%v", rc.ProjectID))
	builder.WriteString(", ")
	builder.WriteString("environment_id=")
	builder.WriteString(fmt.Sprintf("%v", rc.EnvironmentID))
	builder.WriteString(", ")
	builder.WriteString("resource_id=")
	builder.WriteString(fmt.Sprintf("%v", rc.ResourceID))
	builder.WriteString(", ")
	builder.WriteString("connector_id=")
	builder.WriteString(fmt.Sprintf("%v", rc.ConnectorID))
	builder.WriteString(", ")
	builder.WriteString("composition_id=")
	builder.WriteString(fmt.Sprintf("%v", rc.CompositionID))
	builder.WriteString(", ")
	builder.WriteString("class_id=")
	builder.WriteString(fmt.Sprintf("%v", rc.ClassID))
	builder.WriteString(", ")
	builder.WriteString("mode=")
	builder.WriteString(rc.Mode)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(rc.Type)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(rc.Name)
	builder.WriteString(", ")
	builder.WriteString("deployer_type=")
	builder.WriteString(rc.DeployerType)
	builder.WriteString(", ")
	builder.WriteString("shape=")
	builder.WriteString(rc.Shape)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", rc.Status))
	builder.WriteByte(')')
	return builder.String()
}

// ResourceComponents is a parsable slice of ResourceComponent.
type ResourceComponents []*ResourceComponent