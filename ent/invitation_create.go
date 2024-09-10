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
	"github.com/mikestefanello/pagoda/ent/invitation"
	"github.com/mikestefanello/pagoda/ent/profile"
)

// InvitationCreate is the builder for creating a Invitation entity.
type InvitationCreate struct {
	config
	mutation *InvitationMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (ic *InvitationCreate) SetCreatedAt(t time.Time) *InvitationCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ic *InvitationCreate) SetNillableCreatedAt(t *time.Time) *InvitationCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetUpdatedAt sets the "updated_at" field.
func (ic *InvitationCreate) SetUpdatedAt(t time.Time) *InvitationCreate {
	ic.mutation.SetUpdatedAt(t)
	return ic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ic *InvitationCreate) SetNillableUpdatedAt(t *time.Time) *InvitationCreate {
	if t != nil {
		ic.SetUpdatedAt(*t)
	}
	return ic
}

// SetInviteeName sets the "invitee_name" field.
func (ic *InvitationCreate) SetInviteeName(s string) *InvitationCreate {
	ic.mutation.SetInviteeName(s)
	return ic
}

// SetConfirmationCode sets the "confirmation_code" field.
func (ic *InvitationCreate) SetConfirmationCode(s string) *InvitationCreate {
	ic.mutation.SetConfirmationCode(s)
	return ic
}

// SetInviterID sets the "inviter" edge to the Profile entity by ID.
func (ic *InvitationCreate) SetInviterID(id int) *InvitationCreate {
	ic.mutation.SetInviterID(id)
	return ic
}

// SetInviter sets the "inviter" edge to the Profile entity.
func (ic *InvitationCreate) SetInviter(p *Profile) *InvitationCreate {
	return ic.SetInviterID(p.ID)
}

// Mutation returns the InvitationMutation object of the builder.
func (ic *InvitationCreate) Mutation() *InvitationMutation {
	return ic.mutation
}

// Save creates the Invitation in the database.
func (ic *InvitationCreate) Save(ctx context.Context) (*Invitation, error) {
	ic.defaults()
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *InvitationCreate) SaveX(ctx context.Context) *Invitation {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *InvitationCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *InvitationCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *InvitationCreate) defaults() {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		v := invitation.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		v := invitation.DefaultUpdatedAt()
		ic.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *InvitationCreate) check() error {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Invitation.created_at"`)}
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Invitation.updated_at"`)}
	}
	if _, ok := ic.mutation.InviteeName(); !ok {
		return &ValidationError{Name: "invitee_name", err: errors.New(`ent: missing required field "Invitation.invitee_name"`)}
	}
	if v, ok := ic.mutation.InviteeName(); ok {
		if err := invitation.InviteeNameValidator(v); err != nil {
			return &ValidationError{Name: "invitee_name", err: fmt.Errorf(`ent: validator failed for field "Invitation.invitee_name": %w`, err)}
		}
	}
	if _, ok := ic.mutation.ConfirmationCode(); !ok {
		return &ValidationError{Name: "confirmation_code", err: errors.New(`ent: missing required field "Invitation.confirmation_code"`)}
	}
	if v, ok := ic.mutation.ConfirmationCode(); ok {
		if err := invitation.ConfirmationCodeValidator(v); err != nil {
			return &ValidationError{Name: "confirmation_code", err: fmt.Errorf(`ent: validator failed for field "Invitation.confirmation_code": %w`, err)}
		}
	}
	if len(ic.mutation.InviterIDs()) == 0 {
		return &ValidationError{Name: "inviter", err: errors.New(`ent: missing required edge "Invitation.inviter"`)}
	}
	return nil
}

