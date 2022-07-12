// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testTHSeasonEpisodeCaches(t *testing.T) {
	t.Parallel()

	query := THSeasonEpisodeCaches()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTHSeasonEpisodeCachesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSeasonEpisodeCach{}
	if err = randomize.Struct(seed, o, thSeasonEpisodeCachDBTypes, true, thSeasonEpisodeCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeasonEpisodeCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := THSeasonEpisodeCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTHSeasonEpisodeCachesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSeasonEpisodeCach{}
	if err = randomize.Struct(seed, o, thSeasonEpisodeCachDBTypes, true, thSeasonEpisodeCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeasonEpisodeCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := THSeasonEpisodeCaches().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := THSeasonEpisodeCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTHSeasonEpisodeCachesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSeasonEpisodeCach{}
	if err = randomize.Struct(seed, o, thSeasonEpisodeCachDBTypes, true, thSeasonEpisodeCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeasonEpisodeCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := THSeasonEpisodeCachSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := THSeasonEpisodeCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTHSeasonEpisodeCachesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSeasonEpisodeCach{}
	if err = randomize.Struct(seed, o, thSeasonEpisodeCachDBTypes, true, thSeasonEpisodeCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeasonEpisodeCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := THSeasonEpisodeCachExists(ctx, tx, o.EpisodeID)
	if err != nil {
		t.Errorf("Unable to check if THSeasonEpisodeCach exists: %s", err)
	}
	if !e {
		t.Errorf("Expected THSeasonEpisodeCachExists to return true, but got false.")
	}
}

func testTHSeasonEpisodeCachesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSeasonEpisodeCach{}
	if err = randomize.Struct(seed, o, thSeasonEpisodeCachDBTypes, true, thSeasonEpisodeCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeasonEpisodeCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	thSeasonEpisodeCachFound, err := FindTHSeasonEpisodeCach(ctx, tx, o.EpisodeID)
	if err != nil {
		t.Error(err)
	}

	if thSeasonEpisodeCachFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTHSeasonEpisodeCachesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSeasonEpisodeCach{}
	if err = randomize.Struct(seed, o, thSeasonEpisodeCachDBTypes, true, thSeasonEpisodeCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeasonEpisodeCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = THSeasonEpisodeCaches().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTHSeasonEpisodeCachesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSeasonEpisodeCach{}
	if err = randomize.Struct(seed, o, thSeasonEpisodeCachDBTypes, true, thSeasonEpisodeCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeasonEpisodeCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := THSeasonEpisodeCaches().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTHSeasonEpisodeCachesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	thSeasonEpisodeCachOne := &THSeasonEpisodeCach{}
	thSeasonEpisodeCachTwo := &THSeasonEpisodeCach{}
	if err = randomize.Struct(seed, thSeasonEpisodeCachOne, thSeasonEpisodeCachDBTypes, false, thSeasonEpisodeCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeasonEpisodeCach struct: %s", err)
	}
	if err = randomize.Struct(seed, thSeasonEpisodeCachTwo, thSeasonEpisodeCachDBTypes, false, thSeasonEpisodeCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeasonEpisodeCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = thSeasonEpisodeCachOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = thSeasonEpisodeCachTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := THSeasonEpisodeCaches().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTHSeasonEpisodeCachesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	thSeasonEpisodeCachOne := &THSeasonEpisodeCach{}
	thSeasonEpisodeCachTwo := &THSeasonEpisodeCach{}
	if err = randomize.Struct(seed, thSeasonEpisodeCachOne, thSeasonEpisodeCachDBTypes, false, thSeasonEpisodeCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeasonEpisodeCach struct: %s", err)
	}
	if err = randomize.Struct(seed, thSeasonEpisodeCachTwo, thSeasonEpisodeCachDBTypes, false, thSeasonEpisodeCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeasonEpisodeCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = thSeasonEpisodeCachOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = thSeasonEpisodeCachTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := THSeasonEpisodeCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func thSeasonEpisodeCachBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *THSeasonEpisodeCach) error {
	*o = THSeasonEpisodeCach{}
	return nil
}

func thSeasonEpisodeCachAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *THSeasonEpisodeCach) error {
	*o = THSeasonEpisodeCach{}
	return nil
}

func thSeasonEpisodeCachAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *THSeasonEpisodeCach) error {
	*o = THSeasonEpisodeCach{}
	return nil
}

func thSeasonEpisodeCachBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *THSeasonEpisodeCach) error {
	*o = THSeasonEpisodeCach{}
	return nil
}

func thSeasonEpisodeCachAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *THSeasonEpisodeCach) error {
	*o = THSeasonEpisodeCach{}
	return nil
}

func thSeasonEpisodeCachBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *THSeasonEpisodeCach) error {
	*o = THSeasonEpisodeCach{}
	return nil
}

func thSeasonEpisodeCachAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *THSeasonEpisodeCach) error {
	*o = THSeasonEpisodeCach{}
	return nil
}

func thSeasonEpisodeCachBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *THSeasonEpisodeCach) error {
	*o = THSeasonEpisodeCach{}
	return nil
}

func thSeasonEpisodeCachAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *THSeasonEpisodeCach) error {
	*o = THSeasonEpisodeCach{}
	return nil
}

func testTHSeasonEpisodeCachesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &THSeasonEpisodeCach{}
	o := &THSeasonEpisodeCach{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, thSeasonEpisodeCachDBTypes, false); err != nil {
		t.Errorf("Unable to randomize THSeasonEpisodeCach object: %s", err)
	}

	AddTHSeasonEpisodeCachHook(boil.BeforeInsertHook, thSeasonEpisodeCachBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	thSeasonEpisodeCachBeforeInsertHooks = []THSeasonEpisodeCachHook{}

	AddTHSeasonEpisodeCachHook(boil.AfterInsertHook, thSeasonEpisodeCachAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	thSeasonEpisodeCachAfterInsertHooks = []THSeasonEpisodeCachHook{}

	AddTHSeasonEpisodeCachHook(boil.AfterSelectHook, thSeasonEpisodeCachAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	thSeasonEpisodeCachAfterSelectHooks = []THSeasonEpisodeCachHook{}

	AddTHSeasonEpisodeCachHook(boil.BeforeUpdateHook, thSeasonEpisodeCachBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	thSeasonEpisodeCachBeforeUpdateHooks = []THSeasonEpisodeCachHook{}

	AddTHSeasonEpisodeCachHook(boil.AfterUpdateHook, thSeasonEpisodeCachAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	thSeasonEpisodeCachAfterUpdateHooks = []THSeasonEpisodeCachHook{}

	AddTHSeasonEpisodeCachHook(boil.BeforeDeleteHook, thSeasonEpisodeCachBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	thSeasonEpisodeCachBeforeDeleteHooks = []THSeasonEpisodeCachHook{}

	AddTHSeasonEpisodeCachHook(boil.AfterDeleteHook, thSeasonEpisodeCachAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	thSeasonEpisodeCachAfterDeleteHooks = []THSeasonEpisodeCachHook{}

	AddTHSeasonEpisodeCachHook(boil.BeforeUpsertHook, thSeasonEpisodeCachBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	thSeasonEpisodeCachBeforeUpsertHooks = []THSeasonEpisodeCachHook{}

	AddTHSeasonEpisodeCachHook(boil.AfterUpsertHook, thSeasonEpisodeCachAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	thSeasonEpisodeCachAfterUpsertHooks = []THSeasonEpisodeCachHook{}
}

func testTHSeasonEpisodeCachesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSeasonEpisodeCach{}
	if err = randomize.Struct(seed, o, thSeasonEpisodeCachDBTypes, true, thSeasonEpisodeCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeasonEpisodeCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := THSeasonEpisodeCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTHSeasonEpisodeCachesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSeasonEpisodeCach{}
	if err = randomize.Struct(seed, o, thSeasonEpisodeCachDBTypes, true); err != nil {
		t.Errorf("Unable to randomize THSeasonEpisodeCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(thSeasonEpisodeCachColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := THSeasonEpisodeCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTHSeasonEpisodeCachesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSeasonEpisodeCach{}
	if err = randomize.Struct(seed, o, thSeasonEpisodeCachDBTypes, true, thSeasonEpisodeCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeasonEpisodeCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTHSeasonEpisodeCachesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSeasonEpisodeCach{}
	if err = randomize.Struct(seed, o, thSeasonEpisodeCachDBTypes, true, thSeasonEpisodeCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeasonEpisodeCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := THSeasonEpisodeCachSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTHSeasonEpisodeCachesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSeasonEpisodeCach{}
	if err = randomize.Struct(seed, o, thSeasonEpisodeCachDBTypes, true, thSeasonEpisodeCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeasonEpisodeCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := THSeasonEpisodeCaches().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	thSeasonEpisodeCachDBTypes = map[string]string{`EpisodeID`: `bigint`, `SeasonID`: `bigint`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_                          = bytes.MinRead
)

func testTHSeasonEpisodeCachesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(thSeasonEpisodeCachPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(thSeasonEpisodeCachAllColumns) == len(thSeasonEpisodeCachPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &THSeasonEpisodeCach{}
	if err = randomize.Struct(seed, o, thSeasonEpisodeCachDBTypes, true, thSeasonEpisodeCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeasonEpisodeCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := THSeasonEpisodeCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, thSeasonEpisodeCachDBTypes, true, thSeasonEpisodeCachPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize THSeasonEpisodeCach struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTHSeasonEpisodeCachesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(thSeasonEpisodeCachAllColumns) == len(thSeasonEpisodeCachPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &THSeasonEpisodeCach{}
	if err = randomize.Struct(seed, o, thSeasonEpisodeCachDBTypes, true, thSeasonEpisodeCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeasonEpisodeCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := THSeasonEpisodeCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, thSeasonEpisodeCachDBTypes, true, thSeasonEpisodeCachPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize THSeasonEpisodeCach struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(thSeasonEpisodeCachAllColumns, thSeasonEpisodeCachPrimaryKeyColumns) {
		fields = thSeasonEpisodeCachAllColumns
	} else {
		fields = strmangle.SetComplement(
			thSeasonEpisodeCachAllColumns,
			thSeasonEpisodeCachPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := THSeasonEpisodeCachSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTHSeasonEpisodeCachesUpsert(t *testing.T) {
	t.Parallel()

	if len(thSeasonEpisodeCachAllColumns) == len(thSeasonEpisodeCachPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := THSeasonEpisodeCach{}
	if err = randomize.Struct(seed, &o, thSeasonEpisodeCachDBTypes, true); err != nil {
		t.Errorf("Unable to randomize THSeasonEpisodeCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert THSeasonEpisodeCach: %s", err)
	}

	count, err := THSeasonEpisodeCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, thSeasonEpisodeCachDBTypes, false, thSeasonEpisodeCachPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize THSeasonEpisodeCach struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert THSeasonEpisodeCach: %s", err)
	}

	count, err = THSeasonEpisodeCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
