// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/phoneverificationcode"
	"github.com/mikestefanello/pagoda/ent/predicate"
	"github.com/mikestefanello/pagoda/ent/profile"
)

// PhoneVerificationCodeUpdate is the builder for updating PhoneVerificationCode entities.
type PhoneVerificationCodeUpdate struct {
	config
	hooks    []Hook
	mutation *PhoneVerificationCodeMutation
}

// Where appends a list predicates to the PhoneVerificationCodeUpdate builder.
func (pvcu *PhoneVerificationCodeUpdate) Where(ps ...predicate.PhoneVerificationCode) *PhoneVerificationCodeUpdate {
	pvcu.mutation.Where(ps...)
	return pvcu
}

// SetUpdatedAt sets the "updated_at" field.
func (pvcu *PhoneVerificationCodeUpdate) SetUpdatedAt(t time.Time) *PhoneVerificationCodeUpdate {
	pvcu.mutation.SetUpdatedAt(t)
	return pvcu
}

// SetCode sets the "code" field.
func (pvcu *PhoneVerificationCodeUpdate) SetCode(s string) *PhoneVerificationCodeUpdate {
	pvcu.mutation.SetCode(s)
	return pvcu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (pvcu *PhoneVerificationCodeUpdate) SetNillableCode(s *string) *PhoneVerificationCodeUpdate {
	if s != nil {
		pvcu.SetCode(*s)
	}
	return pvcu
}

// SetProfileID sets the "profile_id" field.
func (pvcu *PhoneVerificationCodeUpdate) SetProfileID(i int) *PhoneVerificationCodeUpdate {
	pvcu.mutation.SetProfileID(i)
	return pvcu
}

// SetNillableProfileID sets the "profile_id" field if the given value is not nil.
func (pvcu *PhoneVerificationCodeUpdate) SetNillableProfileID(i *int) *PhoneVerificationCodeUpdate {
	if i != nil {
		pvcu.SetProfileID(*i)
	}
	return pvcu
}

// SetProfile sets the "profile" edge to the Profile entity.
func (pvcu *PhoneVerificationCodeUpdate) SetProfile(p *Profile) *PhoneVerificationCodeUpdate {
	return pvcu.SetProfileID(p.ID)
}

// Mutation returns the PhoneVerificationCodeMutation object of the builder.
func (pvcu *PhoneVerificationCodeUpdate) Mutation() *PhoneVerificationCodeMutation {
	return pvcu.mutation
}

// ClearProfile clears the "profile" edge to the Profile entity.
func (pvcu *PhoneVerificationCodeUpdate) ClearProfile() *PhoneVerificationCodeUpdate {
	pvcu.mutation.ClearProfile()
	return pvcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pvcu *PhoneVerificationCodeUpdate) Save(ctx context.Context) (int, error) {
	pvcu.defaults()
	return withHooks(ctx, pvcu.sqlSave, pvcu.mutation, pvcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pvcu *PhoneVerificationCodeUpdate) SaveX(ctx context.Context) int {
	affected, err := pvcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pvcu *PhoneVerificationCodeUpdate) Exec(ctx context.Context) error {
	_, err := pvcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pvcu *PhoneVerificationCodeUpdate) ExecX(ctx context.Context) {
	if err := pvcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pvcu *PhoneVerificationCodeUpdate) defaults() {
	if _, ok := pvcu.mutation.UpdatedAt(); !ok {
		v := phoneverificationcode.UpdateDefaultUpdatedAt()
		pvcu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pvcu *PhoneVerificationCodeUpdate) check() error {
	if pvcu.mutation.ProfileCleared() && len(pvcu.mutation.ProfileIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "PhoneVerificationCode.profile"`)
	}
	return nil
}

func (pvcu *PhoneVerificationCodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pvcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(phoneverificationcode.Table, phoneverificationcode.Columns, sqlgraph.NewFieldSpec(phoneverificationcode.FieldID, field.TypeInt))
	if ps := pvcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pvcu.mutation.UpdatedAt(); ok {
		_spec.SetField(phoneverificationcode.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pvcu.mutation.Code(); ok {
		_spec.SetField(phoneverificationcode.FieldCode, field.TypeString, value)
	}
	if pvcu.mutation.ProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   phoneverificationcode.ProfileTable,
			Columns: []string{phoneverificationcode.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvcu.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   phoneverificationcode.ProfileTable,
			Columns: []string{phoneverificationcode.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pvcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{phoneverificationcode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pvcu.mutation.done = true
	return n, nil
}

// PhoneVerificationCodeUpdateOne is the builder for updating a single PhoneVerificationCode entity.
type PhoneVerificationCodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PhoneVerificationCodeMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (pvcuo *PhoneVerificationCodeUpdateOne) SetUpdatedAt(t time.Time) *PhoneVerificationCodeUpdateOne {
	pvcuo.mutation.SetUpdatedAt(t)
	return pvcuo
}

// SetCode sets the "code" field.
func (pvcuo *PhoneVerificationCodeUpdateOne) SetCode(s string) *PhoneVerificationCodeUpdateOne {
	pvcuo.mutation.SetCode(s)
	return pvcuo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (pvcuo *PhoneVerificationCodeUpdateOne) SetNillableCode(s *string) *PhoneVerificationCodeUpdateOne {
	if s != nil {
		pvcuo.SetCode(*s)
	}
	return pvcuo
}

// SetProfileID sets the "profile_id" field.
func (pvcuo *PhoneVerificationCodeUpdateOne) SetProfileID(i int) *PhoneVerificationCodeUpdateOne {
	pvcuo.mutation.SetProfileID(i)
	return pvcuo
}

// SetNillableProfileID sets the "profile_id" field if the given value is not nil.
func (pvcuo *PhoneVerificationCodeUpdateOne) SetNillableProfileID(i *int) *PhoneVerificationCodeUpdateOne {
	if i != nil {
		pvcuo.SetProfileID(*i)
	}
	return pvcuo
}

// SetProfile sets the "profile" edge to the Profile entity.
func (pvcuo *PhoneVerificationCodeUpdateOne) SetProfile(p *Profile) *PhoneVerificationCodeUpdateOne {
	return pvcuo.SetProfileID(p.ID)
}

// Mutation returns the PhoneVerificationCodeMutation object of the builder.
func (pvcuo *PhoneVerificationCodeUpdateOne) Mutation() *PhoneVerificationCodeMutation {
	return pvcuo.mutation
}

// ClearProfile clears the "profile" edge to the Profile entity.
func (pvcuo *PhoneVerificationCodeUpdateOne) ClearProfile() *PhoneVerificationCodeUpdateOne {
	pvcuo.mutation.ClearProfile()
	return pvcuo
}

// Where appends a list predicates to the PhoneVerificationCodeUpdate builder.
func (pvcuo *PhoneVerificationCodeUpdateOne) Where(ps ...predicate.PhoneVerificationCode) *PhoneVerificationCodeUpdateOne {
	pvcuo.mutation.Where(ps...)
	return pvcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pvcuo *PhoneVerificationCodeUpdateOne) Select(field string, fields ...string) *PhoneVerificationCodeUpdateOne {
	pvcuo.fields = append([]string{field}, fields...)
	return pvcuo
}

// Save executes the query and returns the updated PhoneVerificationCode entity.
func (pvcuo *PhoneVerificationCodeUpdateOne) Save(ctx context.Context) (*PhoneVerificationCode, error) {
	pvcuo.defaults()
	return withHooks(ctx, pvcuo.sqlSave, pvcuo.mutation, pvcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pvcuo *PhoneVerificationCodeUpdateOne) SaveX(ctx context.Context) *PhoneVerificationCode {
	node, err := pvcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pvcuo *PhoneVerificationCodeUpdateOne) Exec(ctx context.Context) error {
	_, err := pvcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pvcuo *PhoneVerificationCodeUpdateOne) ExecX(ctx context.Context) {
	if err := pvcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pvcuo *PhoneVerificationCodeUpdateOne) defaults() {
	if _, ok := pvcuo.mutation.UpdatedAt(); !ok {
		v := phoneverificationcode.UpdateDefaultUpdatedAt()
		pvcuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pvcuo *PhoneVerificationCodeUpdateOne) check() error {
	if pvcuo.mutation.ProfileCleared() && len(pvcuo.mutation.ProfileIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "PhoneVerificationCode.profile"`)
	}
	return nil
}

func (pvcuo *PhoneVerificationCodeUpdateOne) sqlSave(ctx context.Context) (_node *PhoneVerificationCode, err error) {
	if err := pvcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(phoneverificationcode.Table, phoneverificationcode.Columns, sqlgraph.NewFieldSpec(phoneverificationcode.FieldID, field.TypeInt))
	id, ok := pvcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PhoneVerificationCode.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pvcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, phoneverificationcode.FieldID)
		for _, f := range fields {
			if !phoneverificationcode.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != phoneverificationcode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pvcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pvcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(phoneverificationcode.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pvcuo.mutation.Code(); ok {
		_spec.SetField(phoneverificationcode.FieldCode, field.TypeString, value)
	}
	if pvcuo.mutation.ProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   phoneverificationcode.ProfileTable,
			Columns: []string{phoneverificationcode.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvcuo.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   phoneverificationcode.ProfileTable,
			Columns: []string{phoneverificationcode.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PhoneVerificationCode{config: pvcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pvcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{phoneverificationcode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pvcuo.mutation.done = true
	return _node, nil
}
