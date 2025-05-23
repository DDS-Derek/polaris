// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"polaris/ent/blacklist"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BlacklistCreate is the builder for creating a Blacklist entity.
type BlacklistCreate struct {
	config
	mutation *BlacklistMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (bc *BlacklistCreate) SetType(b blacklist.Type) *BlacklistCreate {
	bc.mutation.SetType(b)
	return bc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (bc *BlacklistCreate) SetNillableType(b *blacklist.Type) *BlacklistCreate {
	if b != nil {
		bc.SetType(*b)
	}
	return bc
}

// SetTorrentHash sets the "torrent_hash" field.
func (bc *BlacklistCreate) SetTorrentHash(s string) *BlacklistCreate {
	bc.mutation.SetTorrentHash(s)
	return bc
}

// SetNillableTorrentHash sets the "torrent_hash" field if the given value is not nil.
func (bc *BlacklistCreate) SetNillableTorrentHash(s *string) *BlacklistCreate {
	if s != nil {
		bc.SetTorrentHash(*s)
	}
	return bc
}

// SetTorrentName sets the "torrent_name" field.
func (bc *BlacklistCreate) SetTorrentName(s string) *BlacklistCreate {
	bc.mutation.SetTorrentName(s)
	return bc
}

// SetNillableTorrentName sets the "torrent_name" field if the given value is not nil.
func (bc *BlacklistCreate) SetNillableTorrentName(s *string) *BlacklistCreate {
	if s != nil {
		bc.SetTorrentName(*s)
	}
	return bc
}

// SetMediaID sets the "media_id" field.
func (bc *BlacklistCreate) SetMediaID(i int) *BlacklistCreate {
	bc.mutation.SetMediaID(i)
	return bc
}

// SetNillableMediaID sets the "media_id" field if the given value is not nil.
func (bc *BlacklistCreate) SetNillableMediaID(i *int) *BlacklistCreate {
	if i != nil {
		bc.SetMediaID(*i)
	}
	return bc
}

// SetCreateTime sets the "create_time" field.
func (bc *BlacklistCreate) SetCreateTime(t time.Time) *BlacklistCreate {
	bc.mutation.SetCreateTime(t)
	return bc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (bc *BlacklistCreate) SetNillableCreateTime(t *time.Time) *BlacklistCreate {
	if t != nil {
		bc.SetCreateTime(*t)
	}
	return bc
}

// SetNotes sets the "notes" field.
func (bc *BlacklistCreate) SetNotes(s string) *BlacklistCreate {
	bc.mutation.SetNotes(s)
	return bc
}

// SetNillableNotes sets the "notes" field if the given value is not nil.
func (bc *BlacklistCreate) SetNillableNotes(s *string) *BlacklistCreate {
	if s != nil {
		bc.SetNotes(*s)
	}
	return bc
}

// Mutation returns the BlacklistMutation object of the builder.
func (bc *BlacklistCreate) Mutation() *BlacklistMutation {
	return bc.mutation
}

// Save creates the Blacklist in the database.
func (bc *BlacklistCreate) Save(ctx context.Context) (*Blacklist, error) {
	bc.defaults()
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BlacklistCreate) SaveX(ctx context.Context) *Blacklist {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BlacklistCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BlacklistCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BlacklistCreate) defaults() {
	if _, ok := bc.mutation.GetType(); !ok {
		v := blacklist.DefaultType
		bc.mutation.SetType(v)
	}
	if _, ok := bc.mutation.CreateTime(); !ok {
		v := blacklist.DefaultCreateTime()
		bc.mutation.SetCreateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BlacklistCreate) check() error {
	if _, ok := bc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Blacklist.type"`)}
	}
	if v, ok := bc.mutation.GetType(); ok {
		if err := blacklist.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Blacklist.type": %w`, err)}
		}
	}
	return nil
}

func (bc *BlacklistCreate) sqlSave(ctx context.Context) (*Blacklist, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BlacklistCreate) createSpec() (*Blacklist, *sqlgraph.CreateSpec) {
	var (
		_node = &Blacklist{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(blacklist.Table, sqlgraph.NewFieldSpec(blacklist.FieldID, field.TypeInt))
	)
	if value, ok := bc.mutation.GetType(); ok {
		_spec.SetField(blacklist.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := bc.mutation.TorrentHash(); ok {
		_spec.SetField(blacklist.FieldTorrentHash, field.TypeString, value)
		_node.TorrentHash = value
	}
	if value, ok := bc.mutation.TorrentName(); ok {
		_spec.SetField(blacklist.FieldTorrentName, field.TypeString, value)
		_node.TorrentName = value
	}
	if value, ok := bc.mutation.MediaID(); ok {
		_spec.SetField(blacklist.FieldMediaID, field.TypeInt, value)
		_node.MediaID = value
	}
	if value, ok := bc.mutation.CreateTime(); ok {
		_spec.SetField(blacklist.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := bc.mutation.Notes(); ok {
		_spec.SetField(blacklist.FieldNotes, field.TypeString, value)
		_node.Notes = value
	}
	return _node, _spec
}

// BlacklistCreateBulk is the builder for creating many Blacklist entities in bulk.
type BlacklistCreateBulk struct {
	config
	err      error
	builders []*BlacklistCreate
}

// Save creates the Blacklist entities in the database.
func (bcb *BlacklistCreateBulk) Save(ctx context.Context) ([]*Blacklist, error) {
	if bcb.err != nil {
		return nil, bcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Blacklist, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BlacklistMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BlacklistCreateBulk) SaveX(ctx context.Context) []*Blacklist {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BlacklistCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BlacklistCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
