// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// THSeasonCach is an object representing the database table.
type THSeasonCach struct {
	ID        int        `boil:"id" json:"id" toml:"id" yaml:"id"`
	SeasonID  int64      `boil:"season_id" json:"season_id" toml:"season_id" yaml:"season_id"`
	IsVip     bool       `boil:"is_vip" json:"is_vip" toml:"is_vip" yaml:"is_vip"`
	Data      types.JSON `boil:"data" json:"data" toml:"data" yaml:"data"`
	CreatedAt time.Time  `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time  `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *thSeasonCachR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L thSeasonCachL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var THSeasonCachColumns = struct {
	ID        string
	SeasonID  string
	IsVip     string
	Data      string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	SeasonID:  "season_id",
	IsVip:     "is_vip",
	Data:      "data",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var THSeasonCachTableColumns = struct {
	ID        string
	SeasonID  string
	IsVip     string
	Data      string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "th_season_caches.id",
	SeasonID:  "th_season_caches.season_id",
	IsVip:     "th_season_caches.is_vip",
	Data:      "th_season_caches.data",
	CreatedAt: "th_season_caches.created_at",
	UpdatedAt: "th_season_caches.updated_at",
}

// Generated where

var THSeasonCachWhere = struct {
	ID        whereHelperint
	SeasonID  whereHelperint64
	IsVip     whereHelperbool
	Data      whereHelpertypes_JSON
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint{field: "\"th_season_caches\".\"id\""},
	SeasonID:  whereHelperint64{field: "\"th_season_caches\".\"season_id\""},
	IsVip:     whereHelperbool{field: "\"th_season_caches\".\"is_vip\""},
	Data:      whereHelpertypes_JSON{field: "\"th_season_caches\".\"data\""},
	CreatedAt: whereHelpertime_Time{field: "\"th_season_caches\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"th_season_caches\".\"updated_at\""},
}

// THSeasonCachRels is where relationship names are stored.
var THSeasonCachRels = struct {
}{}

// thSeasonCachR is where relationships are stored.
type thSeasonCachR struct {
}

// NewStruct creates a new relationship struct
func (*thSeasonCachR) NewStruct() *thSeasonCachR {
	return &thSeasonCachR{}
}

// thSeasonCachL is where Load methods for each relationship are stored.
type thSeasonCachL struct{}

var (
	thSeasonCachAllColumns            = []string{"id", "season_id", "is_vip", "data", "created_at", "updated_at"}
	thSeasonCachColumnsWithoutDefault = []string{"season_id", "is_vip", "data", "created_at", "updated_at"}
	thSeasonCachColumnsWithDefault    = []string{"id"}
	thSeasonCachPrimaryKeyColumns     = []string{"id"}
	thSeasonCachGeneratedColumns      = []string{}
)

type (
	// THSeasonCachSlice is an alias for a slice of pointers to THSeasonCach.
	// This should almost always be used instead of []THSeasonCach.
	THSeasonCachSlice []*THSeasonCach
	// THSeasonCachHook is the signature for custom THSeasonCach hook methods
	THSeasonCachHook func(context.Context, boil.ContextExecutor, *THSeasonCach) error

	thSeasonCachQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	thSeasonCachType                 = reflect.TypeOf(&THSeasonCach{})
	thSeasonCachMapping              = queries.MakeStructMapping(thSeasonCachType)
	thSeasonCachPrimaryKeyMapping, _ = queries.BindMapping(thSeasonCachType, thSeasonCachMapping, thSeasonCachPrimaryKeyColumns)
	thSeasonCachInsertCacheMut       sync.RWMutex
	thSeasonCachInsertCache          = make(map[string]insertCache)
	thSeasonCachUpdateCacheMut       sync.RWMutex
	thSeasonCachUpdateCache          = make(map[string]updateCache)
	thSeasonCachUpsertCacheMut       sync.RWMutex
	thSeasonCachUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var thSeasonCachAfterSelectHooks []THSeasonCachHook

var thSeasonCachBeforeInsertHooks []THSeasonCachHook
var thSeasonCachAfterInsertHooks []THSeasonCachHook

var thSeasonCachBeforeUpdateHooks []THSeasonCachHook
var thSeasonCachAfterUpdateHooks []THSeasonCachHook

var thSeasonCachBeforeDeleteHooks []THSeasonCachHook
var thSeasonCachAfterDeleteHooks []THSeasonCachHook

var thSeasonCachBeforeUpsertHooks []THSeasonCachHook
var thSeasonCachAfterUpsertHooks []THSeasonCachHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *THSeasonCach) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeasonCachAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *THSeasonCach) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeasonCachBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *THSeasonCach) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeasonCachAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *THSeasonCach) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeasonCachBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *THSeasonCach) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeasonCachAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *THSeasonCach) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeasonCachBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *THSeasonCach) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeasonCachAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *THSeasonCach) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeasonCachBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *THSeasonCach) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeasonCachAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTHSeasonCachHook registers your hook function for all future operations.
func AddTHSeasonCachHook(hookPoint boil.HookPoint, thSeasonCachHook THSeasonCachHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		thSeasonCachAfterSelectHooks = append(thSeasonCachAfterSelectHooks, thSeasonCachHook)
	case boil.BeforeInsertHook:
		thSeasonCachBeforeInsertHooks = append(thSeasonCachBeforeInsertHooks, thSeasonCachHook)
	case boil.AfterInsertHook:
		thSeasonCachAfterInsertHooks = append(thSeasonCachAfterInsertHooks, thSeasonCachHook)
	case boil.BeforeUpdateHook:
		thSeasonCachBeforeUpdateHooks = append(thSeasonCachBeforeUpdateHooks, thSeasonCachHook)
	case boil.AfterUpdateHook:
		thSeasonCachAfterUpdateHooks = append(thSeasonCachAfterUpdateHooks, thSeasonCachHook)
	case boil.BeforeDeleteHook:
		thSeasonCachBeforeDeleteHooks = append(thSeasonCachBeforeDeleteHooks, thSeasonCachHook)
	case boil.AfterDeleteHook:
		thSeasonCachAfterDeleteHooks = append(thSeasonCachAfterDeleteHooks, thSeasonCachHook)
	case boil.BeforeUpsertHook:
		thSeasonCachBeforeUpsertHooks = append(thSeasonCachBeforeUpsertHooks, thSeasonCachHook)
	case boil.AfterUpsertHook:
		thSeasonCachAfterUpsertHooks = append(thSeasonCachAfterUpsertHooks, thSeasonCachHook)
	}
}

