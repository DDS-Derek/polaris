// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"polaris/ent/episode"
	"polaris/ent/media"
	"polaris/ent/schema"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MediaCreate is the builder for creating a Media entity.
type MediaCreate struct {
	config
	mutation *MediaMutation
	hooks    []Hook
}

// SetTmdbID sets the "tmdb_id" field.
func (mc *MediaCreate) SetTmdbID(i int) *MediaCreate {
	mc.mutation.SetTmdbID(i)
	return mc
}

// SetImdbID sets the "imdb_id" field.
func (mc *MediaCreate) SetImdbID(s string) *MediaCreate {
	mc.mutation.SetImdbID(s)
	return mc
}

// SetNillableImdbID sets the "imdb_id" field if the given value is not nil.
func (mc *MediaCreate) SetNillableImdbID(s *string) *MediaCreate {
	if s != nil {
		mc.SetImdbID(*s)
	}
	return mc
}

// SetMediaType sets the "media_type" field.
func (mc *MediaCreate) SetMediaType(mt media.MediaType) *MediaCreate {
	mc.mutation.SetMediaType(mt)
	return mc
}

// SetNameCn sets the "name_cn" field.
func (mc *MediaCreate) SetNameCn(s string) *MediaCreate {
	mc.mutation.SetNameCn(s)
	return mc
}

// SetNameEn sets the "name_en" field.
func (mc *MediaCreate) SetNameEn(s string) *MediaCreate {
	mc.mutation.SetNameEn(s)
	return mc
}

// SetOriginalName sets the "original_name" field.
func (mc *MediaCreate) SetOriginalName(s string) *MediaCreate {
	mc.mutation.SetOriginalName(s)
	return mc
}

// SetOverview sets the "overview" field.
func (mc *MediaCreate) SetOverview(s string) *MediaCreate {
	mc.mutation.SetOverview(s)
	return mc
}

// SetCreatedAt sets the "created_at" field.
func (mc *MediaCreate) SetCreatedAt(t time.Time) *MediaCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MediaCreate) SetNillableCreatedAt(t *time.Time) *MediaCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetAirDate sets the "air_date" field.
func (mc *MediaCreate) SetAirDate(s string) *MediaCreate {
	mc.mutation.SetAirDate(s)
	return mc
}

// SetNillableAirDate sets the "air_date" field if the given value is not nil.
func (mc *MediaCreate) SetNillableAirDate(s *string) *MediaCreate {
	if s != nil {
		mc.SetAirDate(*s)
	}
	return mc
}

// SetResolution sets the "resolution" field.
func (mc *MediaCreate) SetResolution(m media.Resolution) *MediaCreate {
	mc.mutation.SetResolution(m)
	return mc
}

// SetNillableResolution sets the "resolution" field if the given value is not nil.
func (mc *MediaCreate) SetNillableResolution(m *media.Resolution) *MediaCreate {
	if m != nil {
		mc.SetResolution(*m)
	}
	return mc
}

// SetStorageID sets the "storage_id" field.
func (mc *MediaCreate) SetStorageID(i int) *MediaCreate {
	mc.mutation.SetStorageID(i)
	return mc
}

// SetNillableStorageID sets the "storage_id" field if the given value is not nil.
func (mc *MediaCreate) SetNillableStorageID(i *int) *MediaCreate {
	if i != nil {
		mc.SetStorageID(*i)
	}
	return mc
}

// SetTargetDir sets the "target_dir" field.
func (mc *MediaCreate) SetTargetDir(s string) *MediaCreate {
	mc.mutation.SetTargetDir(s)
	return mc
}

// SetNillableTargetDir sets the "target_dir" field if the given value is not nil.
func (mc *MediaCreate) SetNillableTargetDir(s *string) *MediaCreate {
	if s != nil {
		mc.SetTargetDir(*s)
	}
	return mc
}

// SetDownloadHistoryEpisodes sets the "download_history_episodes" field.
func (mc *MediaCreate) SetDownloadHistoryEpisodes(b bool) *MediaCreate {
	mc.mutation.SetDownloadHistoryEpisodes(b)
	return mc
}

// SetNillableDownloadHistoryEpisodes sets the "download_history_episodes" field if the given value is not nil.
func (mc *MediaCreate) SetNillableDownloadHistoryEpisodes(b *bool) *MediaCreate {
	if b != nil {
		mc.SetDownloadHistoryEpisodes(*b)
	}
	return mc
}

