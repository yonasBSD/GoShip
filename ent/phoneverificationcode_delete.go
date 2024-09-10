// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/phoneverificationcode"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// PhoneVerificationCodeDelete is the builder for deleting a PhoneVerificationCode entity.
type PhoneVerificationCodeDelete struct {
	config
	hooks    []Hook
	mutation *PhoneVerificationCodeMutation
}

// Where appends a list predicates to the PhoneVerificationCodeDelete builder.
func (pvcd *PhoneVerificationCodeDelete) Where(ps ...predicate.PhoneVerificationCode) *PhoneVerificationCodeDelete {
	pvcd.mutation.Where(ps...)
	return pvcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pvcd *PhoneVerificationCodeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pvcd.sqlExec, pvcd.mutation, pvcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pvcd *PhoneVerificationCodeDelete) ExecX(ctx context.Context) int {
	n, err := pvcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pvcd *PhoneVerificationCodeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(phoneverificationcode.Table, sqlgraph.NewFieldSpec(phoneverificationcode.FieldID, field.TypeInt))
	if ps := pvcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pvcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pvcd.mutation.done = true
	return affected, err
}

// PhoneVerificationCodeDeleteOne is the builder for deleting a single PhoneVerificationCode entity.
type PhoneVerificationCodeDeleteOne struct {
	pvcd *PhoneVerificationCodeDelete
}

// Where appends a list predicates to the PhoneVerificationCodeDelete builder.
func (pvcdo *PhoneVerificationCodeDeleteOne) Where(ps ...predicate.PhoneVerificationCode) *PhoneVerificationCodeDeleteOne {
	pvcdo.pvcd.mutation.Where(ps...)
	return pvcdo
}

// Exec executes the deletion query.
func (pvcdo *PhoneVerificationCodeDeleteOne) Exec(ctx context.Context) error {
	n, err := pvcdo.pvcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{phoneverificationcode.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pvcdo *PhoneVerificationCodeDeleteOne) ExecX(ctx context.Context) {
	if err := pvcdo.Exec(ctx); err != nil {
		panic(err)
	}
}