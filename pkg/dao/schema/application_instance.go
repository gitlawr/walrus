package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/seal-io/seal/pkg/dao/schema/mixin"
	"github.com/seal-io/seal/pkg/dao/types/id"
)

type ApplicationInstance struct {
	ent.Schema
}

func (ApplicationInstance) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ID{},
		mixin.Status{},
		mixin.Time{},
	}
}

func (ApplicationInstance) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("applicationID", "environmentID", "name").
			Unique(),
	}
}

func (ApplicationInstance) Fields() []ent.Field {
	return []ent.Field{
		id.Field("applicationID").
			Comment("ID of the application to which the instance belongs.").
			NotEmpty().
			Immutable(),
		id.Field("environmentID").
			Comment("ID of the environment to which the instance deploys.").
			NotEmpty().
			Immutable(),
		field.String("name").
			Comment("Name of the instance.").
			NotEmpty().
			Immutable(),
		field.JSON("variables", map[string]interface{}{}).
			Comment("Variables of the instance.").
			Optional(),
	}
}

func (ApplicationInstance) Edges() []ent.Edge {
	return []ent.Edge{
		// application 1-* application instances.
		edge.From("application", Application.Type).
			Ref("instances").
			Field("applicationID").
			Comment("Application to which the instance belongs.").
			Unique().
			Required().
			Immutable(),
		// environment 1-* application instances.
		edge.From("environment", Environment.Type).
			Ref("instances").
			Field("environmentID").
			Comment("Environment to which the instance belongs.").
			Unique().
			Required().
			Immutable(),
		// application instance 1-* application revisions.
		edge.To("revisions", ApplicationRevision.Type).
			Comment("Application revisions that belong to this instance.").
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		// application instance 1-* application resources.
		edge.To("resources", ApplicationResource.Type).
			Comment("Application resources that belong to the instance.").
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}