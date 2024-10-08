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
	"github.com/mikestefanello/pagoda/ent/lastseenonline"
	"github.com/mikestefanello/pagoda/ent/user"
)

// LastSeenOnlineCreate is the builder for creating a LastSeenOnline entity.
type LastSeenOnlineCreate struct {
	config
	mutation *LastSeenOnlineMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetSeenAt sets the "seen_at" field.
func (lsoc *LastSeenOnlineCreate) SetSeenAt(t time.Time) *LastSeenOnlineCreate {
	lsoc.mutation.SetSeenAt(t)
	return lsoc
}

// SetNillableSeenAt sets the "seen_at" field if the given value is not nil.
func (lsoc *LastSeenOnlineCreate) SetNillableSeenAt(t *time.Time) *LastSeenOnlineCreate {
	if t != nil {
		lsoc.SetSeenAt(*t)
	}
	return lsoc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (lsoc *LastSeenOnlineCreate) SetUserID(id int) *LastSeenOnlineCreate {
	lsoc.mutation.SetUserID(id)
	return lsoc
}

// SetUser sets the "user" edge to the User entity.
func (lsoc *LastSeenOnlineCreate) SetUser(u *User) *LastSeenOnlineCreate {
	return lsoc.SetUserID(u.ID)
}

// Mutation returns the LastSeenOnlineMutation object of the builder.
func (lsoc *LastSeenOnlineCreate) Mutation() *LastSeenOnlineMutation {
	return lsoc.mutation
}

// Save creates the LastSeenOnline in the database.
func (lsoc *LastSeenOnlineCreate) Save(ctx context.Context) (*LastSeenOnline, error) {
	lsoc.defaults()
	return withHooks(ctx, lsoc.sqlSave, lsoc.mutation, lsoc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lsoc *LastSeenOnlineCreate) SaveX(ctx context.Context) *LastSeenOnline {
	v, err := lsoc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lsoc *LastSeenOnlineCreate) Exec(ctx context.Context) error {
	_, err := lsoc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lsoc *LastSeenOnlineCreate) ExecX(ctx context.Context) {
	if err := lsoc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lsoc *LastSeenOnlineCreate) defaults() {
	if _, ok := lsoc.mutation.SeenAt(); !ok {
		v := lastseenonline.DefaultSeenAt()
		lsoc.mutation.SetSeenAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lsoc *LastSeenOnlineCreate) check() error {
	if _, ok := lsoc.mutation.SeenAt(); !ok {
		return &ValidationError{Name: "seen_at", err: errors.New(`ent: missing required field "LastSeenOnline.seen_at"`)}
	}
	if len(lsoc.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "LastSeenOnline.user"`)}
	}
	return nil
}

func (lsoc *LastSeenOnlineCreate) sqlSave(ctx context.Context) (*LastSeenOnline, error) {
	if err := lsoc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lsoc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lsoc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	lsoc.mutation.id = &_node.ID
	lsoc.mutation.done = true
	return _node, nil
}

func (lsoc *LastSeenOnlineCreate) createSpec() (*LastSeenOnline, *sqlgraph.CreateSpec) {
	var (
		_node = &LastSeenOnline{config: lsoc.config}
		_spec = sqlgraph.NewCreateSpec(lastseenonline.Table, sqlgraph.NewFieldSpec(lastseenonline.FieldID, field.TypeInt))
	)
	_spec.OnConflict = lsoc.conflict
	if value, ok := lsoc.mutation.SeenAt(); ok {
		_spec.SetField(lastseenonline.FieldSeenAt, field.TypeTime, value)
		_node.SeenAt = value
	}
	if nodes := lsoc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lastseenonline.UserTable,
			Columns: []string{lastseenonline.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_last_seen_at = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.LastSeenOnline.Create().
//		SetSeenAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LastSeenOnlineUpsert) {
//			SetSeenAt(v+v).
//		}).
//		Exec(ctx)
func (lsoc *LastSeenOnlineCreate) OnConflict(opts ...sql.ConflictOption) *LastSeenOnlineUpsertOne {
	lsoc.conflict = opts
	return &LastSeenOnlineUpsertOne{
		create: lsoc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.LastSeenOnline.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (lsoc *LastSeenOnlineCreate) OnConflictColumns(columns ...string) *LastSeenOnlineUpsertOne {
	lsoc.conflict = append(lsoc.conflict, sql.ConflictColumns(columns...))
	return &LastSeenOnlineUpsertOne{
		create: lsoc,
	}
}

type (
	// LastSeenOnlineUpsertOne is the builder for "upsert"-ing
	//  one LastSeenOnline node.
	LastSeenOnlineUpsertOne struct {
		create *LastSeenOnlineCreate
	}

	// LastSeenOnlineUpsert is the "OnConflict" setter.
	LastSeenOnlineUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.LastSeenOnline.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *LastSeenOnlineUpsertOne) UpdateNewValues() *LastSeenOnlineUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.SeenAt(); exists {
			s.SetIgnore(lastseenonline.FieldSeenAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.LastSeenOnline.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *LastSeenOnlineUpsertOne) Ignore() *LastSeenOnlineUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LastSeenOnlineUpsertOne) DoNothing() *LastSeenOnlineUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LastSeenOnlineCreate.OnConflict
// documentation for more info.
func (u *LastSeenOnlineUpsertOne) Update(set func(*LastSeenOnlineUpsert)) *LastSeenOnlineUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LastSeenOnlineUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *LastSeenOnlineUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LastSeenOnlineCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LastSeenOnlineUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *LastSeenOnlineUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *LastSeenOnlineUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// LastSeenOnlineCreateBulk is the builder for creating many LastSeenOnline entities in bulk.
type LastSeenOnlineCreateBulk struct {
	config
	err      error
	builders []*LastSeenOnlineCreate
	conflict []sql.ConflictOption
}

// Save creates the LastSeenOnline entities in the database.
func (lsocb *LastSeenOnlineCreateBulk) Save(ctx context.Context) ([]*LastSeenOnline, error) {
	if lsocb.err != nil {
		return nil, lsocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(lsocb.builders))
	nodes := make([]*LastSeenOnline, len(lsocb.builders))
	mutators := make([]Mutator, len(lsocb.builders))
	for i := range lsocb.builders {
		func(i int, root context.Context) {
			builder := lsocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LastSeenOnlineMutation)
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
					_, err = mutators[i+1].Mutate(root, lsocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = lsocb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lsocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lsocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lsocb *LastSeenOnlineCreateBulk) SaveX(ctx context.Context) []*LastSeenOnline {
	v, err := lsocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lsocb *LastSeenOnlineCreateBulk) Exec(ctx context.Context) error {
	_, err := lsocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lsocb *LastSeenOnlineCreateBulk) ExecX(ctx context.Context) {
	if err := lsocb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.LastSeenOnline.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LastSeenOnlineUpsert) {
//			SetSeenAt(v+v).
//		}).
//		Exec(ctx)
func (lsocb *LastSeenOnlineCreateBulk) OnConflict(opts ...sql.ConflictOption) *LastSeenOnlineUpsertBulk {
	lsocb.conflict = opts
	return &LastSeenOnlineUpsertBulk{
		create: lsocb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.LastSeenOnline.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (lsocb *LastSeenOnlineCreateBulk) OnConflictColumns(columns ...string) *LastSeenOnlineUpsertBulk {
	lsocb.conflict = append(lsocb.conflict, sql.ConflictColumns(columns...))
	return &LastSeenOnlineUpsertBulk{
		create: lsocb,
	}
}

// LastSeenOnlineUpsertBulk is the builder for "upsert"-ing
// a bulk of LastSeenOnline nodes.
type LastSeenOnlineUpsertBulk struct {
	create *LastSeenOnlineCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.LastSeenOnline.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *LastSeenOnlineUpsertBulk) UpdateNewValues() *LastSeenOnlineUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.SeenAt(); exists {
				s.SetIgnore(lastseenonline.FieldSeenAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.LastSeenOnline.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *LastSeenOnlineUpsertBulk) Ignore() *LastSeenOnlineUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LastSeenOnlineUpsertBulk) DoNothing() *LastSeenOnlineUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LastSeenOnlineCreateBulk.OnConflict
// documentation for more info.
func (u *LastSeenOnlineUpsertBulk) Update(set func(*LastSeenOnlineUpsert)) *LastSeenOnlineUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LastSeenOnlineUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *LastSeenOnlineUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the LastSeenOnlineCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LastSeenOnlineCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LastSeenOnlineUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
