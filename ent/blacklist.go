// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"polaris/ent/blacklist"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Blacklist is the model entity for the Blacklist schema.
type Blacklist struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type blacklist.Type `json:"type,omitempty"`
	// TorrentHash holds the value of the "torrent_hash" field.
	TorrentHash string `json:"torrent_hash,omitempty"`
	// TorrentName holds the value of the "torrent_name" field.
	TorrentName string `json:"torrent_name,omitempty"`
	// MediaID holds the value of the "media_id" field.
	MediaID int `json:"media_id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// Notes holds the value of the "notes" field.
	Notes        string `json:"notes,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Blacklist) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case blacklist.FieldID, blacklist.FieldMediaID:
			values[i] = new(sql.NullInt64)
		case blacklist.FieldType, blacklist.FieldTorrentHash, blacklist.FieldTorrentName, blacklist.FieldNotes:
			values[i] = new(sql.NullString)
		case blacklist.FieldCreateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Blacklist fields.
func (b *Blacklist) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case blacklist.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case blacklist.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				b.Type = blacklist.Type(value.String)
			}
		case blacklist.FieldTorrentHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field torrent_hash", values[i])
			} else if value.Valid {
				b.TorrentHash = value.String
			}
		case blacklist.FieldTorrentName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field torrent_name", values[i])
			} else if value.Valid {
				b.TorrentName = value.String
			}
		case blacklist.FieldMediaID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field media_id", values[i])
			} else if value.Valid {
				b.MediaID = int(value.Int64)
			}
		case blacklist.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				b.CreateTime = value.Time
			}
		case blacklist.FieldNotes:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field notes", values[i])
			} else if value.Valid {
				b.Notes = value.String
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Blacklist.
// This includes values selected through modifiers, order, etc.
func (b *Blacklist) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// Update returns a builder for updating this Blacklist.
// Note that you need to call Blacklist.Unwrap() before calling this method if this Blacklist
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Blacklist) Update() *BlacklistUpdateOne {
	return NewBlacklistClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Blacklist entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Blacklist) Unwrap() *Blacklist {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Blacklist is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Blacklist) String() string {
	var builder strings.Builder
	builder.WriteString("Blacklist(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", b.Type))
	builder.WriteString(", ")
	builder.WriteString("torrent_hash=")
	builder.WriteString(b.TorrentHash)
	builder.WriteString(", ")
	builder.WriteString("torrent_name=")
	builder.WriteString(b.TorrentName)
	builder.WriteString(", ")
	builder.WriteString("media_id=")
	builder.WriteString(fmt.Sprintf("%v", b.MediaID))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(b.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("notes=")
	builder.WriteString(b.Notes)
	builder.WriteByte(')')
	return builder.String()
}

// Blacklists is a parsable slice of Blacklist.
type Blacklists []*Blacklist
