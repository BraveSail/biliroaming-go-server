// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/strmangle"
)

// THSeason2EpisodeCach is an object representing the database table.
type THSeason2EpisodeCach struct {
	EpisodeID int64     `boil:"episode_id" json:"episode_id" toml:"episode_id" yaml:"episode_id"`
	SeasonID  int64     `boil:"season_id" json:"season_id" toml:"season_id" yaml:"season_id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *thSeason2EpisodeCachR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L thSeason2EpisodeCachL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var THSeason2EpisodeCachColumns = struct {
	EpisodeID string
	SeasonID  string
	CreatedAt string
	UpdatedAt string
}{
	EpisodeID: "episode_id",
	SeasonID:  "season_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var THSeason2EpisodeCachTableColumns = struct {
	EpisodeID string
	SeasonID  string
	CreatedAt string
	UpdatedAt string
}{
	EpisodeID: "th_season2_episode_caches.episode_id",
	SeasonID:  "th_season2_episode_caches.season_id",
	CreatedAt: "th_season2_episode_caches.created_at",
	UpdatedAt: "th_season2_episode_caches.updated_at",
}

// Generated where

var THSeason2EpisodeCachWhere = struct {
	EpisodeID whereHelperint64
	SeasonID  whereHelperint64
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	EpisodeID: whereHelperint64{field: "\"th_season2_episode_caches\".\"episode_id\""},
	SeasonID:  whereHelperint64{field: "\"th_season2_episode_caches\".\"season_id\""},
	CreatedAt: whereHelpertime_Time{field: "\"th_season2_episode_caches\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"th_season2_episode_caches\".\"updated_at\""},
}

// THSeason2EpisodeCachRels is where relationship names are stored.
var THSeason2EpisodeCachRels = struct {
}{}

// thSeason2EpisodeCachR is where relationships are stored.
type thSeason2EpisodeCachR struct {
}

// NewStruct creates a new relationship struct
func (*thSeason2EpisodeCachR) NewStruct() *thSeason2EpisodeCachR {
	return &thSeason2EpisodeCachR{}
}

// thSeason2EpisodeCachL is where Load methods for each relationship are stored.
type thSeason2EpisodeCachL struct{}

var (
	thSeason2EpisodeCachAllColumns            = []string{"episode_id", "season_id", "created_at", "updated_at"}
	thSeason2EpisodeCachColumnsWithoutDefault = []string{"episode_id", "season_id", "created_at", "updated_at"}
	thSeason2EpisodeCachColumnsWithDefault    = []string{}
	thSeason2EpisodeCachPrimaryKeyColumns     = []string{"episode_id"}
	thSeason2EpisodeCachGeneratedColumns      = []string{}
)

type (
	// THSeason2EpisodeCachSlice is an alias for a slice of pointers to THSeason2EpisodeCach.
	// This should almost always be used instead of []THSeason2EpisodeCach.
	THSeason2EpisodeCachSlice []*THSeason2EpisodeCach
	// THSeason2EpisodeCachHook is the signature for custom THSeason2EpisodeCach hook methods
	THSeason2EpisodeCachHook func(context.Context, boil.ContextExecutor, *THSeason2EpisodeCach) error

	thSeason2EpisodeCachQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	thSeason2EpisodeCachType                 = reflect.TypeOf(&THSeason2EpisodeCach{})
	thSeason2EpisodeCachMapping              = queries.MakeStructMapping(thSeason2EpisodeCachType)
	thSeason2EpisodeCachPrimaryKeyMapping, _ = queries.BindMapping(thSeason2EpisodeCachType, thSeason2EpisodeCachMapping, thSeason2EpisodeCachPrimaryKeyColumns)
	thSeason2EpisodeCachInsertCacheMut       sync.RWMutex
	thSeason2EpisodeCachInsertCache          = make(map[string]insertCache)
	thSeason2EpisodeCachUpdateCacheMut       sync.RWMutex
	thSeason2EpisodeCachUpdateCache          = make(map[string]updateCache)
	thSeason2EpisodeCachUpsertCacheMut       sync.RWMutex
	thSeason2EpisodeCachUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var thSeason2EpisodeCachAfterSelectHooks []THSeason2EpisodeCachHook

var thSeason2EpisodeCachBeforeInsertHooks []THSeason2EpisodeCachHook
var thSeason2EpisodeCachAfterInsertHooks []THSeason2EpisodeCachHook

var thSeason2EpisodeCachBeforeUpdateHooks []THSeason2EpisodeCachHook
var thSeason2EpisodeCachAfterUpdateHooks []THSeason2EpisodeCachHook

var thSeason2EpisodeCachBeforeDeleteHooks []THSeason2EpisodeCachHook
var thSeason2EpisodeCachAfterDeleteHooks []THSeason2EpisodeCachHook

var thSeason2EpisodeCachBeforeUpsertHooks []THSeason2EpisodeCachHook
var thSeason2EpisodeCachAfterUpsertHooks []THSeason2EpisodeCachHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *THSeason2EpisodeCach) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeason2EpisodeCachAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *THSeason2EpisodeCach) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeason2EpisodeCachBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *THSeason2EpisodeCach) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeason2EpisodeCachAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *THSeason2EpisodeCach) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeason2EpisodeCachBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *THSeason2EpisodeCach) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeason2EpisodeCachAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *THSeason2EpisodeCach) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeason2EpisodeCachBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *THSeason2EpisodeCach) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeason2EpisodeCachAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *THSeason2EpisodeCach) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeason2EpisodeCachBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *THSeason2EpisodeCach) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thSeason2EpisodeCachAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTHSeason2EpisodeCachHook registers your hook function for all future operations.
func AddTHSeason2EpisodeCachHook(hookPoint boil.HookPoint, thSeason2EpisodeCachHook THSeason2EpisodeCachHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		thSeason2EpisodeCachAfterSelectHooks = append(thSeason2EpisodeCachAfterSelectHooks, thSeason2EpisodeCachHook)
	case boil.BeforeInsertHook:
		thSeason2EpisodeCachBeforeInsertHooks = append(thSeason2EpisodeCachBeforeInsertHooks, thSeason2EpisodeCachHook)
	case boil.AfterInsertHook:
		thSeason2EpisodeCachAfterInsertHooks = append(thSeason2EpisodeCachAfterInsertHooks, thSeason2EpisodeCachHook)
	case boil.BeforeUpdateHook:
		thSeason2EpisodeCachBeforeUpdateHooks = append(thSeason2EpisodeCachBeforeUpdateHooks, thSeason2EpisodeCachHook)
	case boil.AfterUpdateHook:
		thSeason2EpisodeCachAfterUpdateHooks = append(thSeason2EpisodeCachAfterUpdateHooks, thSeason2EpisodeCachHook)
	case boil.BeforeDeleteHook:
		thSeason2EpisodeCachBeforeDeleteHooks = append(thSeason2EpisodeCachBeforeDeleteHooks, thSeason2EpisodeCachHook)
	case boil.AfterDeleteHook:
		thSeason2EpisodeCachAfterDeleteHooks = append(thSeason2EpisodeCachAfterDeleteHooks, thSeason2EpisodeCachHook)
	case boil.BeforeUpsertHook:
		thSeason2EpisodeCachBeforeUpsertHooks = append(thSeason2EpisodeCachBeforeUpsertHooks, thSeason2EpisodeCachHook)
	case boil.AfterUpsertHook:
		thSeason2EpisodeCachAfterUpsertHooks = append(thSeason2EpisodeCachAfterUpsertHooks, thSeason2EpisodeCachHook)
	}
}

// One returns a single thSeason2EpisodeCach record from the query.
func (q thSeason2EpisodeCachQuery) One(ctx context.Context, exec boil.ContextExecutor) (*THSeason2EpisodeCach, error) {
	o := &THSeason2EpisodeCach{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for th_season2_episode_caches")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all THSeason2EpisodeCach records from the query.
func (q thSeason2EpisodeCachQuery) All(ctx context.Context, exec boil.ContextExecutor) (THSeason2EpisodeCachSlice, error) {
	var o []*THSeason2EpisodeCach

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to THSeason2EpisodeCach slice")
	}

	if len(thSeason2EpisodeCachAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all THSeason2EpisodeCach records in the query.
func (q thSeason2EpisodeCachQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count th_season2_episode_caches rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q thSeason2EpisodeCachQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if th_season2_episode_caches exists")
	}

	return count > 0, nil
}

// THSeason2EpisodeCaches retrieves all the records using an executor.
func THSeason2EpisodeCaches(mods ...qm.QueryMod) thSeason2EpisodeCachQuery {
	mods = append(mods, qm.From("\"th_season2_episode_caches\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"th_season2_episode_caches\".*"})
	}

	return thSeason2EpisodeCachQuery{q}
}

// FindTHSeason2EpisodeCach retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTHSeason2EpisodeCach(ctx context.Context, exec boil.ContextExecutor, episodeID int64, selectCols ...string) (*THSeason2EpisodeCach, error) {
	thSeason2EpisodeCachObj := &THSeason2EpisodeCach{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"th_season2_episode_caches\" where \"episode_id\"=$1", sel,
	)

	q := queries.Raw(query, episodeID)

	err := q.Bind(ctx, exec, thSeason2EpisodeCachObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from th_season2_episode_caches")
	}

	if err = thSeason2EpisodeCachObj.doAfterSelectHooks(ctx, exec); err != nil {
		return thSeason2EpisodeCachObj, err
	}

	return thSeason2EpisodeCachObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *THSeason2EpisodeCach) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no th_season2_episode_caches provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(thSeason2EpisodeCachColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	thSeason2EpisodeCachInsertCacheMut.RLock()
	cache, cached := thSeason2EpisodeCachInsertCache[key]
	thSeason2EpisodeCachInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			thSeason2EpisodeCachAllColumns,
			thSeason2EpisodeCachColumnsWithDefault,
			thSeason2EpisodeCachColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(thSeason2EpisodeCachType, thSeason2EpisodeCachMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(thSeason2EpisodeCachType, thSeason2EpisodeCachMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"th_season2_episode_caches\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"th_season2_episode_caches\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into th_season2_episode_caches")
	}

	if !cached {
		thSeason2EpisodeCachInsertCacheMut.Lock()
		thSeason2EpisodeCachInsertCache[key] = cache
		thSeason2EpisodeCachInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the THSeason2EpisodeCach.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *THSeason2EpisodeCach) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	thSeason2EpisodeCachUpdateCacheMut.RLock()
	cache, cached := thSeason2EpisodeCachUpdateCache[key]
	thSeason2EpisodeCachUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			thSeason2EpisodeCachAllColumns,
			thSeason2EpisodeCachPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update th_season2_episode_caches, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"th_season2_episode_caches\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, thSeason2EpisodeCachPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(thSeason2EpisodeCachType, thSeason2EpisodeCachMapping, append(wl, thSeason2EpisodeCachPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update th_season2_episode_caches row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for th_season2_episode_caches")
	}

	if !cached {
		thSeason2EpisodeCachUpdateCacheMut.Lock()
		thSeason2EpisodeCachUpdateCache[key] = cache
		thSeason2EpisodeCachUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q thSeason2EpisodeCachQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for th_season2_episode_caches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for th_season2_episode_caches")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o THSeason2EpisodeCachSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), thSeason2EpisodeCachPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"th_season2_episode_caches\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, thSeason2EpisodeCachPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in thSeason2EpisodeCach slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all thSeason2EpisodeCach")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *THSeason2EpisodeCach) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no th_season2_episode_caches provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(thSeason2EpisodeCachColumnsWithDefault, o)

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

	thSeason2EpisodeCachUpsertCacheMut.RLock()
	cache, cached := thSeason2EpisodeCachUpsertCache[key]
	thSeason2EpisodeCachUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			thSeason2EpisodeCachAllColumns,
			thSeason2EpisodeCachColumnsWithDefault,
			thSeason2EpisodeCachColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			thSeason2EpisodeCachAllColumns,
			thSeason2EpisodeCachPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert th_season2_episode_caches, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(thSeason2EpisodeCachPrimaryKeyColumns))
			copy(conflict, thSeason2EpisodeCachPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"th_season2_episode_caches\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(thSeason2EpisodeCachType, thSeason2EpisodeCachMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(thSeason2EpisodeCachType, thSeason2EpisodeCachMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert th_season2_episode_caches")
	}

	if !cached {
		thSeason2EpisodeCachUpsertCacheMut.Lock()
		thSeason2EpisodeCachUpsertCache[key] = cache
		thSeason2EpisodeCachUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single THSeason2EpisodeCach record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *THSeason2EpisodeCach) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no THSeason2EpisodeCach provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), thSeason2EpisodeCachPrimaryKeyMapping)
	sql := "DELETE FROM \"th_season2_episode_caches\" WHERE \"episode_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from th_season2_episode_caches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for th_season2_episode_caches")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q thSeason2EpisodeCachQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no thSeason2EpisodeCachQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from th_season2_episode_caches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for th_season2_episode_caches")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o THSeason2EpisodeCachSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(thSeason2EpisodeCachBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), thSeason2EpisodeCachPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"th_season2_episode_caches\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, thSeason2EpisodeCachPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from thSeason2EpisodeCach slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for th_season2_episode_caches")
	}

	if len(thSeason2EpisodeCachAfterDeleteHooks) != 0 {
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
func (o *THSeason2EpisodeCach) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTHSeason2EpisodeCach(ctx, exec, o.EpisodeID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *THSeason2EpisodeCachSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := THSeason2EpisodeCachSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), thSeason2EpisodeCachPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"th_season2_episode_caches\".* FROM \"th_season2_episode_caches\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, thSeason2EpisodeCachPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in THSeason2EpisodeCachSlice")
	}

	*o = slice

	return nil
}

// THSeason2EpisodeCachExists checks if the THSeason2EpisodeCach row exists.
func THSeason2EpisodeCachExists(ctx context.Context, exec boil.ContextExecutor, episodeID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"th_season2_episode_caches\" where \"episode_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, episodeID)
	}
	row := exec.QueryRowContext(ctx, sql, episodeID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if th_season2_episode_caches exists")
	}

	return exists, nil
}

// Exists checks if the THSeason2EpisodeCach row exists.
func (o *THSeason2EpisodeCach) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return THSeason2EpisodeCachExists(ctx, exec, o.EpisodeID)
}
