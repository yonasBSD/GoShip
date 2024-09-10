// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/phoneverificationcode"
	"github.com/mikestefanello/pagoda/ent/predicate"
	"github.com/mikestefanello/pagoda/ent/profile"
)

// PhoneVerificationCodeQuery is the builder for querying PhoneVerificationCode entities.
type PhoneVerificationCodeQuery struct {
	config
	ctx         *QueryContext
	order       []phoneverificationcode.OrderOption
	inters      []Interceptor
	predicates  []predicate.PhoneVerificationCode
	withProfile *ProfileQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PhoneVerificationCodeQuery builder.
func (pvcq *PhoneVerificationCodeQuery) Where(ps ...predicate.PhoneVerificationCode) *PhoneVerificationCodeQuery {
	pvcq.predicates = append(pvcq.predicates, ps...)
	return pvcq
}

// Limit the number of records to be returned by this query.
func (pvcq *PhoneVerificationCodeQuery) Limit(limit int) *PhoneVerificationCodeQuery {
	pvcq.ctx.Limit = &limit
	return pvcq
}

// Offset to start from.
func (pvcq *PhoneVerificationCodeQuery) Offset(offset int) *PhoneVerificationCodeQuery {
	pvcq.ctx.Offset = &offset
	return pvcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pvcq *PhoneVerificationCodeQuery) Unique(unique bool) *PhoneVerificationCodeQuery {
	pvcq.ctx.Unique = &unique
	return pvcq
}

// Order specifies how the records should be ordered.
func (pvcq *PhoneVerificationCodeQuery) Order(o ...phoneverificationcode.OrderOption) *PhoneVerificationCodeQuery {
	pvcq.order = append(pvcq.order, o...)
	return pvcq
}

// QueryProfile chains the current query on the "profile" edge.
func (pvcq *PhoneVerificationCodeQuery) QueryProfile() *ProfileQuery {
	query := (&ProfileClient{config: pvcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pvcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pvcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(phoneverificationcode.Table, phoneverificationcode.FieldID, selector),
			sqlgraph.To(profile.Table, profile.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, phoneverificationcode.ProfileTable, phoneverificationcode.ProfileColumn),
		)
		fromU = sqlgraph.SetNeighbors(pvcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PhoneVerificationCode entity from the query.
// Returns a *NotFoundError when no PhoneVerificationCode was found.
func (pvcq *PhoneVerificationCodeQuery) First(ctx context.Context) (*PhoneVerificationCode, error) {
	nodes, err := pvcq.Limit(1).All(setContextOp(ctx, pvcq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{phoneverificationcode.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pvcq *PhoneVerificationCodeQuery) FirstX(ctx context.Context) *PhoneVerificationCode {
	node, err := pvcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PhoneVerificationCode ID from the query.
// Returns a *NotFoundError when no PhoneVerificationCode ID was found.
func (pvcq *PhoneVerificationCodeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pvcq.Limit(1).IDs(setContextOp(ctx, pvcq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{phoneverificationcode.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pvcq *PhoneVerificationCodeQuery) FirstIDX(ctx context.Context) int {
	id, err := pvcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PhoneVerificationCode entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PhoneVerificationCode entity is found.
// Returns a *NotFoundError when no PhoneVerificationCode entities are found.
func (pvcq *PhoneVerificationCodeQuery) Only(ctx context.Context) (*PhoneVerificationCode, error) {
	nodes, err := pvcq.Limit(2).All(setContextOp(ctx, pvcq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{phoneverificationcode.Label}
	default:
		return nil, &NotSingularError{phoneverificationcode.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pvcq *PhoneVerificationCodeQuery) OnlyX(ctx context.Context) *PhoneVerificationCode {
	node, err := pvcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PhoneVerificationCode ID in the query.
// Returns a *NotSingularError when more than one PhoneVerificationCode ID is found.
// Returns a *NotFoundError when no entities are found.
func (pvcq *PhoneVerificationCodeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pvcq.Limit(2).IDs(setContextOp(ctx, pvcq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{phoneverificationcode.Label}
	default:
		err = &NotSingularError{phoneverificationcode.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pvcq *PhoneVerificationCodeQuery) OnlyIDX(ctx context.Context) int {
	id, err := pvcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PhoneVerificationCodes.
func (pvcq *PhoneVerificationCodeQuery) All(ctx context.Context) ([]*PhoneVerificationCode, error) {
	ctx = setContextOp(ctx, pvcq.ctx, ent.OpQueryAll)
	if err := pvcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PhoneVerificationCode, *PhoneVerificationCodeQuery]()
	return withInterceptors[[]*PhoneVerificationCode](ctx, pvcq, qr, pvcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pvcq *PhoneVerificationCodeQuery) AllX(ctx context.Context) []*PhoneVerificationCode {
	nodes, err := pvcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PhoneVerificationCode IDs.
func (pvcq *PhoneVerificationCodeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pvcq.ctx.Unique == nil && pvcq.path != nil {
		pvcq.Unique(true)
	}
	ctx = setContextOp(ctx, pvcq.ctx, ent.OpQueryIDs)
	if err = pvcq.Select(phoneverificationcode.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pvcq *PhoneVerificationCodeQuery) IDsX(ctx context.Context) []int {
	ids, err := pvcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pvcq *PhoneVerificationCodeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pvcq.ctx, ent.OpQueryCount)
	if err := pvcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pvcq, querierCount[*PhoneVerificationCodeQuery](), pvcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pvcq *PhoneVerificationCodeQuery) CountX(ctx context.Context) int {
	count, err := pvcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pvcq *PhoneVerificationCodeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pvcq.ctx, ent.OpQueryExist)
	switch _, err := pvcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pvcq *PhoneVerificationCodeQuery) ExistX(ctx context.Context) bool {
	exist, err := pvcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PhoneVerificationCodeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pvcq *PhoneVerificationCodeQuery) Clone() *PhoneVerificationCodeQuery {
	if pvcq == nil {
		return nil
	}
	return &PhoneVerificationCodeQuery{
		config:      pvcq.config,
		ctx:         pvcq.ctx.Clone(),
		order:       append([]phoneverificationcode.OrderOption{}, pvcq.order...),
		inters:      append([]Interceptor{}, pvcq.inters...),
		predicates:  append([]predicate.PhoneVerificationCode{}, pvcq.predicates...),
		withProfile: pvcq.withProfile.Clone(),
		// clone intermediate query.
		sql:  pvcq.sql.Clone(),
		path: pvcq.path,
	}
}

// WithProfile tells the query-builder to eager-load the nodes that are connected to
// the "profile" edge. The optional arguments are used to configure the query builder of the edge.
func (pvcq *PhoneVerificationCodeQuery) WithProfile(opts ...func(*ProfileQuery)) *PhoneVerificationCodeQuery {
	query := (&ProfileClient{config: pvcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pvcq.withProfile = query
	return pvcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PhoneVerificationCode.Query().
//		GroupBy(phoneverificationcode.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pvcq *PhoneVerificationCodeQuery) GroupBy(field string, fields ...string) *PhoneVerificationCodeGroupBy {
	pvcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PhoneVerificationCodeGroupBy{build: pvcq}
	grbuild.flds = &pvcq.ctx.Fields
	grbuild.label = phoneverificationcode.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.PhoneVerificationCode.Query().
//		Select(phoneverificationcode.FieldCreatedAt).
//		Scan(ctx, &v)
func (pvcq *PhoneVerificationCodeQuery) Select(fields ...string) *PhoneVerificationCodeSelect {
	pvcq.ctx.Fields = append(pvcq.ctx.Fields, fields...)
	sbuild := &PhoneVerificationCodeSelect{PhoneVerificationCodeQuery: pvcq}
	sbuild.label = phoneverificationcode.Label
	sbuild.flds, sbuild.scan = &pvcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PhoneVerificationCodeSelect configured with the given aggregations.
func (pvcq *PhoneVerificationCodeQuery) Aggregate(fns ...AggregateFunc) *PhoneVerificationCodeSelect {
	return pvcq.Select().Aggregate(fns...)
}

func (pvcq *PhoneVerificationCodeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pvcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pvcq); err != nil {
				return err
			}
		}
	}
	for _, f := range pvcq.ctx.Fields {
		if !phoneverificationcode.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pvcq.path != nil {
		prev, err := pvcq.path(ctx)
		if err != nil {
			return err
		}
		pvcq.sql = prev
	}
	return nil
}

func (pvcq *PhoneVerificationCodeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PhoneVerificationCode, error) {
	var (
		nodes       = []*PhoneVerificationCode{}
		_spec       = pvcq.querySpec()
		loadedTypes = [1]bool{
			pvcq.withProfile != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PhoneVerificationCode).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PhoneVerificationCode{config: pvcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pvcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pvcq.withProfile; query != nil {
		if err := pvcq.loadProfile(ctx, query, nodes, nil,
			func(n *PhoneVerificationCode, e *Profile) { n.Edges.Profile = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pvcq *PhoneVerificationCodeQuery) loadProfile(ctx context.Context, query *ProfileQuery, nodes []*PhoneVerificationCode, init func(*PhoneVerificationCode), assign func(*PhoneVerificationCode, *Profile)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PhoneVerificationCode)
	for i := range nodes {
		fk := nodes[i].ProfileID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(profile.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "profile_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pvcq *PhoneVerificationCodeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pvcq.querySpec()
	_spec.Node.Columns = pvcq.ctx.Fields
	if len(pvcq.ctx.Fields) > 0 {
		_spec.Unique = pvcq.ctx.Unique != nil && *pvcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pvcq.driver, _spec)
}

func (pvcq *PhoneVerificationCodeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(phoneverificationcode.Table, phoneverificationcode.Columns, sqlgraph.NewFieldSpec(phoneverificationcode.FieldID, field.TypeInt))
	_spec.From = pvcq.sql
	if unique := pvcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pvcq.path != nil {
		_spec.Unique = true
	}
	if fields := pvcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, phoneverificationcode.FieldID)
		for i := range fields {
			if fields[i] != phoneverificationcode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if pvcq.withProfile != nil {
			_spec.Node.AddColumnOnce(phoneverificationcode.FieldProfileID)
		}
	}
	if ps := pvcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pvcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pvcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pvcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pvcq *PhoneVerificationCodeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pvcq.driver.Dialect())
	t1 := builder.Table(phoneverificationcode.Table)
	columns := pvcq.ctx.Fields
	if len(columns) == 0 {
		columns = phoneverificationcode.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pvcq.sql != nil {
		selector = pvcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pvcq.ctx.Unique != nil && *pvcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pvcq.predicates {
		p(selector)
	}
	for _, p := range pvcq.order {
		p(selector)
	}
	if offset := pvcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pvcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PhoneVerificationCodeGroupBy is the group-by builder for PhoneVerificationCode entities.
type PhoneVerificationCodeGroupBy struct {
	selector
	build *PhoneVerificationCodeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pvcgb *PhoneVerificationCodeGroupBy) Aggregate(fns ...AggregateFunc) *PhoneVerificationCodeGroupBy {
	pvcgb.fns = append(pvcgb.fns, fns...)
	return pvcgb
}

// Scan applies the selector query and scans the result into the given value.
func (pvcgb *PhoneVerificationCodeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pvcgb.build.ctx, ent.OpQueryGroupBy)
	if err := pvcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PhoneVerificationCodeQuery, *PhoneVerificationCodeGroupBy](ctx, pvcgb.build, pvcgb, pvcgb.build.inters, v)
}

func (pvcgb *PhoneVerificationCodeGroupBy) sqlScan(ctx context.Context, root *PhoneVerificationCodeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pvcgb.fns))
	for _, fn := range pvcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pvcgb.flds)+len(pvcgb.fns))
		for _, f := range *pvcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pvcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pvcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PhoneVerificationCodeSelect is the builder for selecting fields of PhoneVerificationCode entities.
type PhoneVerificationCodeSelect struct {
	*PhoneVerificationCodeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pvcs *PhoneVerificationCodeSelect) Aggregate(fns ...AggregateFunc) *PhoneVerificationCodeSelect {
	pvcs.fns = append(pvcs.fns, fns...)
	return pvcs
}

// Scan applies the selector query and scans the result into the given value.
func (pvcs *PhoneVerificationCodeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pvcs.ctx, ent.OpQuerySelect)
	if err := pvcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PhoneVerificationCodeQuery, *PhoneVerificationCodeSelect](ctx, pvcs.PhoneVerificationCodeQuery, pvcs, pvcs.inters, v)
}

func (pvcs *PhoneVerificationCodeSelect) sqlScan(ctx context.Context, root *PhoneVerificationCodeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pvcs.fns))
	for _, fn := range pvcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pvcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pvcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}