// SetLimiter sets the "limiter" field.
func (mc *MediaCreate) SetLimiter(sl schema.MediaLimiter) *MediaCreate {
	mc.mutation.SetLimiter(sl)
	return mc
}

// SetNillableLimiter sets the "limiter" field if the given value is not nil.
func (mc *MediaCreate) SetNillableLimiter(sl *schema.MediaLimiter) *MediaCreate {
	if sl != nil {
		mc.SetLimiter(*sl)
	}
	return mc
}

// SetExtras sets the "extras" field.
func (mc *MediaCreate) SetExtras(se schema.MediaExtras) *MediaCreate {
	mc.mutation.SetExtras(se)
	return mc
}

// SetNillableExtras sets the "extras" field if the given value is not nil.
func (mc *MediaCreate) SetNillableExtras(se *schema.MediaExtras) *MediaCreate {
	if se != nil {
		mc.SetExtras(*se)
	}
	return mc
}

// SetAlternativeTitles sets the "alternative_titles" field.
func (mc *MediaCreate) SetAlternativeTitles(st []schema.AlternativeTilte) *MediaCreate {
	mc.mutation.SetAlternativeTitles(st)
	return mc
}

// SetCreateTime sets the "create_time" field.
func (mc *MediaCreate) SetCreateTime(t time.Time) *MediaCreate {
	mc.mutation.SetCreateTime(t)
	return mc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (mc *MediaCreate) SetNillableCreateTime(t *time.Time) *MediaCreate {
	if t != nil {
		mc.SetCreateTime(*t)
	}
	return mc
}

// AddEpisodeIDs adds the "episodes" edge to the Episode entity by IDs.
func (mc *MediaCreate) AddEpisodeIDs(ids ...int) *MediaCreate {
	mc.mutation.AddEpisodeIDs(ids...)
	return mc
}

// AddEpisodes adds the "episodes" edges to the Episode entity.
func (mc *MediaCreate) AddEpisodes(e ...*Episode) *MediaCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return mc.AddEpisodeIDs(ids...)
}

// Mutation returns the MediaMutation object of the builder.
func (mc *MediaCreate) Mutation() *MediaMutation {
	return mc.mutation
}

// Save creates the Media in the database.
func (mc *MediaCreate) Save(ctx context.Context) (*Media, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MediaCreate) SaveX(ctx context.Context) *Media {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MediaCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MediaCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MediaCreate) defaults() {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := media.DefaultCreatedAt
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.AirDate(); !ok {
		v := media.DefaultAirDate
		mc.mutation.SetAirDate(v)
	}
	if _, ok := mc.mutation.Resolution(); !ok {
		v := media.DefaultResolution
		mc.mutation.SetResolution(v)
	}
	if _, ok := mc.mutation.DownloadHistoryEpisodes(); !ok {
		v := media.DefaultDownloadHistoryEpisodes
		mc.mutation.SetDownloadHistoryEpisodes(v)
	}
	if _, ok := mc.mutation.CreateTime(); !ok {
		v := media.DefaultCreateTime()
		mc.mutation.SetCreateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MediaCreate) check() error {
	if _, ok := mc.mutation.TmdbID(); !ok {
		return &ValidationError{Name: "tmdb_id", err: errors.New(`ent: missing required field "Media.tmdb_id"`)}
	}
	if _, ok := mc.mutation.MediaType(); !ok {
		return &ValidationError{Name: "media_type", err: errors.New(`ent: missing required field "Media.media_type"`)}
	}
	if v, ok := mc.mutation.MediaType(); ok {
		if err := media.MediaTypeValidator(v); err != nil {
			return &ValidationError{Name: "media_type", err: fmt.Errorf(`ent: validator failed for field "Media.media_type": %w`, err)}
		}
	}
	if _, ok := mc.mutation.NameCn(); !ok {
		return &ValidationError{Name: "name_cn", err: errors.New(`ent: missing required field "Media.name_cn"`)}
	}
	if _, ok := mc.mutation.NameEn(); !ok {
		return &ValidationError{Name: "name_en", err: errors.New(`ent: missing required field "Media.name_en"`)}
	}
	if _, ok := mc.mutation.OriginalName(); !ok {
		return &ValidationError{Name: "original_name", err: errors.New(`ent: missing required field "Media.original_name"`)}
	}
	if _, ok := mc.mutation.Overview(); !ok {
		return &ValidationError{Name: "overview", err: errors.New(`ent: missing required field "Media.overview"`)}
	}
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Media.created_at"`)}
	}
	if _, ok := mc.mutation.AirDate(); !ok {
		return &ValidationError{Name: "air_date", err: errors.New(`ent: missing required field "Media.air_date"`)}
	}
	if _, ok := mc.mutation.Resolution(); !ok {
		return &ValidationError{Name: "resolution", err: errors.New(`ent: missing required field "Media.resolution"`)}
	}
	if v, ok := mc.mutation.Resolution(); ok {
		if err := media.ResolutionValidator(v); err != nil {
			return &ValidationError{Name: "resolution", err: fmt.Errorf(`ent: validator failed for field "Media.resolution": %w`, err)}
		}
	}
	return nil
}

