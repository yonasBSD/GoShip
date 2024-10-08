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
	"github.com/mikestefanello/pagoda/ent/emailsubscription"
	"github.com/mikestefanello/pagoda/ent/emailsubscriptiontype"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// EmailSubscriptionUpdate is the builder for updating EmailSubscription entities.
type EmailSubscriptionUpdate struct {
	config
	hooks    []Hook
	mutation *EmailSubscriptionMutation
}

// Where appends a list predicates to the EmailSubscriptionUpdate builder.
func (esu *EmailSubscriptionUpdate) Where(ps ...predicate.EmailSubscription) *EmailSubscriptionUpdate {
	esu.mutation.Where(ps...)
	return esu
}

// SetUpdatedAt sets the "updated_at" field.
func (esu *EmailSubscriptionUpdate) SetUpdatedAt(t time.Time) *EmailSubscriptionUpdate {
	esu.mutation.SetUpdatedAt(t)
	return esu
}

// SetEmail sets the "email" field.
func (esu *EmailSubscriptionUpdate) SetEmail(s string) *EmailSubscriptionUpdate {
	esu.mutation.SetEmail(s)
	return esu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (esu *EmailSubscriptionUpdate) SetNillableEmail(s *string) *EmailSubscriptionUpdate {
	if s != nil {
		esu.SetEmail(*s)
	}
	return esu
}

// SetVerified sets the "verified" field.
func (esu *EmailSubscriptionUpdate) SetVerified(b bool) *EmailSubscriptionUpdate {
	esu.mutation.SetVerified(b)
	return esu
}

// SetNillableVerified sets the "verified" field if the given value is not nil.
func (esu *EmailSubscriptionUpdate) SetNillableVerified(b *bool) *EmailSubscriptionUpdate {
	if b != nil {
		esu.SetVerified(*b)
	}
	return esu
}

// SetConfirmationCode sets the "confirmation_code" field.
func (esu *EmailSubscriptionUpdate) SetConfirmationCode(s string) *EmailSubscriptionUpdate {
	esu.mutation.SetConfirmationCode(s)
	return esu
}

// SetNillableConfirmationCode sets the "confirmation_code" field if the given value is not nil.
func (esu *EmailSubscriptionUpdate) SetNillableConfirmationCode(s *string) *EmailSubscriptionUpdate {
	if s != nil {
		esu.SetConfirmationCode(*s)
	}
	return esu
}

// SetLatitude sets the "latitude" field.
func (esu *EmailSubscriptionUpdate) SetLatitude(f float64) *EmailSubscriptionUpdate {
	esu.mutation.ResetLatitude()
	esu.mutation.SetLatitude(f)
	return esu
}

// SetNillableLatitude sets the "latitude" field if the given value is not nil.
func (esu *EmailSubscriptionUpdate) SetNillableLatitude(f *float64) *EmailSubscriptionUpdate {
	if f != nil {
		esu.SetLatitude(*f)
	}
	return esu
}

// AddLatitude adds f to the "latitude" field.
func (esu *EmailSubscriptionUpdate) AddLatitude(f float64) *EmailSubscriptionUpdate {
	esu.mutation.AddLatitude(f)
	return esu
}

// ClearLatitude clears the value of the "latitude" field.
func (esu *EmailSubscriptionUpdate) ClearLatitude() *EmailSubscriptionUpdate {
	esu.mutation.ClearLatitude()
	return esu
}

// SetLongitude sets the "longitude" field.
func (esu *EmailSubscriptionUpdate) SetLongitude(f float64) *EmailSubscriptionUpdate {
	esu.mutation.ResetLongitude()
	esu.mutation.SetLongitude(f)
	return esu
}

// SetNillableLongitude sets the "longitude" field if the given value is not nil.
func (esu *EmailSubscriptionUpdate) SetNillableLongitude(f *float64) *EmailSubscriptionUpdate {
	if f != nil {
		esu.SetLongitude(*f)
	}
	return esu
}

// AddLongitude adds f to the "longitude" field.
func (esu *EmailSubscriptionUpdate) AddLongitude(f float64) *EmailSubscriptionUpdate {
	esu.mutation.AddLongitude(f)
	return esu
}

// ClearLongitude clears the value of the "longitude" field.
func (esu *EmailSubscriptionUpdate) ClearLongitude() *EmailSubscriptionUpdate {
	esu.mutation.ClearLongitude()
	return esu
}

// AddSubscriptionIDs adds the "subscriptions" edge to the EmailSubscriptionType entity by IDs.
func (esu *EmailSubscriptionUpdate) AddSubscriptionIDs(ids ...int) *EmailSubscriptionUpdate {
	esu.mutation.AddSubscriptionIDs(ids...)
	return esu
}

// AddSubscriptions adds the "subscriptions" edges to the EmailSubscriptionType entity.
func (esu *EmailSubscriptionUpdate) AddSubscriptions(e ...*EmailSubscriptionType) *EmailSubscriptionUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return esu.AddSubscriptionIDs(ids...)
}