func (ic *InvitationCreate) sqlSave(ctx context.Context) (*Invitation, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *InvitationCreate) createSpec() (*Invitation, *sqlgraph.CreateSpec) {
	var (
		_node = &Invitation{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(invitation.Table, sqlgraph.NewFieldSpec(invitation.FieldID, field.TypeInt))
	)
	_spec.OnConflict = ic.conflict
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.SetField(invitation.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ic.mutation.UpdatedAt(); ok {
		_spec.SetField(invitation.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ic.mutation.InviteeName(); ok {
		_spec.SetField(invitation.FieldInviteeName, field.TypeString, value)
		_node.InviteeName = value
	}
	if value, ok := ic.mutation.ConfirmationCode(); ok {
		_spec.SetField(invitation.FieldConfirmationCode, field.TypeString, value)
		_node.ConfirmationCode = value
	}
	if nodes := ic.mutation.InviterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invitation.InviterTable,
			Columns: []string{invitation.InviterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.profile_invitations = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Invitation.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.InvitationUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ic *InvitationCreate) OnConflict(opts ...sql.ConflictOption) *InvitationUpsertOne {
	ic.conflict = opts
	return &InvitationUpsertOne{
		create: ic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Invitation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ic *InvitationCreate) OnConflictColumns(columns ...string) *InvitationUpsertOne {
	ic.conflict = append(ic.conflict, sql.ConflictColumns(columns...))
	return &InvitationUpsertOne{
		create: ic,
	}
}

type (
	// InvitationUpsertOne is the builder for "upsert"-ing
	//  one Invitation node.
	InvitationUpsertOne struct {
		create *InvitationCreate
	}

	// InvitationUpsert is the "OnConflict" setter.
	InvitationUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *InvitationUpsert) SetUpdatedAt(v time.Time) *InvitationUpsert {
	u.Set(invitation.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *InvitationUpsert) UpdateUpdatedAt() *InvitationUpsert {
	u.SetExcluded(invitation.FieldUpdatedAt)
	return u
}

// SetInviteeName sets the "invitee_name" field.
func (u *InvitationUpsert) SetInviteeName(v string) *InvitationUpsert {
	u.Set(invitation.FieldInviteeName, v)
	return u
}

// UpdateInviteeName sets the "invitee_name" field to the value that was provided on create.
func (u *InvitationUpsert) UpdateInviteeName() *InvitationUpsert {
	u.SetExcluded(invitation.FieldInviteeName)
	return u
}

// SetConfirmationCode sets the "confirmation_code" field.
func (u *InvitationUpsert) SetConfirmationCode(v string) *InvitationUpsert {
	u.Set(invitation.FieldConfirmationCode, v)
	return u
}

// UpdateConfirmationCode sets the "confirmation_code" field to the value that was provided on create.
func (u *InvitationUpsert) UpdateConfirmationCode() *InvitationUpsert {
	u.SetExcluded(invitation.FieldConfirmationCode)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Invitation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *InvitationUpsertOne) UpdateNewValues() *InvitationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(invitation.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Invitation.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *InvitationUpsertOne) Ignore() *InvitationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *InvitationUpsertOne) DoNothing() *InvitationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the InvitationCreate.OnConflict
// documentation for more info.
func (u *InvitationUpsertOne) Update(set func(*InvitationUpsert)) *InvitationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&InvitationUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *InvitationUpsertOne) SetUpdatedAt(v time.Time) *InvitationUpsertOne {
	return u.Update(func(s *InvitationUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *InvitationUpsertOne) UpdateUpdatedAt() *InvitationUpsertOne {
	return u.Update(func(s *InvitationUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetInviteeName sets the "invitee_name" field.
func (u *InvitationUpsertOne) SetInviteeName(v string) *InvitationUpsertOne {
	return u.Update(func(s *InvitationUpsert) {
		s.SetInviteeName(v)
	})
}

// UpdateInviteeName sets the "invitee_name" field to the value that was provided on create.
func (u *InvitationUpsertOne) UpdateInviteeName() *InvitationUpsertOne {
	return u.Update(func(s *InvitationUpsert) {
		s.UpdateInviteeName()
	})
}

// SetConfirmationCode sets the "confirmation_code" field.
func (u *InvitationUpsertOne) SetConfirmationCode(v string) *InvitationUpsertOne {
	return u.Update(func(s *InvitationUpsert) {
		s.SetConfirmationCode(v)
	})
}

// UpdateConfirmationCode sets the "confirmation_code" field to the value that was provided on create.
func (u *InvitationUpsertOne) UpdateConfirmationCode() *InvitationUpsertOne {
	return u.Update(func(s *InvitationUpsert) {
		s.UpdateConfirmationCode()
	})
}

// Exec executes the query.
func (u *InvitationUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for InvitationCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *InvitationUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *InvitationUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *InvitationUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// InvitationCreateBulk is the builder for creating many Invitation entities in bulk.
type InvitationCreateBulk struct {
	config
	err      error
	builders []*InvitationCreate
	conflict []sql.ConflictOption
}

// Save creates the Invitation entities in the database.
func (icb *InvitationCreateBulk) Save(ctx context.Context) ([]*Invitation, error) {
	if icb.err != nil {
		return nil, icb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Invitation, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InvitationMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = icb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *InvitationCreateBulk) SaveX(ctx context.Context) []*Invitation {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *InvitationCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *InvitationCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Invitation.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.InvitationUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (icb *InvitationCreateBulk) OnConflict(opts ...sql.ConflictOption) *InvitationUpsertBulk {
	icb.conflict = opts
	return &InvitationUpsertBulk{
		create: icb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Invitation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (icb *InvitationCreateBulk) OnConflictColumns(columns ...string) *InvitationUpsertBulk {
	icb.conflict = append(icb.conflict, sql.ConflictColumns(columns...))
	return &InvitationUpsertBulk{
		create: icb,
	}
}

// InvitationUpsertBulk is the builder for "upsert"-ing
// a bulk of Invitation nodes.
type InvitationUpsertBulk struct {
	create *InvitationCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Invitation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *InvitationUpsertBulk) UpdateNewValues() *InvitationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(invitation.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Invitation.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *InvitationUpsertBulk) Ignore() *InvitationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *InvitationUpsertBulk) DoNothing() *InvitationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the InvitationCreateBulk.OnConflict
// documentation for more info.
func (u *InvitationUpsertBulk) Update(set func(*InvitationUpsert)) *InvitationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&InvitationUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *InvitationUpsertBulk) SetUpdatedAt(v time.Time) *InvitationUpsertBulk {
	return u.Update(func(s *InvitationUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *InvitationUpsertBulk) UpdateUpdatedAt() *InvitationUpsertBulk {
	return u.Update(func(s *InvitationUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetInviteeName sets the "invitee_name" field.
func (u *InvitationUpsertBulk) SetInviteeName(v string) *InvitationUpsertBulk {
	return u.Update(func(s *InvitationUpsert) {
		s.SetInviteeName(v)
	})
}

// UpdateInviteeName sets the "invitee_name" field to the value that was provided on create.
func (u *InvitationUpsertBulk) UpdateInviteeName() *InvitationUpsertBulk {
	return u.Update(func(s *InvitationUpsert) {
		s.UpdateInviteeName()
	})
}

// SetConfirmationCode sets the "confirmation_code" field.
func (u *InvitationUpsertBulk) SetConfirmationCode(v string) *InvitationUpsertBulk {
	return u.Update(func(s *InvitationUpsert) {
		s.SetConfirmationCode(v)
	})
}

// UpdateConfirmationCode sets the "confirmation_code" field to the value that was provided on create.
func (u *InvitationUpsertBulk) UpdateConfirmationCode() *InvitationUpsertBulk {
	return u.Update(func(s *InvitationUpsert) {
		s.UpdateConfirmationCode()
	})
}

// Exec executes the query.
func (u *InvitationUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the InvitationCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for InvitationCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *InvitationUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}