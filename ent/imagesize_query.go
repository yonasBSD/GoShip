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
	"github.com/mikestefanello/pagoda/ent/filestorage"
	"github.com/mikestefanello/pagoda/ent/image"
	"github.com/mikestefanello/pagoda/ent/imagesize"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// ImageSizeQuery is the builder for querying ImageSize entities.
type ImageSizeQuery struct {
	config
	ctx        *QueryContext
	order      []imagesize.OrderOption
	inters     []Interceptor
	predicates []predicate.ImageSize
	withFile   *FileStorageQuery
	withImage  *ImageQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ImageSizeQuery builder.
func (isq *ImageSizeQuery) Where(ps ...predicate.ImageSize) *ImageSizeQuery {
	isq.predicates = append(isq.predicates, ps...)
	return isq
}

// Limit the number of records to be returned by this query.
func (isq *ImageSizeQuery) Limit(limit int) *ImageSizeQuery {
	isq.ctx.Limit = &limit
	return isq
}

// Offset to start from.
func (isq *ImageSizeQuery) Offset(offset int) *ImageSizeQuery {
	isq.ctx.Offset = &offset
	return isq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (isq *ImageSizeQuery) Unique(unique bool) *ImageSizeQuery {
	isq.ctx.Unique = &unique
	return isq
}

// Order specifies how the records should be ordered.
func (isq *ImageSizeQuery) Order(o ...imagesize.OrderOption) *ImageSizeQuery {
	isq.order = append(isq.order, o...)
	return isq
}

// QueryFile chains the current query on the "file" edge.
func (isq *ImageSizeQuery) QueryFile() *FileStorageQuery {
	query := (&FileStorageClient{config: isq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := isq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := isq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(imagesize.Table, imagesize.FieldID, selector),
			sqlgraph.To(filestorage.Table, filestorage.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, imagesize.FileTable, imagesize.FileColumn),
		)
		fromU = sqlgraph.SetNeighbors(isq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryImage chains the current query on the "image" edge.
func (isq *ImageSizeQuery) QueryImage() *ImageQuery {
	query := (&ImageClient{config: isq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := isq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := isq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(imagesize.Table, imagesize.FieldID, selector),
			sqlgraph.To(image.Table, image.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, imagesize.ImageTable, imagesize.ImageColumn),
		)
		fromU = sqlgraph.SetNeighbors(isq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ImageSize entity from the query.
// Returns a *NotFoundError when no ImageSize was found.
func (isq *ImageSizeQuery) First(ctx context.Context) (*ImageSize, error) {
	nodes, err := isq.Limit(1).All(setContextOp(ctx, isq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{imagesize.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (isq *ImageSizeQuery) FirstX(ctx context.Context) *ImageSize {
	node, err := isq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ImageSize ID from the query.
// Returns a *NotFoundError when no ImageSize ID was found.
func (isq *ImageSizeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = isq.Limit(1).IDs(setContextOp(ctx, isq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{imagesize.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (isq *ImageSizeQuery) FirstIDX(ctx context.Context) int {
	id, err := isq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ImageSize entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ImageSize entity is found.
// Returns a *NotFoundError when no ImageSize entities are found.
func (isq *ImageSizeQuery) Only(ctx context.Context) (*ImageSize, error) {
	nodes, err := isq.Limit(2).All(setContextOp(ctx, isq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{imagesize.Label}
	default:
		return nil, &NotSingularError{imagesize.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (isq *ImageSizeQuery) OnlyX(ctx context.Context) *ImageSize {
	node, err := isq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ImageSize ID in the query.
// Returns a *NotSingularError when more than one ImageSize ID is found.
// Returns a *NotFoundError when no entities are found.
func (isq *ImageSizeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = isq.Limit(2).IDs(setContextOp(ctx, isq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{imagesize.Label}
	default:
		err = &NotSingularError{imagesize.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (isq *ImageSizeQuery) OnlyIDX(ctx context.Context) int {
	id, err := isq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ImageSizes.
func (isq *ImageSizeQuery) All(ctx context.Context) ([]*ImageSize, error) {
	ctx = setContextOp(ctx, isq.ctx, ent.OpQueryAll)
	if err := isq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ImageSize, *ImageSizeQuery]()
	return withInterceptors[[]*ImageSize](ctx, isq, qr, isq.inters)
}

// AllX is like All, but panics if an error occurs.
func (isq *ImageSizeQuery) AllX(ctx context.Context) []*ImageSize {
	nodes, err := isq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ImageSize IDs.
func (isq *ImageSizeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if isq.ctx.Unique == nil && isq.path != nil {
		isq.Unique(true)
	}
	ctx = setContextOp(ctx, isq.ctx, ent.OpQueryIDs)
	if err = isq.Select(imagesize.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (isq *ImageSizeQuery) IDsX(ctx context.Context) []int {
	ids, err := isq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (isq *ImageSizeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, isq.ctx, ent.OpQueryCount)
	if err := isq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, isq, querierCount[*ImageSizeQuery](), isq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (isq *ImageSizeQuery) CountX(ctx context.Context) int {
	count, err := isq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (isq *ImageSizeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, isq.ctx, ent.OpQueryExist)
	switch _, err := isq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (isq *ImageSizeQuery) ExistX(ctx context.Context) bool {
	exist, err := isq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ImageSizeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (isq *ImageSizeQuery) Clone() *ImageSizeQuery {
	if isq == nil {
		return nil
	}
	return &ImageSizeQuery{
		config:     isq.config,
		ctx:        isq.ctx.Clone(),
		order:      append([]imagesize.OrderOption{}, isq.order...),
		inters:     append([]Interceptor{}, isq.inters...),
		predicates: append([]predicate.ImageSize{}, isq.predicates...),
		withFile:   isq.withFile.Clone(),
		withImage:  isq.withImage.Clone(),
		// clone intermediate query.
		sql:  isq.sql.Clone(),
		path: isq.path,
	}
}

// WithFile tells the query-builder to eager-load the nodes that are connected to
// the "file" edge. The optional arguments are used to configure the query builder of the edge.
func (isq *ImageSizeQuery) WithFile(opts ...func(*FileStorageQuery)) *ImageSizeQuery {
	query := (&FileStorageClient{config: isq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	isq.withFile = query
	return isq
}

// WithImage tells the query-builder to eager-load the nodes that are connected to
// the "image" edge. The optional arguments are used to configure the query builder of the edge.
func (isq *ImageSizeQuery) WithImage(opts ...func(*ImageQuery)) *ImageSizeQuery {
	query := (&ImageClient{config: isq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	isq.withImage = query
	return isq
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
//	client.ImageSize.Query().
//		GroupBy(imagesize.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (isq *ImageSizeQuery) GroupBy(field string, fields ...string) *ImageSizeGroupBy {
	isq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ImageSizeGroupBy{build: isq}
	grbuild.flds = &isq.ctx.Fields
	grbuild.label = imagesize.Label
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
//	client.ImageSize.Query().
//		Select(imagesize.FieldCreatedAt).
//		Scan(ctx, &v)
func (isq *ImageSizeQuery) Select(fields ...string) *ImageSizeSelect {
	isq.ctx.Fields = append(isq.ctx.Fields, fields...)
	sbuild := &ImageSizeSelect{ImageSizeQuery: isq}
	sbuild.label = imagesize.Label
	sbuild.flds, sbuild.scan = &isq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ImageSizeSelect configured with the given aggregations.
func (isq *ImageSizeQuery) Aggregate(fns ...AggregateFunc) *ImageSizeSelect {
	return isq.Select().Aggregate(fns...)
}

func (isq *ImageSizeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range isq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, isq); err != nil {
				return err
			}
		}
	}
	for _, f := range isq.ctx.Fields {
		if !imagesize.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if isq.path != nil {
		prev, err := isq.path(ctx)
		if err != nil {
			return err
		}
		isq.sql = prev
	}
	return nil
}

func (isq *ImageSizeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ImageSize, error) {
	var (
		nodes       = []*ImageSize{}
		withFKs     = isq.withFKs
		_spec       = isq.querySpec()
		loadedTypes = [2]bool{
			isq.withFile != nil,
			isq.withImage != nil,
		}
	)
	if isq.withFile != nil || isq.withImage != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, imagesize.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ImageSize).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ImageSize{config: isq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, isq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := isq.withFile; query != nil {
		if err := isq.loadFile(ctx, query, nodes, nil,
			func(n *ImageSize, e *FileStorage) { n.Edges.File = e }); err != nil {
			return nil, err
		}
	}
	if query := isq.withImage; query != nil {
		if err := isq.loadImage(ctx, query, nodes, nil,
			func(n *ImageSize, e *Image) { n.Edges.Image = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (isq *ImageSizeQuery) loadFile(ctx context.Context, query *FileStorageQuery, nodes []*ImageSize, init func(*ImageSize), assign func(*ImageSize, *FileStorage)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ImageSize)
	for i := range nodes {
		if nodes[i].image_size_file == nil {
			continue
		}
		fk := *nodes[i].image_size_file
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(filestorage.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "image_size_file" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (isq *ImageSizeQuery) loadImage(ctx context.Context, query *ImageQuery, nodes []*ImageSize, init func(*ImageSize), assign func(*ImageSize, *Image)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ImageSize)
	for i := range nodes {
		if nodes[i].image_sizes == nil {
			continue
		}
		fk := *nodes[i].image_sizes
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(image.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "image_sizes" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (isq *ImageSizeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := isq.querySpec()
	_spec.Node.Columns = isq.ctx.Fields
	if len(isq.ctx.Fields) > 0 {
		_spec.Unique = isq.ctx.Unique != nil && *isq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, isq.driver, _spec)
}

func (isq *ImageSizeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(imagesize.Table, imagesize.Columns, sqlgraph.NewFieldSpec(imagesize.FieldID, field.TypeInt))
	_spec.From = isq.sql
	if unique := isq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if isq.path != nil {
		_spec.Unique = true
	}
	if fields := isq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, imagesize.FieldID)
		for i := range fields {
			if fields[i] != imagesize.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := isq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := isq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := isq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := isq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (isq *ImageSizeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(isq.driver.Dialect())
	t1 := builder.Table(imagesize.Table)
	columns := isq.ctx.Fields
	if len(columns) == 0 {
		columns = imagesize.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if isq.sql != nil {
		selector = isq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if isq.ctx.Unique != nil && *isq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range isq.predicates {
		p(selector)
	}
	for _, p := range isq.order {
		p(selector)
	}
	if offset := isq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := isq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ImageSizeGroupBy is the group-by builder for ImageSize entities.
type ImageSizeGroupBy struct {
	selector
	build *ImageSizeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (isgb *ImageSizeGroupBy) Aggregate(fns ...AggregateFunc) *ImageSizeGroupBy {
	isgb.fns = append(isgb.fns, fns...)
	return isgb
}

// Scan applies the selector query and scans the result into the given value.
func (isgb *ImageSizeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, isgb.build.ctx, ent.OpQueryGroupBy)
	if err := isgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ImageSizeQuery, *ImageSizeGroupBy](ctx, isgb.build, isgb, isgb.build.inters, v)
}

func (isgb *ImageSizeGroupBy) sqlScan(ctx context.Context, root *ImageSizeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(isgb.fns))
	for _, fn := range isgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*isgb.flds)+len(isgb.fns))
		for _, f := range *isgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*isgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := isgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ImageSizeSelect is the builder for selecting fields of ImageSize entities.
type ImageSizeSelect struct {
	*ImageSizeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (iss *ImageSizeSelect) Aggregate(fns ...AggregateFunc) *ImageSizeSelect {
	iss.fns = append(iss.fns, fns...)
	return iss
}

// Scan applies the selector query and scans the result into the given value.
func (iss *ImageSizeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, iss.ctx, ent.OpQuerySelect)
	if err := iss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ImageSizeQuery, *ImageSizeSelect](ctx, iss.ImageSizeQuery, iss, iss.inters, v)
}

func (iss *ImageSizeSelect) sqlScan(ctx context.Context, root *ImageSizeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(iss.fns))
	for _, fn := range iss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*iss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := iss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}