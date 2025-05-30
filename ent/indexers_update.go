// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"polaris/ent/indexers"
	"polaris/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// IndexersUpdate is the builder for updating Indexers entities.
type IndexersUpdate struct {
	config
	hooks    []Hook
	mutation *IndexersMutation
}

// Where appends a list predicates to the IndexersUpdate builder.
func (iu *IndexersUpdate) Where(ps ...predicate.Indexers) *IndexersUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetName sets the "name" field.
func (iu *IndexersUpdate) SetName(s string) *IndexersUpdate {
	iu.mutation.SetName(s)
	return iu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (iu *IndexersUpdate) SetNillableName(s *string) *IndexersUpdate {
	if s != nil {
		iu.SetName(*s)
	}
	return iu
}

// SetImplementation sets the "implementation" field.
func (iu *IndexersUpdate) SetImplementation(s string) *IndexersUpdate {
	iu.mutation.SetImplementation(s)
	return iu
}

// SetNillableImplementation sets the "implementation" field if the given value is not nil.
func (iu *IndexersUpdate) SetNillableImplementation(s *string) *IndexersUpdate {
	if s != nil {
		iu.SetImplementation(*s)
	}
	return iu
}

// SetSettings sets the "settings" field.
func (iu *IndexersUpdate) SetSettings(s string) *IndexersUpdate {
	iu.mutation.SetSettings(s)
	return iu
}

// SetNillableSettings sets the "settings" field if the given value is not nil.
func (iu *IndexersUpdate) SetNillableSettings(s *string) *IndexersUpdate {
	if s != nil {
		iu.SetSettings(*s)
	}
	return iu
}

// ClearSettings clears the value of the "settings" field.
func (iu *IndexersUpdate) ClearSettings() *IndexersUpdate {
	iu.mutation.ClearSettings()
	return iu
}

// SetEnableRss sets the "enable_rss" field.
func (iu *IndexersUpdate) SetEnableRss(b bool) *IndexersUpdate {
	iu.mutation.SetEnableRss(b)
	return iu
}

// SetNillableEnableRss sets the "enable_rss" field if the given value is not nil.
func (iu *IndexersUpdate) SetNillableEnableRss(b *bool) *IndexersUpdate {
	if b != nil {
		iu.SetEnableRss(*b)
	}
	return iu
}

// SetPriority sets the "priority" field.
func (iu *IndexersUpdate) SetPriority(i int) *IndexersUpdate {
	iu.mutation.ResetPriority()
	iu.mutation.SetPriority(i)
	return iu
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (iu *IndexersUpdate) SetNillablePriority(i *int) *IndexersUpdate {
	if i != nil {
		iu.SetPriority(*i)
	}
	return iu
}

// AddPriority adds i to the "priority" field.
func (iu *IndexersUpdate) AddPriority(i int) *IndexersUpdate {
	iu.mutation.AddPriority(i)
	return iu
}

// SetSeedRatio sets the "seed_ratio" field.
func (iu *IndexersUpdate) SetSeedRatio(f float32) *IndexersUpdate {
	iu.mutation.ResetSeedRatio()
	iu.mutation.SetSeedRatio(f)
	return iu
}

// SetNillableSeedRatio sets the "seed_ratio" field if the given value is not nil.
func (iu *IndexersUpdate) SetNillableSeedRatio(f *float32) *IndexersUpdate {
	if f != nil {
		iu.SetSeedRatio(*f)
	}
	return iu
}

// AddSeedRatio adds f to the "seed_ratio" field.
func (iu *IndexersUpdate) AddSeedRatio(f float32) *IndexersUpdate {
	iu.mutation.AddSeedRatio(f)
	return iu
}

// ClearSeedRatio clears the value of the "seed_ratio" field.
func (iu *IndexersUpdate) ClearSeedRatio() *IndexersUpdate {
	iu.mutation.ClearSeedRatio()
	return iu
}

// SetDisabled sets the "disabled" field.
func (iu *IndexersUpdate) SetDisabled(b bool) *IndexersUpdate {
	iu.mutation.SetDisabled(b)
	return iu
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (iu *IndexersUpdate) SetNillableDisabled(b *bool) *IndexersUpdate {
	if b != nil {
		iu.SetDisabled(*b)
	}
	return iu
}

// ClearDisabled clears the value of the "disabled" field.
func (iu *IndexersUpdate) ClearDisabled() *IndexersUpdate {
	iu.mutation.ClearDisabled()
	return iu
}

// SetTvSearch sets the "tv_search" field.
func (iu *IndexersUpdate) SetTvSearch(b bool) *IndexersUpdate {
	iu.mutation.SetTvSearch(b)
	return iu
}

// SetNillableTvSearch sets the "tv_search" field if the given value is not nil.
func (iu *IndexersUpdate) SetNillableTvSearch(b *bool) *IndexersUpdate {
	if b != nil {
		iu.SetTvSearch(*b)
	}
	return iu
}

// ClearTvSearch clears the value of the "tv_search" field.
func (iu *IndexersUpdate) ClearTvSearch() *IndexersUpdate {
	iu.mutation.ClearTvSearch()
	return iu
}

// SetMovieSearch sets the "movie_search" field.
func (iu *IndexersUpdate) SetMovieSearch(b bool) *IndexersUpdate {
	iu.mutation.SetMovieSearch(b)
	return iu
}

// SetNillableMovieSearch sets the "movie_search" field if the given value is not nil.
func (iu *IndexersUpdate) SetNillableMovieSearch(b *bool) *IndexersUpdate {
	if b != nil {
		iu.SetMovieSearch(*b)
	}
	return iu
}

// ClearMovieSearch clears the value of the "movie_search" field.
func (iu *IndexersUpdate) ClearMovieSearch() *IndexersUpdate {
	iu.mutation.ClearMovieSearch()
	return iu
}

// SetAPIKey sets the "api_key" field.
func (iu *IndexersUpdate) SetAPIKey(s string) *IndexersUpdate {
	iu.mutation.SetAPIKey(s)
	return iu
}

// SetNillableAPIKey sets the "api_key" field if the given value is not nil.
func (iu *IndexersUpdate) SetNillableAPIKey(s *string) *IndexersUpdate {
	if s != nil {
		iu.SetAPIKey(*s)
	}
	return iu
}

// ClearAPIKey clears the value of the "api_key" field.
func (iu *IndexersUpdate) ClearAPIKey() *IndexersUpdate {
	iu.mutation.ClearAPIKey()
	return iu
}

// SetURL sets the "url" field.
func (iu *IndexersUpdate) SetURL(s string) *IndexersUpdate {
	iu.mutation.SetURL(s)
	return iu
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (iu *IndexersUpdate) SetNillableURL(s *string) *IndexersUpdate {
	if s != nil {
		iu.SetURL(*s)
	}
	return iu
}

// ClearURL clears the value of the "url" field.
func (iu *IndexersUpdate) ClearURL() *IndexersUpdate {
	iu.mutation.ClearURL()
	return iu
}

// SetSynced sets the "synced" field.
func (iu *IndexersUpdate) SetSynced(b bool) *IndexersUpdate {
	iu.mutation.SetSynced(b)
	return iu
}

// SetNillableSynced sets the "synced" field if the given value is not nil.
func (iu *IndexersUpdate) SetNillableSynced(b *bool) *IndexersUpdate {
	if b != nil {
		iu.SetSynced(*b)
	}
	return iu
}

// ClearSynced clears the value of the "synced" field.
func (iu *IndexersUpdate) ClearSynced() *IndexersUpdate {
	iu.mutation.ClearSynced()
	return iu
}

// Mutation returns the IndexersMutation object of the builder.
func (iu *IndexersUpdate) Mutation() *IndexersMutation {
	return iu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *IndexersUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *IndexersUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *IndexersUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *IndexersUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *IndexersUpdate) check() error {
	if v, ok := iu.mutation.Priority(); ok {
		if err := indexers.PriorityValidator(v); err != nil {
			return &ValidationError{Name: "priority", err: fmt.Errorf(`ent: validator failed for field "Indexers.priority": %w`, err)}
		}
	}
	return nil
}

func (iu *IndexersUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(indexers.Table, indexers.Columns, sqlgraph.NewFieldSpec(indexers.FieldID, field.TypeInt))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.Name(); ok {
		_spec.SetField(indexers.FieldName, field.TypeString, value)
	}
	if value, ok := iu.mutation.Implementation(); ok {
		_spec.SetField(indexers.FieldImplementation, field.TypeString, value)
	}
	if value, ok := iu.mutation.Settings(); ok {
		_spec.SetField(indexers.FieldSettings, field.TypeString, value)
	}
	if iu.mutation.SettingsCleared() {
		_spec.ClearField(indexers.FieldSettings, field.TypeString)
	}
	if value, ok := iu.mutation.EnableRss(); ok {
		_spec.SetField(indexers.FieldEnableRss, field.TypeBool, value)
	}
	if value, ok := iu.mutation.Priority(); ok {
		_spec.SetField(indexers.FieldPriority, field.TypeInt, value)
	}
	if value, ok := iu.mutation.AddedPriority(); ok {
		_spec.AddField(indexers.FieldPriority, field.TypeInt, value)
	}
	if value, ok := iu.mutation.SeedRatio(); ok {
		_spec.SetField(indexers.FieldSeedRatio, field.TypeFloat32, value)
	}
	if value, ok := iu.mutation.AddedSeedRatio(); ok {
		_spec.AddField(indexers.FieldSeedRatio, field.TypeFloat32, value)
	}
	if iu.mutation.SeedRatioCleared() {
		_spec.ClearField(indexers.FieldSeedRatio, field.TypeFloat32)
	}
	if value, ok := iu.mutation.Disabled(); ok {
		_spec.SetField(indexers.FieldDisabled, field.TypeBool, value)
	}
	if iu.mutation.DisabledCleared() {
		_spec.ClearField(indexers.FieldDisabled, field.TypeBool)
	}
	if value, ok := iu.mutation.TvSearch(); ok {
		_spec.SetField(indexers.FieldTvSearch, field.TypeBool, value)
	}
	if iu.mutation.TvSearchCleared() {
		_spec.ClearField(indexers.FieldTvSearch, field.TypeBool)
	}
	if value, ok := iu.mutation.MovieSearch(); ok {
		_spec.SetField(indexers.FieldMovieSearch, field.TypeBool, value)
	}
	if iu.mutation.MovieSearchCleared() {
		_spec.ClearField(indexers.FieldMovieSearch, field.TypeBool)
	}
	if value, ok := iu.mutation.APIKey(); ok {
		_spec.SetField(indexers.FieldAPIKey, field.TypeString, value)
	}
	if iu.mutation.APIKeyCleared() {
		_spec.ClearField(indexers.FieldAPIKey, field.TypeString)
	}
	if value, ok := iu.mutation.URL(); ok {
		_spec.SetField(indexers.FieldURL, field.TypeString, value)
	}
	if iu.mutation.URLCleared() {
		_spec.ClearField(indexers.FieldURL, field.TypeString)
	}
	if value, ok := iu.mutation.Synced(); ok {
		_spec.SetField(indexers.FieldSynced, field.TypeBool, value)
	}
	if iu.mutation.SyncedCleared() {
		_spec.ClearField(indexers.FieldSynced, field.TypeBool)
	}
	if iu.mutation.CreateTimeCleared() {
		_spec.ClearField(indexers.FieldCreateTime, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{indexers.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// IndexersUpdateOne is the builder for updating a single Indexers entity.
type IndexersUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IndexersMutation
}

// SetName sets the "name" field.
func (iuo *IndexersUpdateOne) SetName(s string) *IndexersUpdateOne {
	iuo.mutation.SetName(s)
	return iuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (iuo *IndexersUpdateOne) SetNillableName(s *string) *IndexersUpdateOne {
	if s != nil {
		iuo.SetName(*s)
	}
	return iuo
}

// SetImplementation sets the "implementation" field.
func (iuo *IndexersUpdateOne) SetImplementation(s string) *IndexersUpdateOne {
	iuo.mutation.SetImplementation(s)
	return iuo
}

// SetNillableImplementation sets the "implementation" field if the given value is not nil.
func (iuo *IndexersUpdateOne) SetNillableImplementation(s *string) *IndexersUpdateOne {
	if s != nil {
		iuo.SetImplementation(*s)
	}
	return iuo
}

// SetSettings sets the "settings" field.
func (iuo *IndexersUpdateOne) SetSettings(s string) *IndexersUpdateOne {
	iuo.mutation.SetSettings(s)
	return iuo
}

// SetNillableSettings sets the "settings" field if the given value is not nil.
func (iuo *IndexersUpdateOne) SetNillableSettings(s *string) *IndexersUpdateOne {
	if s != nil {
		iuo.SetSettings(*s)
	}
	return iuo
}

// ClearSettings clears the value of the "settings" field.
func (iuo *IndexersUpdateOne) ClearSettings() *IndexersUpdateOne {
	iuo.mutation.ClearSettings()
	return iuo
}

// SetEnableRss sets the "enable_rss" field.
func (iuo *IndexersUpdateOne) SetEnableRss(b bool) *IndexersUpdateOne {
	iuo.mutation.SetEnableRss(b)
	return iuo
}

// SetNillableEnableRss sets the "enable_rss" field if the given value is not nil.
func (iuo *IndexersUpdateOne) SetNillableEnableRss(b *bool) *IndexersUpdateOne {
	if b != nil {
		iuo.SetEnableRss(*b)
	}
	return iuo
}

// SetPriority sets the "priority" field.
func (iuo *IndexersUpdateOne) SetPriority(i int) *IndexersUpdateOne {
	iuo.mutation.ResetPriority()
	iuo.mutation.SetPriority(i)
	return iuo
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (iuo *IndexersUpdateOne) SetNillablePriority(i *int) *IndexersUpdateOne {
	if i != nil {
		iuo.SetPriority(*i)
	}
	return iuo
}

// AddPriority adds i to the "priority" field.
func (iuo *IndexersUpdateOne) AddPriority(i int) *IndexersUpdateOne {
	iuo.mutation.AddPriority(i)
	return iuo
}

// SetSeedRatio sets the "seed_ratio" field.
func (iuo *IndexersUpdateOne) SetSeedRatio(f float32) *IndexersUpdateOne {
	iuo.mutation.ResetSeedRatio()
	iuo.mutation.SetSeedRatio(f)
	return iuo
}

// SetNillableSeedRatio sets the "seed_ratio" field if the given value is not nil.
func (iuo *IndexersUpdateOne) SetNillableSeedRatio(f *float32) *IndexersUpdateOne {
	if f != nil {
		iuo.SetSeedRatio(*f)
	}
	return iuo
}

// AddSeedRatio adds f to the "seed_ratio" field.
func (iuo *IndexersUpdateOne) AddSeedRatio(f float32) *IndexersUpdateOne {
	iuo.mutation.AddSeedRatio(f)
	return iuo
}

// ClearSeedRatio clears the value of the "seed_ratio" field.
func (iuo *IndexersUpdateOne) ClearSeedRatio() *IndexersUpdateOne {
	iuo.mutation.ClearSeedRatio()
	return iuo
}

// SetDisabled sets the "disabled" field.
func (iuo *IndexersUpdateOne) SetDisabled(b bool) *IndexersUpdateOne {
	iuo.mutation.SetDisabled(b)
	return iuo
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (iuo *IndexersUpdateOne) SetNillableDisabled(b *bool) *IndexersUpdateOne {
	if b != nil {
		iuo.SetDisabled(*b)
	}
	return iuo
}

// ClearDisabled clears the value of the "disabled" field.
func (iuo *IndexersUpdateOne) ClearDisabled() *IndexersUpdateOne {
	iuo.mutation.ClearDisabled()
	return iuo
}

// SetTvSearch sets the "tv_search" field.
func (iuo *IndexersUpdateOne) SetTvSearch(b bool) *IndexersUpdateOne {
	iuo.mutation.SetTvSearch(b)
	return iuo
}

// SetNillableTvSearch sets the "tv_search" field if the given value is not nil.
func (iuo *IndexersUpdateOne) SetNillableTvSearch(b *bool) *IndexersUpdateOne {
	if b != nil {
		iuo.SetTvSearch(*b)
	}
	return iuo
}

// ClearTvSearch clears the value of the "tv_search" field.
func (iuo *IndexersUpdateOne) ClearTvSearch() *IndexersUpdateOne {
	iuo.mutation.ClearTvSearch()
	return iuo
}

// SetMovieSearch sets the "movie_search" field.
func (iuo *IndexersUpdateOne) SetMovieSearch(b bool) *IndexersUpdateOne {
	iuo.mutation.SetMovieSearch(b)
	return iuo
}

// SetNillableMovieSearch sets the "movie_search" field if the given value is not nil.
func (iuo *IndexersUpdateOne) SetNillableMovieSearch(b *bool) *IndexersUpdateOne {
	if b != nil {
		iuo.SetMovieSearch(*b)
	}
	return iuo
}

// ClearMovieSearch clears the value of the "movie_search" field.
func (iuo *IndexersUpdateOne) ClearMovieSearch() *IndexersUpdateOne {
	iuo.mutation.ClearMovieSearch()
	return iuo
}

// SetAPIKey sets the "api_key" field.
func (iuo *IndexersUpdateOne) SetAPIKey(s string) *IndexersUpdateOne {
	iuo.mutation.SetAPIKey(s)
	return iuo
}

// SetNillableAPIKey sets the "api_key" field if the given value is not nil.
func (iuo *IndexersUpdateOne) SetNillableAPIKey(s *string) *IndexersUpdateOne {
	if s != nil {
		iuo.SetAPIKey(*s)
	}
	return iuo
}

// ClearAPIKey clears the value of the "api_key" field.
func (iuo *IndexersUpdateOne) ClearAPIKey() *IndexersUpdateOne {
	iuo.mutation.ClearAPIKey()
	return iuo
}

// SetURL sets the "url" field.
func (iuo *IndexersUpdateOne) SetURL(s string) *IndexersUpdateOne {
	iuo.mutation.SetURL(s)
	return iuo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (iuo *IndexersUpdateOne) SetNillableURL(s *string) *IndexersUpdateOne {
	if s != nil {
		iuo.SetURL(*s)
	}
	return iuo
}

// ClearURL clears the value of the "url" field.
func (iuo *IndexersUpdateOne) ClearURL() *IndexersUpdateOne {
	iuo.mutation.ClearURL()
	return iuo
}

// SetSynced sets the "synced" field.
func (iuo *IndexersUpdateOne) SetSynced(b bool) *IndexersUpdateOne {
	iuo.mutation.SetSynced(b)
	return iuo
}

// SetNillableSynced sets the "synced" field if the given value is not nil.
func (iuo *IndexersUpdateOne) SetNillableSynced(b *bool) *IndexersUpdateOne {
	if b != nil {
		iuo.SetSynced(*b)
	}
	return iuo
}

// ClearSynced clears the value of the "synced" field.
func (iuo *IndexersUpdateOne) ClearSynced() *IndexersUpdateOne {
	iuo.mutation.ClearSynced()
	return iuo
}

// Mutation returns the IndexersMutation object of the builder.
func (iuo *IndexersUpdateOne) Mutation() *IndexersMutation {
	return iuo.mutation
}

// Where appends a list predicates to the IndexersUpdate builder.
func (iuo *IndexersUpdateOne) Where(ps ...predicate.Indexers) *IndexersUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *IndexersUpdateOne) Select(field string, fields ...string) *IndexersUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Indexers entity.
func (iuo *IndexersUpdateOne) Save(ctx context.Context) (*Indexers, error) {
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *IndexersUpdateOne) SaveX(ctx context.Context) *Indexers {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *IndexersUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *IndexersUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *IndexersUpdateOne) check() error {
	if v, ok := iuo.mutation.Priority(); ok {
		if err := indexers.PriorityValidator(v); err != nil {
			return &ValidationError{Name: "priority", err: fmt.Errorf(`ent: validator failed for field "Indexers.priority": %w`, err)}
		}
	}
	return nil
}

func (iuo *IndexersUpdateOne) sqlSave(ctx context.Context) (_node *Indexers, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(indexers.Table, indexers.Columns, sqlgraph.NewFieldSpec(indexers.FieldID, field.TypeInt))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Indexers.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, indexers.FieldID)
		for _, f := range fields {
			if !indexers.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != indexers.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.Name(); ok {
		_spec.SetField(indexers.FieldName, field.TypeString, value)
	}
	if value, ok := iuo.mutation.Implementation(); ok {
		_spec.SetField(indexers.FieldImplementation, field.TypeString, value)
	}
	if value, ok := iuo.mutation.Settings(); ok {
		_spec.SetField(indexers.FieldSettings, field.TypeString, value)
	}
	if iuo.mutation.SettingsCleared() {
		_spec.ClearField(indexers.FieldSettings, field.TypeString)
	}
	if value, ok := iuo.mutation.EnableRss(); ok {
		_spec.SetField(indexers.FieldEnableRss, field.TypeBool, value)
	}
	if value, ok := iuo.mutation.Priority(); ok {
		_spec.SetField(indexers.FieldPriority, field.TypeInt, value)
	}
	if value, ok := iuo.mutation.AddedPriority(); ok {
		_spec.AddField(indexers.FieldPriority, field.TypeInt, value)
	}
	if value, ok := iuo.mutation.SeedRatio(); ok {
		_spec.SetField(indexers.FieldSeedRatio, field.TypeFloat32, value)
	}
	if value, ok := iuo.mutation.AddedSeedRatio(); ok {
		_spec.AddField(indexers.FieldSeedRatio, field.TypeFloat32, value)
	}
	if iuo.mutation.SeedRatioCleared() {
		_spec.ClearField(indexers.FieldSeedRatio, field.TypeFloat32)
	}
	if value, ok := iuo.mutation.Disabled(); ok {
		_spec.SetField(indexers.FieldDisabled, field.TypeBool, value)
	}
	if iuo.mutation.DisabledCleared() {
		_spec.ClearField(indexers.FieldDisabled, field.TypeBool)
	}
	if value, ok := iuo.mutation.TvSearch(); ok {
		_spec.SetField(indexers.FieldTvSearch, field.TypeBool, value)
	}
	if iuo.mutation.TvSearchCleared() {
		_spec.ClearField(indexers.FieldTvSearch, field.TypeBool)
	}
	if value, ok := iuo.mutation.MovieSearch(); ok {
		_spec.SetField(indexers.FieldMovieSearch, field.TypeBool, value)
	}
	if iuo.mutation.MovieSearchCleared() {
		_spec.ClearField(indexers.FieldMovieSearch, field.TypeBool)
	}
	if value, ok := iuo.mutation.APIKey(); ok {
		_spec.SetField(indexers.FieldAPIKey, field.TypeString, value)
	}
	if iuo.mutation.APIKeyCleared() {
		_spec.ClearField(indexers.FieldAPIKey, field.TypeString)
	}
	if value, ok := iuo.mutation.URL(); ok {
		_spec.SetField(indexers.FieldURL, field.TypeString, value)
	}
	if iuo.mutation.URLCleared() {
		_spec.ClearField(indexers.FieldURL, field.TypeString)
	}
	if value, ok := iuo.mutation.Synced(); ok {
		_spec.SetField(indexers.FieldSynced, field.TypeBool, value)
	}
	if iuo.mutation.SyncedCleared() {
		_spec.ClearField(indexers.FieldSynced, field.TypeBool)
	}
	if iuo.mutation.CreateTimeCleared() {
		_spec.ClearField(indexers.FieldCreateTime, field.TypeTime)
	}
	_node = &Indexers{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{indexers.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
