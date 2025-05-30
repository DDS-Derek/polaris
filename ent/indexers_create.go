// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"polaris/ent/indexers"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// IndexersCreate is the builder for creating a Indexers entity.
type IndexersCreate struct {
	config
	mutation *IndexersMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ic *IndexersCreate) SetName(s string) *IndexersCreate {
	ic.mutation.SetName(s)
	return ic
}

// SetImplementation sets the "implementation" field.
func (ic *IndexersCreate) SetImplementation(s string) *IndexersCreate {
	ic.mutation.SetImplementation(s)
	return ic
}

// SetSettings sets the "settings" field.
func (ic *IndexersCreate) SetSettings(s string) *IndexersCreate {
	ic.mutation.SetSettings(s)
	return ic
}

// SetNillableSettings sets the "settings" field if the given value is not nil.
func (ic *IndexersCreate) SetNillableSettings(s *string) *IndexersCreate {
	if s != nil {
		ic.SetSettings(*s)
	}
	return ic
}

// SetEnableRss sets the "enable_rss" field.
func (ic *IndexersCreate) SetEnableRss(b bool) *IndexersCreate {
	ic.mutation.SetEnableRss(b)
	return ic
}

// SetNillableEnableRss sets the "enable_rss" field if the given value is not nil.
func (ic *IndexersCreate) SetNillableEnableRss(b *bool) *IndexersCreate {
	if b != nil {
		ic.SetEnableRss(*b)
	}
	return ic
}

