// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"polaris/ent/blacklist"
	"polaris/ent/predicate"
	"polaris/ent/schema"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BlacklistUpdate is the builder for updating Blacklist entities.
type BlacklistUpdate struct {
	config
	hooks    []Hook
	mutation *BlacklistMutation
}

// Where appends a list predicates to the BlacklistUpdate builder.
func (bu *BlacklistUpdate) Where(ps ...predicate.Blacklist) *BlacklistUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetType sets the "type" field.
func (bu *BlacklistUpdate) SetType(b blacklist.Type) *BlacklistUpdate {
	bu.mutation.SetType(b)
	return bu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (bu *BlacklistUpdate) SetNillableType(b *blacklist.Type) *BlacklistUpdate {
	if b != nil {
		bu.SetType(*b)
	}
	return bu
}

// SetValue sets the "value" field.
func (bu *BlacklistUpdate) SetValue(sv schema.BlacklistValue) *BlacklistUpdate {
	bu.mutation.SetValue(sv)
	return bu
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (bu *BlacklistUpdate) SetNillableValue(sv *schema.BlacklistValue) *BlacklistUpdate {
	if sv != nil {
		bu.SetValue(*sv)
	}
	return bu
}

// SetNotes sets the "notes" field.
func (bu *BlacklistUpdate) SetNotes(s string) *BlacklistUpdate {
	bu.mutation.SetNotes(s)
	return bu
}

// SetNillableNotes sets the "notes" field if the given value is not nil.
func (bu *BlacklistUpdate) SetNillableNotes(s *string) *BlacklistUpdate {
	if s != nil {
		bu.SetNotes(*s)
	}
	return bu
}

// ClearNotes clears the value of the "notes" field.
func (bu *BlacklistUpdate) ClearNotes() *BlacklistUpdate {
	bu.mutation.ClearNotes()
	return bu
}

// Mutation returns the BlacklistMutation object of the builder.
func (bu *BlacklistUpdate) Mutation() *BlacklistMutation {
	return bu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BlacklistUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BlacklistUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BlacklistUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BlacklistUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BlacklistUpdate) check() error {
	if v, ok := bu.mutation.GetType(); ok {
		if err := blacklist.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Blacklist.type": %w`, err)}
		}
	}
	return nil
}

func (bu *BlacklistUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(blacklist.Table, blacklist.Columns, sqlgraph.NewFieldSpec(blacklist.FieldID, field.TypeInt))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.GetType(); ok {
		_spec.SetField(blacklist.FieldType, field.TypeEnum, value)
	}
	if value, ok := bu.mutation.Value(); ok {
		_spec.SetField(blacklist.FieldValue, field.TypeJSON, value)
	}
	if value, ok := bu.mutation.Notes(); ok {
		_spec.SetField(blacklist.FieldNotes, field.TypeString, value)
	}
	if bu.mutation.NotesCleared() {
		_spec.ClearField(blacklist.FieldNotes, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blacklist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BlacklistUpdateOne is the builder for updating a single Blacklist entity.
type BlacklistUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BlacklistMutation
}

// SetType sets the "type" field.
func (buo *BlacklistUpdateOne) SetType(b blacklist.Type) *BlacklistUpdateOne {
	buo.mutation.SetType(b)
	return buo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (buo *BlacklistUpdateOne) SetNillableType(b *blacklist.Type) *BlacklistUpdateOne {
	if b != nil {
		buo.SetType(*b)
	}
	return buo
}

// SetValue sets the "value" field.
func (buo *BlacklistUpdateOne) SetValue(sv schema.BlacklistValue) *BlacklistUpdateOne {
	buo.mutation.SetValue(sv)
	return buo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (buo *BlacklistUpdateOne) SetNillableValue(sv *schema.BlacklistValue) *BlacklistUpdateOne {
	if sv != nil {
		buo.SetValue(*sv)
	}
	return buo
}

// SetNotes sets the "notes" field.
func (buo *BlacklistUpdateOne) SetNotes(s string) *BlacklistUpdateOne {
	buo.mutation.SetNotes(s)
	return buo
}

// SetNillableNotes sets the "notes" field if the given value is not nil.
func (buo *BlacklistUpdateOne) SetNillableNotes(s *string) *BlacklistUpdateOne {
	if s != nil {
		buo.SetNotes(*s)
	}
	return buo
}

// ClearNotes clears the value of the "notes" field.
func (buo *BlacklistUpdateOne) ClearNotes() *BlacklistUpdateOne {
	buo.mutation.ClearNotes()
	return buo
}

// Mutation returns the BlacklistMutation object of the builder.
func (buo *BlacklistUpdateOne) Mutation() *BlacklistMutation {
	return buo.mutation
}

// Where appends a list predicates to the BlacklistUpdate builder.
func (buo *BlacklistUpdateOne) Where(ps ...predicate.Blacklist) *BlacklistUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BlacklistUpdateOne) Select(field string, fields ...string) *BlacklistUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Blacklist entity.
func (buo *BlacklistUpdateOne) Save(ctx context.Context) (*Blacklist, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BlacklistUpdateOne) SaveX(ctx context.Context) *Blacklist {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BlacklistUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BlacklistUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BlacklistUpdateOne) check() error {
	if v, ok := buo.mutation.GetType(); ok {
		if err := blacklist.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Blacklist.type": %w`, err)}
		}
	}
	return nil
}

func (buo *BlacklistUpdateOne) sqlSave(ctx context.Context) (_node *Blacklist, err error) {
	if err := buo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(blacklist.Table, blacklist.Columns, sqlgraph.NewFieldSpec(blacklist.FieldID, field.TypeInt))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Blacklist.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blacklist.FieldID)
		for _, f := range fields {
			if !blacklist.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != blacklist.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.GetType(); ok {
		_spec.SetField(blacklist.FieldType, field.TypeEnum, value)
	}
	if value, ok := buo.mutation.Value(); ok {
		_spec.SetField(blacklist.FieldValue, field.TypeJSON, value)
	}
	if value, ok := buo.mutation.Notes(); ok {
		_spec.SetField(blacklist.FieldNotes, field.TypeString, value)
	}
	if buo.mutation.NotesCleared() {
		_spec.ClearField(blacklist.FieldNotes, field.TypeString)
	}
	_node = &Blacklist{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blacklist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
