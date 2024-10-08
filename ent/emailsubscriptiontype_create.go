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
)

// EmailSubscriptionTypeCreate is the builder for creating a EmailSubscriptionType entity.
type EmailSubscriptionTypeCreate struct {
	config
	mutation *EmailSubscriptionTypeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (estc *EmailSubscriptionTypeCreate) SetCreatedAt(t time.Time) *EmailSubscriptionTypeCreate {
	estc.mutation.SetCreatedAt(t)
	return estc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (estc *EmailSubscriptionTypeCreate) SetNillableCreatedAt(t *time.Time) *EmailSubscriptionTypeCreate {
	if t != nil {
		estc.SetCreatedAt(*t)
	}
	return estc
}

// SetUpdatedAt sets the "updated_at" field.
func (estc *EmailSubscriptionTypeCreate) SetUpdatedAt(t time.Time) *EmailSubscriptionTypeCreate {
	estc.mutation.SetUpdatedAt(t)
	return estc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (estc *EmailSubscriptionTypeCreate) SetNillableUpdatedAt(t *time.Time) *EmailSubscriptionTypeCreate {
	if t != nil {
		estc.SetUpdatedAt(*t)
	}
	return estc
}

// SetName sets the "name" field.
func (estc *EmailSubscriptionTypeCreate) SetName(e emailsubscriptiontype.Name) *EmailSubscriptionTypeCreate {
	estc.mutation.SetName(e)
	return estc
}

// SetActive sets the "active" field.
func (estc *EmailSubscriptionTypeCreate) SetActive(b bool) *EmailSubscriptionTypeCreate {
	estc.mutation.SetActive(b)
	return estc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (estc *EmailSubscriptionTypeCreate) SetNillableActive(b *bool) *EmailSubscriptionTypeCreate {
	if b != nil {
		estc.SetActive(*b)
	}
	return estc
}

// AddSubscriberIDs adds the "subscriber" edge to the EmailSubscription entity by IDs.
func (estc *EmailSubscriptionTypeCreate) AddSubscriberIDs(ids ...int) *EmailSubscriptionTypeCreate {
	estc.mutation.AddSubscriberIDs(ids...)
	return estc
}

// AddSubscriber adds the "subscriber" edges to the EmailSubscription entity.
func (estc *EmailSubscriptionTypeCreate) AddSubscriber(e ...*EmailSubscription) *EmailSubscriptionTypeCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return estc.AddSubscriberIDs(ids...)
}

// Mutation returns the EmailSubscriptionTypeMutation object of the builder.
func (estc *EmailSubscriptionTypeCreate) Mutation() *EmailSubscriptionTypeMutation {
	return estc.mutation
}

// Save creates the EmailSubscriptionType in the database.
func (estc *EmailSubscriptionTypeCreate) Save(ctx context.Context) (*EmailSubscriptionType, error) {
	estc.defaults()
	return withHooks(ctx, estc.sqlSave, estc.mutation, estc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (estc *EmailSubscriptionTypeCreate) SaveX(ctx context.Context) *EmailSubscriptionType {
	v, err := estc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (estc *EmailSubscriptionTypeCreate) Exec(ctx context.Context) error {
	_, err := estc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (estc *EmailSubscriptionTypeCreate) ExecX(ctx context.Context) {
	if err := estc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (estc *EmailSubscriptionTypeCreate) defaults() {
	if _, ok := estc.mutation.CreatedAt(); !ok {
		v := emailsubscriptiontype.DefaultCreatedAt()
		estc.mutation.SetCreatedAt(v)
	}
	if _, ok := estc.mutation.UpdatedAt(); !ok {
		v := emailsubscriptiontype.DefaultUpdatedAt()
		estc.mutation.SetUpdatedAt(v)
	}
	if _, ok := estc.mutation.Active(); !ok {
		v := emailsubscriptiontype.DefaultActive
		estc.mutation.SetActive(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (estc *EmailSubscriptionTypeCreate) check() error {
	if _, ok := estc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "EmailSubscriptionType.created_at"`)}
	}
	if _, ok := estc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "EmailSubscriptionType.updated_at"`)}
	}
	if _, ok := estc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "EmailSubscriptionType.name"`)}
	}
	if v, ok := estc.mutation.Name(); ok {
		if err := emailsubscriptiontype.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "EmailSubscriptionType.name": %w`, err)}
		}
	}
	if _, ok := estc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`ent: missing required field "EmailSubscriptionType.active"`)}
	}
	return nil
}

func (estc *EmailSubscriptionTypeCreate) sqlSave(ctx context.Context) (*EmailSubscriptionType, error) {
	if err := estc.check(); err != nil {
		return nil, err
	}
	_node, _spec := estc.createSpec()
	if err := sqlgraph.CreateNode(ctx, estc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	estc.mutation.id = &_node.ID
	estc.mutation.done = true
	return _node, nil
}

func (estc *EmailSubscriptionTypeCreate) createSpec() (*EmailSubscriptionType, *sqlgraph.CreateSpec) {
	var (
		_node = &EmailSubscriptionType{config: estc.config}
		_spec = sqlgraph.NewCreateSpec(emailsubscriptiontype.Table, sqlgraph.NewFieldSpec(emailsubscriptiontype.FieldID, field.TypeInt))
	)
	_spec.OnConflict = estc.conflict
	if value, ok := estc.mutation.CreatedAt(); ok {
		_spec.SetField(emailsubscriptiontype.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := estc.mutation.UpdatedAt(); ok {
		_spec.SetField(emailsubscriptiontype.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := estc.mutation.Name(); ok {
		_spec.SetField(emailsubscriptiontype.FieldName, field.TypeEnum, value)
		_node.Name = value
	}
	if value, ok := estc.mutation.Active(); ok {
		_spec.SetField(emailsubscriptiontype.FieldActive, field.TypeBool, value)
		_node.Active = value
	}
	if nodes := estc.mutation.SubscriberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   emailsubscriptiontype.SubscriberTable,
			Columns: emailsubscriptiontype.SubscriberPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(emailsubscription.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.EmailSubscriptionType.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EmailSubscriptionTypeUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (estc *EmailSubscriptionTypeCreate) OnConflict(opts ...sql.ConflictOption) *EmailSubscriptionTypeUpsertOne {
	estc.conflict = opts
	return &EmailSubscriptionTypeUpsertOne{
		create: estc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.EmailSubscriptionType.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (estc *EmailSubscriptionTypeCreate) OnConflictColumns(columns ...string) *EmailSubscriptionTypeUpsertOne {
	estc.conflict = append(estc.conflict, sql.ConflictColumns(columns...))
	return &EmailSubscriptionTypeUpsertOne{
		create: estc,
	}
}

type (
	// EmailSubscriptionTypeUpsertOne is the builder for "upsert"-ing
	//  one EmailSubscriptionType node.
	EmailSubscriptionTypeUpsertOne struct {
		create *EmailSubscriptionTypeCreate
	}

	// EmailSubscriptionTypeUpsert is the "OnConflict" setter.
	EmailSubscriptionTypeUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *EmailSubscriptionTypeUpsert) SetUpdatedAt(v time.Time) *EmailSubscriptionTypeUpsert {
	u.Set(emailsubscriptiontype.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *EmailSubscriptionTypeUpsert) UpdateUpdatedAt() *EmailSubscriptionTypeUpsert {
	u.SetExcluded(emailsubscriptiontype.FieldUpdatedAt)
	return u
}

// SetName sets the "name" field.
func (u *EmailSubscriptionTypeUpsert) SetName(v emailsubscriptiontype.Name) *EmailSubscriptionTypeUpsert {
	u.Set(emailsubscriptiontype.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *EmailSubscriptionTypeUpsert) UpdateName() *EmailSubscriptionTypeUpsert {
	u.SetExcluded(emailsubscriptiontype.FieldName)
	return u
}

// SetActive sets the "active" field.
func (u *EmailSubscriptionTypeUpsert) SetActive(v bool) *EmailSubscriptionTypeUpsert {
	u.Set(emailsubscriptiontype.FieldActive, v)
	return u
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *EmailSubscriptionTypeUpsert) UpdateActive() *EmailSubscriptionTypeUpsert {
	u.SetExcluded(emailsubscriptiontype.FieldActive)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.EmailSubscriptionType.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *EmailSubscriptionTypeUpsertOne) UpdateNewValues() *EmailSubscriptionTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(emailsubscriptiontype.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.EmailSubscriptionType.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *EmailSubscriptionTypeUpsertOne) Ignore() *EmailSubscriptionTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EmailSubscriptionTypeUpsertOne) DoNothing() *EmailSubscriptionTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EmailSubscriptionTypeCreate.OnConflict
// documentation for more info.
func (u *EmailSubscriptionTypeUpsertOne) Update(set func(*EmailSubscriptionTypeUpsert)) *EmailSubscriptionTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EmailSubscriptionTypeUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *EmailSubscriptionTypeUpsertOne) SetUpdatedAt(v time.Time) *EmailSubscriptionTypeUpsertOne {
	return u.Update(func(s *EmailSubscriptionTypeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *EmailSubscriptionTypeUpsertOne) UpdateUpdatedAt() *EmailSubscriptionTypeUpsertOne {
	return u.Update(func(s *EmailSubscriptionTypeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *EmailSubscriptionTypeUpsertOne) SetName(v emailsubscriptiontype.Name) *EmailSubscriptionTypeUpsertOne {
	return u.Update(func(s *EmailSubscriptionTypeUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *EmailSubscriptionTypeUpsertOne) UpdateName() *EmailSubscriptionTypeUpsertOne {
	return u.Update(func(s *EmailSubscriptionTypeUpsert) {
		s.UpdateName()
	})
}

// SetActive sets the "active" field.
func (u *EmailSubscriptionTypeUpsertOne) SetActive(v bool) *EmailSubscriptionTypeUpsertOne {
	return u.Update(func(s *EmailSubscriptionTypeUpsert) {
		s.SetActive(v)
	})
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *EmailSubscriptionTypeUpsertOne) UpdateActive() *EmailSubscriptionTypeUpsertOne {
	return u.Update(func(s *EmailSubscriptionTypeUpsert) {
		s.UpdateActive()
	})
}

// Exec executes the query.
func (u *EmailSubscriptionTypeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EmailSubscriptionTypeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EmailSubscriptionTypeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *EmailSubscriptionTypeUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *EmailSubscriptionTypeUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// EmailSubscriptionTypeCreateBulk is the builder for creating many EmailSubscriptionType entities in bulk.
type EmailSubscriptionTypeCreateBulk struct {
	config
	err      error
	builders []*EmailSubscriptionTypeCreate
	conflict []sql.ConflictOption
}

// Save creates the EmailSubscriptionType entities in the database.
func (estcb *EmailSubscriptionTypeCreateBulk) Save(ctx context.Context) ([]*EmailSubscriptionType, error) {
	if estcb.err != nil {
		return nil, estcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(estcb.builders))
	nodes := make([]*EmailSubscriptionType, len(estcb.builders))
	mutators := make([]Mutator, len(estcb.builders))
	for i := range estcb.builders {
		func(i int, root context.Context) {
			builder := estcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EmailSubscriptionTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, estcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = estcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, estcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, estcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (estcb *EmailSubscriptionTypeCreateBulk) SaveX(ctx context.Context) []*EmailSubscriptionType {
	v, err := estcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (estcb *EmailSubscriptionTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := estcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (estcb *EmailSubscriptionTypeCreateBulk) ExecX(ctx context.Context) {
	if err := estcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.EmailSubscriptionType.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EmailSubscriptionTypeUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (estcb *EmailSubscriptionTypeCreateBulk) OnConflict(opts ...sql.ConflictOption) *EmailSubscriptionTypeUpsertBulk {
	estcb.conflict = opts
	return &EmailSubscriptionTypeUpsertBulk{
		create: estcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.EmailSubscriptionType.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (estcb *EmailSubscriptionTypeCreateBulk) OnConflictColumns(columns ...string) *EmailSubscriptionTypeUpsertBulk {
	estcb.conflict = append(estcb.conflict, sql.ConflictColumns(columns...))
	return &EmailSubscriptionTypeUpsertBulk{
		create: estcb,
	}
}

// EmailSubscriptionTypeUpsertBulk is the builder for "upsert"-ing
// a bulk of EmailSubscriptionType nodes.
type EmailSubscriptionTypeUpsertBulk struct {
	create *EmailSubscriptionTypeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.EmailSubscriptionType.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *EmailSubscriptionTypeUpsertBulk) UpdateNewValues() *EmailSubscriptionTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(emailsubscriptiontype.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.EmailSubscriptionType.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *EmailSubscriptionTypeUpsertBulk) Ignore() *EmailSubscriptionTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EmailSubscriptionTypeUpsertBulk) DoNothing() *EmailSubscriptionTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EmailSubscriptionTypeCreateBulk.OnConflict
// documentation for more info.
func (u *EmailSubscriptionTypeUpsertBulk) Update(set func(*EmailSubscriptionTypeUpsert)) *EmailSubscriptionTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EmailSubscriptionTypeUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *EmailSubscriptionTypeUpsertBulk) SetUpdatedAt(v time.Time) *EmailSubscriptionTypeUpsertBulk {
	return u.Update(func(s *EmailSubscriptionTypeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *EmailSubscriptionTypeUpsertBulk) UpdateUpdatedAt() *EmailSubscriptionTypeUpsertBulk {
	return u.Update(func(s *EmailSubscriptionTypeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *EmailSubscriptionTypeUpsertBulk) SetName(v emailsubscriptiontype.Name) *EmailSubscriptionTypeUpsertBulk {
	return u.Update(func(s *EmailSubscriptionTypeUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *EmailSubscriptionTypeUpsertBulk) UpdateName() *EmailSubscriptionTypeUpsertBulk {
	return u.Update(func(s *EmailSubscriptionTypeUpsert) {
		s.UpdateName()
	})
}

// SetActive sets the "active" field.
func (u *EmailSubscriptionTypeUpsertBulk) SetActive(v bool) *EmailSubscriptionTypeUpsertBulk {
	return u.Update(func(s *EmailSubscriptionTypeUpsert) {
		s.SetActive(v)
	})
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *EmailSubscriptionTypeUpsertBulk) UpdateActive() *EmailSubscriptionTypeUpsertBulk {
	return u.Update(func(s *EmailSubscriptionTypeUpsert) {
		s.UpdateActive()
	})
}

// Exec executes the query.
func (u *EmailSubscriptionTypeUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the EmailSubscriptionTypeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EmailSubscriptionTypeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EmailSubscriptionTypeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
