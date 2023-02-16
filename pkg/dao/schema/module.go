package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/seal/pkg/dao/schema/mixin"
)

type Module struct {
	schema
}

func (Module) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Status{},
		mixin.Time{},
	}
}

func (Module) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Comment("It is also the name of the module.").
			Unique().
			Immutable(),
		field.String("description").
			Comment("Description of the module.").
			Optional(),
		field.JSON("labels", map[string]string{}).
			Comment("Labels of the module.").
			Optional(),
		field.String("source").
			Comment("Source of the module."),
		field.String("version").
			Comment("Version of the module."),
		field.JSON("inputSchema", map[string]interface{}{}).
			Comment("Input schema of the module.").
			Optional(),
		field.JSON("outputSchema", map[string]interface{}{}).
			Comment("Output schema of the module.").
			Optional(),
	}
}

func (Module) Edges() []ent.Edge {
	return []ent.Edge{
		// applications *-* modules.
		edge.From("application", Application.Type).
			Ref("modules").
			Comment("Applications to which the module configures.").
			Through("applicationModuleRelationships", ApplicationModuleRelationship.Type),
	}
}
