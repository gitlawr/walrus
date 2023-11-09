package modelchange

import (
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/seal-io/walrus/pkg/dao/model/migrate"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/utils/topic"
)

// Available topics,
// which is used for model change subscription.
var (
	// Catalog is the topic for model.Catalog.
	Catalog = topic.Topic(migrate.CatalogsTable.Name)
	// Connector is the topic for model.Connector.
	Connector = topic.Topic(migrate.ConnectorsTable.Name)
	// Resource is the topic for model.Resource.
	Resource = topic.Topic(migrate.ResourcesTable.Name)
	// ResourceComponent is the topic for model.ResourceComponent.
	ResourceComponent = topic.Topic(migrate.ResourceComponentsTable.Name)
	// ResourceRevision is the topic for model.ResourceRevision.
	ResourceRevision = topic.Topic(migrate.ResourceRevisionsTable.Name)
	// Template is the topic for model.Template.
	Template = topic.Topic(migrate.TemplatesTable.Name)
	// Workflow is the topic for model.Workflow.
	Workflow = topic.Topic(migrate.WorkflowsTable.Name)
	// WorkflowExecution is the topic for model.WorkflowExecution.
	WorkflowExecution = topic.Topic(migrate.WorkflowExecutionsTable.Name)
)

// tableNameSet holds the set for interested table names,
// which should correspond to the topics above.
var tableNameSet = sets.NewString(
	// Allow subscribing from topic.
	string(Catalog),
	string(Connector),
	string(Resource),
	string(ResourceComponent),
	string(ResourceRevision),
	string(Template),
	string(Workflow),
	string(WorkflowExecution),
	// Disallow subscribing from topic.
	migrate.SettingsTable.Name,
)

// EventType indicates the type of model change event.
type EventType uint8

const (
	_EventTypeUnknown EventType = iota
	EventTypeCreate
	EventTypeUpdate
	EventTypeDelete
	_EventTypeLength
)

func (t EventType) String() string {
	switch t {
	case EventTypeCreate:
		return "create"
	case EventTypeUpdate:
		return "update"
	case EventTypeDelete:
		return "delete"
	}

	return "unknown"
}

// Event indicates the event of model change,
// includes Type and changed IDs.
type Event struct {
	Type EventType
	IDs  []object.ID
}