// Mutation returns the EmailSubscriptionMutation object of the builder.
func (esu *EmailSubscriptionUpdate) Mutation() *EmailSubscriptionMutation {
	return esu.mutation
}

// ClearSubscriptions clears all "subscriptions" edges to the EmailSubscriptionType entity.
func (esu *EmailSubscriptionUpdate) ClearSubscriptions() *EmailSubscriptionUpdate {
	esu.mutation.ClearSubscriptions()
	return esu
}

// RemoveSubscriptionIDs removes the "subscriptions" edge to EmailSubscriptionType entities by IDs.
func (esu *EmailSubscriptionUpdate) RemoveSubscriptionIDs(ids ...int) *EmailSubscriptionUpdate {
	esu.mutation.RemoveSubscriptionIDs(ids...)
	return esu
}

// RemoveSubscriptions removes "subscriptions" edges to EmailSubscriptionType entities.
func (esu *EmailSubscriptionUpdate) RemoveSubscriptions(e ...*EmailSubscriptionType) *EmailSubscriptionUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return esu.RemoveSubscriptionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (esu *EmailSubscriptionUpdate) Save(ctx context.Context) (int, error) {
	esu.defaults()
	return withHooks(ctx, esu.sqlSave, esu.mutation, esu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (esu *EmailSubscriptionUpdate) SaveX(ctx context.Context) int {
	affected, err := esu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (esu *EmailSubscriptionUpdate) Exec(ctx context.Context) error {
	_, err := esu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (esu *EmailSubscriptionUpdate) ExecX(ctx context.Context) {
	if err := esu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (esu *EmailSubscriptionUpdate) defaults() {
	if _, ok := esu.mutation.UpdatedAt(); !ok {
		v := emailsubscription.UpdateDefaultUpdatedAt()
		esu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (esu *EmailSubscriptionUpdate) check() error {
	if v, ok := esu.mutation.Email(); ok {
		if err := emailsubscription.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "EmailSubscription.email": %w`, err)}
		}
	}
	if v, ok := esu.mutation.ConfirmationCode(); ok {
		if err := emailsubscription.ConfirmationCodeValidator(v); err != nil {
			return &ValidationError{Name: "confirmation_code", err: fmt.Errorf(`ent: validator failed for field "EmailSubscription.confirmation_code": %w`, err)}
		}
	}
	return nil
}

func (esu *EmailSubscriptionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := esu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(emailsubscription.Table, emailsubscription.Columns, sqlgraph.NewFieldSpec(emailsubscription.FieldID, field.TypeInt))
	if ps := esu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := esu.mutation.UpdatedAt(); ok {
		_spec.SetField(emailsubscription.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := esu.mutation.Email(); ok {
		_spec.SetField(emailsubscription.FieldEmail, field.TypeString, value)
	}
	if value, ok := esu.mutation.Verified(); ok {
		_spec.SetField(emailsubscription.FieldVerified, field.TypeBool, value)
	}
	if value, ok := esu.mutation.ConfirmationCode(); ok {
		_spec.SetField(emailsubscription.FieldConfirmationCode, field.TypeString, value)
	}
	if value, ok := esu.mutation.Latitude(); ok {
		_spec.SetField(emailsubscription.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := esu.mutation.AddedLatitude(); ok {
		_spec.AddField(emailsubscription.FieldLatitude, field.TypeFloat64, value)
	}
	if esu.mutation.LatitudeCleared() {
		_spec.ClearField(emailsubscription.FieldLatitude, field.TypeFloat64)
	}
	if value, ok := esu.mutation.Longitude(); ok {
		_spec.SetField(emailsubscription.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := esu.mutation.AddedLongitude(); ok {
		_spec.AddField(emailsubscription.FieldLongitude, field.TypeFloat64, value)
	}
	if esu.mutation.LongitudeCleared() {
		_spec.ClearField(emailsubscription.FieldLongitude, field.TypeFloat64)
	}
	if esu.mutation.SubscriptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   emailsubscription.SubscriptionsTable,
			Columns: emailsubscription.SubscriptionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(emailsubscriptiontype.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := esu.mutation.RemovedSubscriptionsIDs(); len(nodes) > 0 && !esu.mutation.SubscriptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   emailsubscription.SubscriptionsTable,
			Columns: emailsubscription.SubscriptionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(emailsubscriptiontype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := esu.mutation.SubscriptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   emailsubscription.SubscriptionsTable,
			Columns: emailsubscription.SubscriptionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(emailsubscriptiontype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, esu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{emailsubscription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	esu.mutation.done = true
	return n, nil
}

// EmailSubscriptionUpdateOne is the builder for updating a single EmailSubscription entity.
type EmailSubscriptionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EmailSubscriptionMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (esuo *EmailSubscriptionUpdateOne) SetUpdatedAt(t time.Time) *EmailSubscriptionUpdateOne {
	esuo.mutation.SetUpdatedAt(t)
	return esuo
}

// SetEmail sets the "email" field.
func (esuo *EmailSubscriptionUpdateOne) SetEmail(s string) *EmailSubscriptionUpdateOne {
	esuo.mutation.SetEmail(s)
	return esuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (esuo *EmailSubscriptionUpdateOne) SetNillableEmail(s *string) *EmailSubscriptionUpdateOne {
	if s != nil {
		esuo.SetEmail(*s)
	}
	return esuo
}

// SetVerified sets the "verified" field.
func (esuo *EmailSubscriptionUpdateOne) SetVerified(b bool) *EmailSubscriptionUpdateOne {
	esuo.mutation.SetVerified(b)
	return esuo
}

// SetNillableVerified sets the "verified" field if the given value is not nil.
func (esuo *EmailSubscriptionUpdateOne) SetNillableVerified(b *bool) *EmailSubscriptionUpdateOne {
	if b != nil {
		esuo.SetVerified(*b)
	}
	return esuo
}

// SetConfirmationCode sets the "confirmation_code" field.
func (esuo *EmailSubscriptionUpdateOne) SetConfirmationCode(s string) *EmailSubscriptionUpdateOne {
	esuo.mutation.SetConfirmationCode(s)
	return esuo
}

// SetNillableConfirmationCode sets the "confirmation_code" field if the given value is not nil.
func (esuo *EmailSubscriptionUpdateOne) SetNillableConfirmationCode(s *string) *EmailSubscriptionUpdateOne {
	if s != nil {
		esuo.SetConfirmationCode(*s)
	}
	return esuo
}

// SetLatitude sets the "latitude" field.
func (esuo *EmailSubscriptionUpdateOne) SetLatitude(f float64) *EmailSubscriptionUpdateOne {
	esuo.mutation.ResetLatitude()
	esuo.mutation.SetLatitude(f)
	return esuo
}

// SetNillableLatitude sets the "latitude" field if the given value is not nil.
func (esuo *EmailSubscriptionUpdateOne) SetNillableLatitude(f *float64) *EmailSubscriptionUpdateOne {
	if f != nil {
		esuo.SetLatitude(*f)
	}
	return esuo
}

// AddLatitude adds f to the "latitude" field.
func (esuo *EmailSubscriptionUpdateOne) AddLatitude(f float64) *EmailSubscriptionUpdateOne {
	esuo.mutation.AddLatitude(f)
	return esuo
}

// ClearLatitude clears the value of the "latitude" field.
func (esuo *EmailSubscriptionUpdateOne) ClearLatitude() *EmailSubscriptionUpdateOne {
	esuo.mutation.ClearLatitude()
	return esuo
}

// SetLongitude sets the "longitude" field.
func (esuo *EmailSubscriptionUpdateOne) SetLongitude(f float64) *EmailSubscriptionUpdateOne {
	esuo.mutation.ResetLongitude()
	esuo.mutation.SetLongitude(f)
	return esuo
}

// SetNillableLongitude sets the "longitude" field if the given value is not nil.
func (esuo *EmailSubscriptionUpdateOne) SetNillableLongitude(f *float64) *EmailSubscriptionUpdateOne {
	if f != nil {
		esuo.SetLongitude(*f)
	}
	return esuo
}

// AddLongitude adds f to the "longitude" field.
func (esuo *EmailSubscriptionUpdateOne) AddLongitude(f float64) *EmailSubscriptionUpdateOne {
	esuo.mutation.AddLongitude(f)
	return esuo
}

// ClearLongitude clears the value of the "longitude" field.
func (esuo *EmailSubscriptionUpdateOne) ClearLongitude() *EmailSubscriptionUpdateOne {
	esuo.mutation.ClearLongitude()
	return esuo
}

// AddSubscriptionIDs adds the "subscriptions" edge to the EmailSubscriptionType entity by IDs.
func (esuo *EmailSubscriptionUpdateOne) AddSubscriptionIDs(ids ...int) *EmailSubscriptionUpdateOne {
	esuo.mutation.AddSubscriptionIDs(ids...)
	return esuo
}

// AddSubscriptions adds the "subscriptions" edges to the EmailSubscriptionType entity.
func (esuo *EmailSubscriptionUpdateOne) AddSubscriptions(e ...*EmailSubscriptionType) *EmailSubscriptionUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return esuo.AddSubscriptionIDs(ids...)
}

// Mutation returns the EmailSubscriptionMutation object of the builder.
func (esuo *EmailSubscriptionUpdateOne) Mutation() *EmailSubscriptionMutation {
	return esuo.mutation
}

// ClearSubscriptions clears all "subscriptions" edges to the EmailSubscriptionType entity.
func (esuo *EmailSubscriptionUpdateOne) ClearSubscriptions() *EmailSubscriptionUpdateOne {
	esuo.mutation.ClearSubscriptions()
	return esuo
}

// RemoveSubscriptionIDs removes the "subscriptions" edge to EmailSubscriptionType entities by IDs.
func (esuo *EmailSubscriptionUpdateOne) RemoveSubscriptionIDs(ids ...int) *EmailSubscriptionUpdateOne {
	esuo.mutation.RemoveSubscriptionIDs(ids...)
	return esuo
}

// RemoveSubscriptions removes "subscriptions" edges to EmailSubscriptionType entities.
func (esuo *EmailSubscriptionUpdateOne) RemoveSubscriptions(e ...*EmailSubscriptionType) *EmailSubscriptionUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return esuo.RemoveSubscriptionIDs(ids...)
}

// Where appends a list predicates to the EmailSubscriptionUpdate builder.
func (esuo *EmailSubscriptionUpdateOne) Where(ps ...predicate.EmailSubscription) *EmailSubscriptionUpdateOne {
	esuo.mutation.Where(ps...)
	return esuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (esuo *EmailSubscriptionUpdateOne) Select(field string, fields ...string) *EmailSubscriptionUpdateOne {
	esuo.fields = append([]string{field}, fields...)
	return esuo
}

// Save executes the query and returns the updated EmailSubscription entity.
func (esuo *EmailSubscriptionUpdateOne) Save(ctx context.Context) (*EmailSubscription, error) {
	esuo.defaults()
	return withHooks(ctx, esuo.sqlSave, esuo.mutation, esuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (esuo *EmailSubscriptionUpdateOne) SaveX(ctx context.Context) *EmailSubscription {
	node, err := esuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (esuo *EmailSubscriptionUpdateOne) Exec(ctx context.Context) error {
	_, err := esuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (esuo *EmailSubscriptionUpdateOne) ExecX(ctx context.Context) {
	if err := esuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (esuo *EmailSubscriptionUpdateOne) defaults() {
	if _, ok := esuo.mutation.UpdatedAt(); !ok {
		v := emailsubscription.UpdateDefaultUpdatedAt()
		esuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (esuo *EmailSubscriptionUpdateOne) check() error {
	if v, ok := esuo.mutation.Email(); ok {
		if err := emailsubscription.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "EmailSubscription.email": %w`, err)}
		}
	}
	if v, ok := esuo.mutation.ConfirmationCode(); ok {
		if err := emailsubscription.ConfirmationCodeValidator(v); err != nil {
			return &ValidationError{Name: "confirmation_code", err: fmt.Errorf(`ent: validator failed for field "EmailSubscription.confirmation_code": %w`, err)}
		}
	}
	return nil
}

func (esuo *EmailSubscriptionUpdateOne) sqlSave(ctx context.Context) (_node *EmailSubscription, err error) {
	if err := esuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(emailsubscription.Table, emailsubscription.Columns, sqlgraph.NewFieldSpec(emailsubscription.FieldID, field.TypeInt))
	id, ok := esuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EmailSubscription.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := esuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, emailsubscription.FieldID)
		for _, f := range fields {
			if !emailsubscription.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != emailsubscription.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := esuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := esuo.mutation.UpdatedAt(); ok {
		_spec.SetField(emailsubscription.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := esuo.mutation.Email(); ok {
		_spec.SetField(emailsubscription.FieldEmail, field.TypeString, value)
	}
	if value, ok := esuo.mutation.Verified(); ok {
		_spec.SetField(emailsubscription.FieldVerified, field.TypeBool, value)
	}
	if value, ok := esuo.mutation.ConfirmationCode(); ok {
		_spec.SetField(emailsubscription.FieldConfirmationCode, field.TypeString, value)
	}
	if value, ok := esuo.mutation.Latitude(); ok {
		_spec.SetField(emailsubscription.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := esuo.mutation.AddedLatitude(); ok {
		_spec.AddField(emailsubscription.FieldLatitude, field.TypeFloat64, value)
	}
	if esuo.mutation.LatitudeCleared() {
		_spec.ClearField(emailsubscription.FieldLatitude, field.TypeFloat64)
	}
	if value, ok := esuo.mutation.Longitude(); ok {
		_spec.SetField(emailsubscription.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := esuo.mutation.AddedLongitude(); ok {
		_spec.AddField(emailsubscription.FieldLongitude, field.TypeFloat64, value)
	}
	if esuo.mutation.LongitudeCleared() {
		_spec.ClearField(emailsubscription.FieldLongitude, field.TypeFloat64)
	}
	if esuo.mutation.SubscriptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   emailsubscription.SubscriptionsTable,
			Columns: emailsubscription.SubscriptionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(emailsubscriptiontype.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := esuo.mutation.RemovedSubscriptionsIDs(); len(nodes) > 0 && !esuo.mutation.SubscriptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   emailsubscription.SubscriptionsTable,
			Columns: emailsubscription.SubscriptionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(emailsubscriptiontype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := esuo.mutation.SubscriptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   emailsubscription.SubscriptionsTable,
			Columns: emailsubscription.SubscriptionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(emailsubscriptiontype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &EmailSubscription{config: esuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, esuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{emailsubscription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	esuo.mutation.done = true
	return _node, nil
}
