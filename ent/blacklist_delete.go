// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"polaris/ent/blacklist"
	"polaris/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BlacklistDelete is the builder for deleting a Blacklist entity.
type BlacklistDelete struct {
	config
	hooks    []Hook
	mutation *BlacklistMutation
}

// Where appends a list predicates to the BlacklistDelete builder.
func (bd *BlacklistDelete) Where(ps ...predicate.Blacklist) *BlacklistDelete {
	bd.mutation.Where(ps...)
	return bd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bd *BlacklistDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bd.sqlExec, bd.mutation, bd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bd *BlacklistDelete) ExecX(ctx context.Context) int {
	n, err := bd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bd *BlacklistDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(blacklist.Table, sqlgraph.NewFieldSpec(blacklist.FieldID, field.TypeInt))
	if ps := bd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bd.mutation.done = true
	return affected, err
}

// BlacklistDeleteOne is the builder for deleting a single Blacklist entity.
type BlacklistDeleteOne struct {
	bd *BlacklistDelete
}

// Where appends a list predicates to the BlacklistDelete builder.
func (bdo *BlacklistDeleteOne) Where(ps ...predicate.Blacklist) *BlacklistDeleteOne {
	bdo.bd.mutation.Where(ps...)
	return bdo
}

// Exec executes the deletion query.
func (bdo *BlacklistDeleteOne) Exec(ctx context.Context) error {
	n, err := bdo.bd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{blacklist.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bdo *BlacklistDeleteOne) ExecX(ctx context.Context) {
	if err := bdo.Exec(ctx); err != nil {
		panic(err)
	}
}