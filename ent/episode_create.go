// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"polaris/ent/episode"
	"polaris/ent/series"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EpisodeCreate is the builder for creating a Episode entity.
type EpisodeCreate struct {
	config
	mutation *EpisodeMutation
	hooks    []Hook
}

// SetSeriesID sets the "series_id" field.
func (ec *EpisodeCreate) SetSeriesID(i int) *EpisodeCreate {
	ec.mutation.SetSeriesID(i)
	return ec
}

// SetNillableSeriesID sets the "series_id" field if the given value is not nil.
func (ec *EpisodeCreate) SetNillableSeriesID(i *int) *EpisodeCreate {
	if i != nil {
		ec.SetSeriesID(*i)
	}
	return ec
}

// SetSeasonNumber sets the "season_number" field.
func (ec *EpisodeCreate) SetSeasonNumber(i int) *EpisodeCreate {
	ec.mutation.SetSeasonNumber(i)
	return ec
}

// SetEpisodeNumber sets the "episode_number" field.
func (ec *EpisodeCreate) SetEpisodeNumber(i int) *EpisodeCreate {
	ec.mutation.SetEpisodeNumber(i)
	return ec
}

// SetTitle sets the "title" field.
func (ec *EpisodeCreate) SetTitle(s string) *EpisodeCreate {
	ec.mutation.SetTitle(s)
	return ec
}

// SetOverview sets the "overview" field.
func (ec *EpisodeCreate) SetOverview(s string) *EpisodeCreate {
	ec.mutation.SetOverview(s)
	return ec
}

// SetAirDate sets the "air_date" field.
func (ec *EpisodeCreate) SetAirDate(s string) *EpisodeCreate {
	ec.mutation.SetAirDate(s)
	return ec
}

// SetStatus sets the "status" field.
func (ec *EpisodeCreate) SetStatus(e episode.Status) *EpisodeCreate {
	ec.mutation.SetStatus(e)
	return ec
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ec *EpisodeCreate) SetNillableStatus(e *episode.Status) *EpisodeCreate {
	if e != nil {
		ec.SetStatus(*e)
	}
	return ec
}

// SetFileInStorage sets the "file_in_storage" field.
func (ec *EpisodeCreate) SetFileInStorage(s string) *EpisodeCreate {
	ec.mutation.SetFileInStorage(s)
	return ec
}

// SetNillableFileInStorage sets the "file_in_storage" field if the given value is not nil.
func (ec *EpisodeCreate) SetNillableFileInStorage(s *string) *EpisodeCreate {
	if s != nil {
		ec.SetFileInStorage(*s)
	}
	return ec
}

// SetSeries sets the "series" edge to the Series entity.
func (ec *EpisodeCreate) SetSeries(s *Series) *EpisodeCreate {
	return ec.SetSeriesID(s.ID)
}

// Mutation returns the EpisodeMutation object of the builder.
func (ec *EpisodeCreate) Mutation() *EpisodeMutation {
	return ec.mutation
}

// Save creates the Episode in the database.
func (ec *EpisodeCreate) Save(ctx context.Context) (*Episode, error) {
	ec.defaults()
	return withHooks(ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EpisodeCreate) SaveX(ctx context.Context) *Episode {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EpisodeCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EpisodeCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EpisodeCreate) defaults() {
	if _, ok := ec.mutation.Status(); !ok {
		v := episode.DefaultStatus
		ec.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EpisodeCreate) check() error {
	if _, ok := ec.mutation.SeasonNumber(); !ok {
		return &ValidationError{Name: "season_number", err: errors.New(`ent: missing required field "Episode.season_number"`)}
	}
	if _, ok := ec.mutation.EpisodeNumber(); !ok {
		return &ValidationError{Name: "episode_number", err: errors.New(`ent: missing required field "Episode.episode_number"`)}
	}
	if _, ok := ec.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Episode.title"`)}
	}
	if _, ok := ec.mutation.Overview(); !ok {
		return &ValidationError{Name: "overview", err: errors.New(`ent: missing required field "Episode.overview"`)}
	}
	if _, ok := ec.mutation.AirDate(); !ok {
		return &ValidationError{Name: "air_date", err: errors.New(`ent: missing required field "Episode.air_date"`)}
	}
	if _, ok := ec.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Episode.status"`)}
	}
	if v, ok := ec.mutation.Status(); ok {
		if err := episode.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Episode.status": %w`, err)}
		}
	}
	return nil
}

func (ec *EpisodeCreate) sqlSave(ctx context.Context) (*Episode, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *EpisodeCreate) createSpec() (*Episode, *sqlgraph.CreateSpec) {
	var (
		_node = &Episode{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(episode.Table, sqlgraph.NewFieldSpec(episode.FieldID, field.TypeInt))
	)
	if value, ok := ec.mutation.SeasonNumber(); ok {
		_spec.SetField(episode.FieldSeasonNumber, field.TypeInt, value)
		_node.SeasonNumber = value
	}
	if value, ok := ec.mutation.EpisodeNumber(); ok {
		_spec.SetField(episode.FieldEpisodeNumber, field.TypeInt, value)
		_node.EpisodeNumber = value
	}
	if value, ok := ec.mutation.Title(); ok {
		_spec.SetField(episode.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := ec.mutation.Overview(); ok {
		_spec.SetField(episode.FieldOverview, field.TypeString, value)
		_node.Overview = value
	}
	if value, ok := ec.mutation.AirDate(); ok {
		_spec.SetField(episode.FieldAirDate, field.TypeString, value)
		_node.AirDate = value
	}
	if value, ok := ec.mutation.Status(); ok {
		_spec.SetField(episode.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := ec.mutation.FileInStorage(); ok {
		_spec.SetField(episode.FieldFileInStorage, field.TypeString, value)
		_node.FileInStorage = value
	}
	if nodes := ec.mutation.SeriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   episode.SeriesTable,
			Columns: []string{episode.SeriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(series.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SeriesID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EpisodeCreateBulk is the builder for creating many Episode entities in bulk.
type EpisodeCreateBulk struct {
	config
	err      error
	builders []*EpisodeCreate
}

// Save creates the Episode entities in the database.
func (ecb *EpisodeCreateBulk) Save(ctx context.Context) ([]*Episode, error) {
	if ecb.err != nil {
		return nil, ecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Episode, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EpisodeMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EpisodeCreateBulk) SaveX(ctx context.Context) []*Episode {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EpisodeCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EpisodeCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
