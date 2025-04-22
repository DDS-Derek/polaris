package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Storage holds the schema definition for the Storage entity.
type Storage struct {
	ent.Schema
}

// Fields of the Storage.
func (Storage) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.Enum("implementation").Values("webdav", "local", "alist"),
		field.String("tv_path").Optional(),
		field.String("movie_path").Optional(),
		field.String("settings").Optional(),
		field.Bool("deleted").Default(false),
		field.Bool("default").Default(false),
		field.Time("create_time").Optional().Default(time.Now).Immutable(),
	}
}

// Edges of the Storage.
func (Storage) Edges() []ent.Edge {
	return nil
}
