// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/imagesize"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// ImageSizeDelete is the builder for deleting a ImageSize entity.
type ImageSizeDelete struct {
	config
	hooks    []Hook
	mutation *ImageSizeMutation
}

// Where appends a list predicates to the ImageSizeDelete builder.
func (isd *ImageSizeDelete) Where(ps ...predicate.ImageSize) *ImageSizeDelete {
	isd.mutation.Where(ps...)
	return isd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (isd *ImageSizeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, isd.sqlExec, isd.mutation, isd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (isd *ImageSizeDelete) ExecX(ctx context.Context) int {
	n, err := isd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (isd *ImageSizeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(imagesize.Table, sqlgraph.NewFieldSpec(imagesize.FieldID, field.TypeInt))
	if ps := isd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, isd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	isd.mutation.done = true
	return affected, err
}

// ImageSizeDeleteOne is the builder for deleting a single ImageSize entity.
type ImageSizeDeleteOne struct {
	isd *ImageSizeDelete
}

// Where appends a list predicates to the ImageSizeDelete builder.
func (isdo *ImageSizeDeleteOne) Where(ps ...predicate.ImageSize) *ImageSizeDeleteOne {
	isdo.isd.mutation.Where(ps...)
	return isdo
}

// Exec executes the deletion query.
func (isdo *ImageSizeDeleteOne) Exec(ctx context.Context) error {
	n, err := isdo.isd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{imagesize.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (isdo *ImageSizeDeleteOne) ExecX(ctx context.Context) {
	if err := isdo.Exec(ctx); err != nil {
		panic(err)
	}
}