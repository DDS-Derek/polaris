package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Episode holds the schema definition for the Epidodes entity.
type Episode struct {
	ent.Schema
}

// Fields of the Episode.
func (Episode) Fields() []ent.Field {
	return []ent.Field{
		field.Int("media_id").Optional(),
		field.Int("season_number").StructTag("json:\"season_number\""),
		field.Int("episode_number").StructTag("json:\"episode_number\""),
		field.String("title"),
		field.String("overview"),
		field.String("air_date"),
		field.Enum("status").Values("missing", "downloading", "downloaded").Default("missing"),
		field.Bool("monitored").Default(false).StructTag("json:\"monitored\""), //whether this episode is monitored
		field.String("target_file").Optional(),
		field.Time("create_time").Optional().Default(time.Now).Immutable(),
	}
}

// Edges of the Episode.
func (Episode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("media", Media.Type).
			Ref("episodes").
			Unique().
			Field("media_id"),
	}

}