// SetPriority sets the "priority" field.
func (ic *IndexersCreate) SetPriority(i int) *IndexersCreate {
	ic.mutation.SetPriority(i)
	return ic
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (ic *IndexersCreate) SetNillablePriority(i *int) *IndexersCreate {
	if i != nil {
		ic.SetPriority(*i)
	}
	return ic
}

// SetSeedRatio sets the "seed_ratio" field.
func (ic *IndexersCreate) SetSeedRatio(f float32) *IndexersCreate {
	ic.mutation.SetSeedRatio(f)
	return ic
}

// SetNillableSeedRatio sets the "seed_ratio" field if the given value is not nil.
func (ic *IndexersCreate) SetNillableSeedRatio(f *float32) *IndexersCreate {
	if f != nil {
		ic.SetSeedRatio(*f)
	}
	return ic
}

// SetDisabled sets the "disabled" field.
func (ic *IndexersCreate) SetDisabled(b bool) *IndexersCreate {
	ic.mutation.SetDisabled(b)
	return ic
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (ic *IndexersCreate) SetNillableDisabled(b *bool) *IndexersCreate {
	if b != nil {
		ic.SetDisabled(*b)
	}
	return ic
}

// SetTvSearch sets the "tv_search" field.
func (ic *IndexersCreate) SetTvSearch(b bool) *IndexersCreate {
	ic.mutation.SetTvSearch(b)
	return ic
}

// SetNillableTvSearch sets the "tv_search" field if the given value is not nil.
func (ic *IndexersCreate) SetNillableTvSearch(b *bool) *IndexersCreate {
	if b != nil {
		ic.SetTvSearch(*b)
	}
	return ic
}

// SetMovieSearch sets the "movie_search" field.
func (ic *IndexersCreate) SetMovieSearch(b bool) *IndexersCreate {
	ic.mutation.SetMovieSearch(b)
	return ic
}

// SetNillableMovieSearch sets the "movie_search" field if the given value is not nil.
func (ic *IndexersCreate) SetNillableMovieSearch(b *bool) *IndexersCreate {
	if b != nil {
		ic.SetMovieSearch(*b)
	}
	return ic
}

// SetAPIKey sets the "api_key" field.
func (ic *IndexersCreate) SetAPIKey(s string) *IndexersCreate {
	ic.mutation.SetAPIKey(s)
	return ic
}

// SetNillableAPIKey sets the "api_key" field if the given value is not nil.
func (ic *IndexersCreate) SetNillableAPIKey(s *string) *IndexersCreate {
	if s != nil {
		ic.SetAPIKey(*s)
	}
	return ic
}

// SetURL sets the "url" field.
func (ic *IndexersCreate) SetURL(s string) *IndexersCreate {
	ic.mutation.SetURL(s)
	return ic
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (ic *IndexersCreate) SetNillableURL(s *string) *IndexersCreate {
	if s != nil {
		ic.SetURL(*s)
	}
	return ic
}

// SetSynced sets the "synced" field.
func (ic *IndexersCreate) SetSynced(b bool) *IndexersCreate {
	ic.mutation.SetSynced(b)
	return ic
}

// SetNillableSynced sets the "synced" field if the given value is not nil.
func (ic *IndexersCreate) SetNillableSynced(b *bool) *IndexersCreate {
	if b != nil {
		ic.SetSynced(*b)
	}
	return ic
}

// SetCreateTime sets the "create_time" field.
func (ic *IndexersCreate) SetCreateTime(t time.Time) *IndexersCreate {
	ic.mutation.SetCreateTime(t)
	return ic
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ic *IndexersCreate) SetNillableCreateTime(t *time.Time) *IndexersCreate {
	if t != nil {
		ic.SetCreateTime(*t)
	}
	return ic
}

// Mutation returns the IndexersMutation object of the builder.
func (ic *IndexersCreate) Mutation() *IndexersMutation {
	return ic.mutation
}

// Save creates the Indexers in the database.
func (ic *IndexersCreate) Save(ctx context.Context) (*Indexers, error) {
	ic.defaults()
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *IndexersCreate) SaveX(ctx context.Context) *Indexers {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *IndexersCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *IndexersCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *IndexersCreate) defaults() {
	if _, ok := ic.mutation.Settings(); !ok {
		v := indexers.DefaultSettings
		ic.mutation.SetSettings(v)
	}
	if _, ok := ic.mutation.EnableRss(); !ok {
		v := indexers.DefaultEnableRss
		ic.mutation.SetEnableRss(v)
	}
	if _, ok := ic.mutation.Priority(); !ok {
		v := indexers.DefaultPriority
		ic.mutation.SetPriority(v)
	}
	if _, ok := ic.mutation.SeedRatio(); !ok {
		v := indexers.DefaultSeedRatio
		ic.mutation.SetSeedRatio(v)
	}
	if _, ok := ic.mutation.Disabled(); !ok {
		v := indexers.DefaultDisabled
		ic.mutation.SetDisabled(v)
	}
	if _, ok := ic.mutation.TvSearch(); !ok {
		v := indexers.DefaultTvSearch
		ic.mutation.SetTvSearch(v)
	}
	if _, ok := ic.mutation.MovieSearch(); !ok {
		v := indexers.DefaultMovieSearch
		ic.mutation.SetMovieSearch(v)
	}
	if _, ok := ic.mutation.Synced(); !ok {
		v := indexers.DefaultSynced
		ic.mutation.SetSynced(v)
	}
	if _, ok := ic.mutation.CreateTime(); !ok {
		v := indexers.DefaultCreateTime()
		ic.mutation.SetCreateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *IndexersCreate) check() error {
	if _, ok := ic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Indexers.name"`)}
	}
	if _, ok := ic.mutation.Implementation(); !ok {
		return &ValidationError{Name: "implementation", err: errors.New(`ent: missing required field "Indexers.implementation"`)}
	}
	if _, ok := ic.mutation.EnableRss(); !ok {
		return &ValidationError{Name: "enable_rss", err: errors.New(`ent: missing required field "Indexers.enable_rss"`)}
	}
	if _, ok := ic.mutation.Priority(); !ok {
		return &ValidationError{Name: "priority", err: errors.New(`ent: missing required field "Indexers.priority"`)}
	}
	if v, ok := ic.mutation.Priority(); ok {
		if err := indexers.PriorityValidator(v); err != nil {
			return &ValidationError{Name: "priority", err: fmt.Errorf(`ent: validator failed for field "Indexers.priority": %w`, err)}
		}
	}
	return nil
}

func (ic *IndexersCreate) sqlSave(ctx context.Context) (*Indexers, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *IndexersCreate) createSpec() (*Indexers, *sqlgraph.CreateSpec) {
	var (
		_node = &Indexers{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(indexers.Table, sqlgraph.NewFieldSpec(indexers.FieldID, field.TypeInt))
	)
	if value, ok := ic.mutation.Name(); ok {
		_spec.SetField(indexers.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ic.mutation.Implementation(); ok {
		_spec.SetField(indexers.FieldImplementation, field.TypeString, value)
		_node.Implementation = value
	}
	if value, ok := ic.mutation.Settings(); ok {
		_spec.SetField(indexers.FieldSettings, field.TypeString, value)
		_node.Settings = value
	}
	if value, ok := ic.mutation.EnableRss(); ok {
		_spec.SetField(indexers.FieldEnableRss, field.TypeBool, value)
		_node.EnableRss = value
	}
	if value, ok := ic.mutation.Priority(); ok {
		_spec.SetField(indexers.FieldPriority, field.TypeInt, value)
		_node.Priority = value
	}
	if value, ok := ic.mutation.SeedRatio(); ok {
		_spec.SetField(indexers.FieldSeedRatio, field.TypeFloat32, value)
		_node.SeedRatio = value
	}
	if value, ok := ic.mutation.Disabled(); ok {
		_spec.SetField(indexers.FieldDisabled, field.TypeBool, value)
		_node.Disabled = value
	}
	if value, ok := ic.mutation.TvSearch(); ok {
		_spec.SetField(indexers.FieldTvSearch, field.TypeBool, value)
		_node.TvSearch = value
	}
	if value, ok := ic.mutation.MovieSearch(); ok {
		_spec.SetField(indexers.FieldMovieSearch, field.TypeBool, value)
		_node.MovieSearch = value
	}
	if value, ok := ic.mutation.APIKey(); ok {
		_spec.SetField(indexers.FieldAPIKey, field.TypeString, value)
		_node.APIKey = value
	}
	if value, ok := ic.mutation.URL(); ok {
		_spec.SetField(indexers.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := ic.mutation.Synced(); ok {
		_spec.SetField(indexers.FieldSynced, field.TypeBool, value)
		_node.Synced = value
	}
	if value, ok := ic.mutation.CreateTime(); ok {
		_spec.SetField(indexers.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	return _node, _spec
}

// IndexersCreateBulk is the builder for creating many Indexers entities in bulk.
type IndexersCreateBulk struct {
	config
	err      error
	builders []*IndexersCreate
}

// Save creates the Indexers entities in the database.
func (icb *IndexersCreateBulk) Save(ctx context.Context) ([]*Indexers, error) {
	if icb.err != nil {
		return nil, icb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Indexers, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IndexersMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *IndexersCreateBulk) SaveX(ctx context.Context) []*Indexers {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *IndexersCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *IndexersCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