// One returns a single thSeasonCach record from the query.
func (q thSeasonCachQuery) One(ctx context.Context, exec boil.ContextExecutor) (*THSeasonCach, error) {
	o := &THSeasonCach{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for th_season_caches")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all THSeasonCach records from the query.
func (q thSeasonCachQuery) All(ctx context.Context, exec boil.ContextExecutor) (THSeasonCachSlice, error) {
	var o []*THSeasonCach

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to THSeasonCach slice")
	}

	if len(thSeasonCachAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all THSeasonCach records in the query.
func (q thSeasonCachQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count th_season_caches rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q thSeasonCachQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if th_season_caches exists")
	}

	return count > 0, nil
}

// THSeasonCaches retrieves all the records using an executor.
func THSeasonCaches(mods ...qm.QueryMod) thSeasonCachQuery {
	mods = append(mods, qm.From("\"th_season_caches\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"th_season_caches\".*"})
	}

	return thSeasonCachQuery{q}
}

// FindTHSeasonCach retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTHSeasonCach(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*THSeasonCach, error) {
	thSeasonCachObj := &THSeasonCach{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"th_season_caches\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, thSeasonCachObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from th_season_caches")
	}

	if err = thSeasonCachObj.doAfterSelectHooks(ctx, exec); err != nil {
		return thSeasonCachObj, err
	}

	return thSeasonCachObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *THSeasonCach) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no th_season_caches provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(thSeasonCachColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	thSeasonCachInsertCacheMut.RLock()
	cache, cached := thSeasonCachInsertCache[key]
	thSeasonCachInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			thSeasonCachAllColumns,
			thSeasonCachColumnsWithDefault,
			thSeasonCachColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(thSeasonCachType, thSeasonCachMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(thSeasonCachType, thSeasonCachMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"th_season_caches\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"th_season_caches\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into th_season_caches")
	}

	if !cached {
		thSeasonCachInsertCacheMut.Lock()
		thSeasonCachInsertCache[key] = cache
		thSeasonCachInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the THSeasonCach.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *THSeasonCach) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	thSeasonCachUpdateCacheMut.RLock()
	cache, cached := thSeasonCachUpdateCache[key]
	thSeasonCachUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			thSeasonCachAllColumns,
			thSeasonCachPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update th_season_caches, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"th_season_caches\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, thSeasonCachPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(thSeasonCachType, thSeasonCachMapping, append(wl, thSeasonCachPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update th_season_caches row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for th_season_caches")
	}

	if !cached {
		thSeasonCachUpdateCacheMut.Lock()
		thSeasonCachUpdateCache[key] = cache
		thSeasonCachUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q thSeasonCachQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for th_season_caches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for th_season_caches")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o THSeasonCachSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), thSeasonCachPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"th_season_caches\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, thSeasonCachPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in thSeasonCach slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all thSeasonCach")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *THSeasonCach) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no th_season_caches provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(thSeasonCachColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	thSeasonCachUpsertCacheMut.RLock()
	cache, cached := thSeasonCachUpsertCache[key]
	thSeasonCachUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			thSeasonCachAllColumns,
			thSeasonCachColumnsWithDefault,
			thSeasonCachColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			thSeasonCachAllColumns,
			thSeasonCachPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert th_season_caches, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(thSeasonCachPrimaryKeyColumns))
			copy(conflict, thSeasonCachPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"th_season_caches\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(thSeasonCachType, thSeasonCachMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(thSeasonCachType, thSeasonCachMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert th_season_caches")
	}

	if !cached {
		thSeasonCachUpsertCacheMut.Lock()
		thSeasonCachUpsertCache[key] = cache
		thSeasonCachUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single THSeasonCach record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *THSeasonCach) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no THSeasonCach provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), thSeasonCachPrimaryKeyMapping)
	sql := "DELETE FROM \"th_season_caches\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from th_season_caches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for th_season_caches")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q thSeasonCachQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no thSeasonCachQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from th_season_caches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for th_season_caches")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o THSeasonCachSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(thSeasonCachBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), thSeasonCachPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"th_season_caches\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, thSeasonCachPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from thSeasonCach slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for th_season_caches")
	}

	if len(thSeasonCachAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *THSeasonCach) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTHSeasonCach(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *THSeasonCachSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := THSeasonCachSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), thSeasonCachPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"th_season_caches\".* FROM \"th_season_caches\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, thSeasonCachPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in THSeasonCachSlice")
	}

	*o = slice

	return nil
}

// THSeasonCachExists checks if the THSeasonCach row exists.
func THSeasonCachExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"th_season_caches\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if th_season_caches exists")
	}

	return exists, nil
}