func (mc *MediaCreate) sqlSave(ctx context.Context) (*Media, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MediaCreate) createSpec() (*Media, *sqlgraph.CreateSpec) {
	var (
		_node = &Media{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(media.Table, sqlgraph.NewFieldSpec(media.FieldID, field.TypeInt))
	)
	if value, ok := mc.mutation.TmdbID(); ok {
		_spec.SetField(media.FieldTmdbID, field.TypeInt, value)
		_node.TmdbID = value
	}
	if value, ok := mc.mutation.ImdbID(); ok {
		_spec.SetField(media.FieldImdbID, field.TypeString, value)
		_node.ImdbID = value
	}
	if value, ok := mc.mutation.MediaType(); ok {
		_spec.SetField(media.FieldMediaType, field.TypeEnum, value)
		_node.MediaType = value
	}
	if value, ok := mc.mutation.NameCn(); ok {
		_spec.SetField(media.FieldNameCn, field.TypeString, value)
		_node.NameCn = value
	}
	if value, ok := mc.mutation.NameEn(); ok {
		_spec.SetField(media.FieldNameEn, field.TypeString, value)
		_node.NameEn = value
	}
	if value, ok := mc.mutation.OriginalName(); ok {
		_spec.SetField(media.FieldOriginalName, field.TypeString, value)
		_node.OriginalName = value
	}
	if value, ok := mc.mutation.Overview(); ok {
		_spec.SetField(media.FieldOverview, field.TypeString, value)
		_node.Overview = value
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(media.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.AirDate(); ok {
		_spec.SetField(media.FieldAirDate, field.TypeString, value)
		_node.AirDate = value
	}
	if value, ok := mc.mutation.Resolution(); ok {
		_spec.SetField(media.FieldResolution, field.TypeEnum, value)
		_node.Resolution = value
	}
	if value, ok := mc.mutation.StorageID(); ok {
		_spec.SetField(media.FieldStorageID, field.TypeInt, value)
		_node.StorageID = value
	}
	if value, ok := mc.mutation.TargetDir(); ok {
		_spec.SetField(media.FieldTargetDir, field.TypeString, value)
		_node.TargetDir = value
	}
	if value, ok := mc.mutation.DownloadHistoryEpisodes(); ok {
		_spec.SetField(media.FieldDownloadHistoryEpisodes, field.TypeBool, value)
		_node.DownloadHistoryEpisodes = value
	}
	if value, ok := mc.mutation.Limiter(); ok {
		_spec.SetField(media.FieldLimiter, field.TypeJSON, value)
		_node.Limiter = value
	}
	if value, ok := mc.mutation.Extras(); ok {
		_spec.SetField(media.FieldExtras, field.TypeJSON, value)
		_node.Extras = value
	}
	if value, ok := mc.mutation.AlternativeTitles(); ok {
		_spec.SetField(media.FieldAlternativeTitles, field.TypeJSON, value)
		_node.AlternativeTitles = value
	}
	if value, ok := mc.mutation.CreateTime(); ok {
		_spec.SetField(media.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if nodes := mc.mutation.EpisodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   media.EpisodesTable,
			Columns: []string{media.EpisodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(episode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MediaCreateBulk is the builder for creating many Media entities in bulk.
type MediaCreateBulk struct {
	config
	err      error
	builders []*MediaCreate
}

// Save creates the Media entities in the database.
func (mcb *MediaCreateBulk) Save(ctx context.Context) ([]*Media, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Media, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MediaMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MediaCreateBulk) SaveX(ctx context.Context) []*Media {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MediaCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MediaCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
