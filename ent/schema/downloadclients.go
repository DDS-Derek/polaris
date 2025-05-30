package schema

import (
	"errors"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// DownloadClients holds the schema definition for the DownloadClients entity.
type DownloadClients struct {
	ent.Schema
}

// Fields of the DownloadClients.
func (DownloadClients) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("enable"),
		field.String("name"),
		field.Enum("implementation").Values("transmission", "qbittorrent", "buildin"),
		field.String("url"),
		field.String("user").Default(""),
		field.String("password").Default(""),
		field.String("settings").Default(""),
		field.Int("priority1").Default(1).Validate(func(i int) error {
			if i > 50 {
				return errors.ErrUnsupported
			}
			if i <= 0 {
				return errors.ErrUnsupported
			}
			return nil
		}),
		field.Bool("use_nat_traversal").Optional().Default(false).Comment("use stun server to do nat traversal, enable download client to do uploading successfully"),
		field.Bool("remove_completed_downloads").Default(true),
		field.Bool("remove_failed_downloads").Default(true),
		field.String("tags").Default(""),
		field.Time("create_time").Optional().Default(time.Now).Immutable(),
	}
}

// Edges of the DownloadClients.
func (DownloadClients) Edges() []ent.Edge {
	return nil
}
