// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

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

// Split is an object representing the database table.
type Split struct {
	ID            int64 `boil:"id" json:"id" toml:"id" yaml:"id"`
	TransactionID int64 `boil:"tx_id" json:"tx_id" toml:"tx_id" yaml:"tx_id"`
	AccountID     int64 `boil:"account_id" json:"account_id" toml:"account_id" yaml:"account_id"`
	ValueNum      int64 `boil:"value_num" json:"value_num" toml:"value_num" yaml:"value_num"`
	ValueDenom    int64 `boil:"value_denom" json:"value_denom" toml:"value_denom" yaml:"value_denom"`

	R *splitR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L splitL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SplitColumns = struct {
	ID            string
	TransactionID string
	AccountID     string
	ValueNum      string
	ValueDenom    string
}{
	ID:            "id",
	TransactionID: "tx_id",
	AccountID:     "account_id",
	ValueNum:      "value_num",
	ValueDenom:    "value_denom",
}

var SplitTableColumns = struct {
	ID            string
	TransactionID string
	AccountID     string
	ValueNum      string
	ValueDenom    string
}{
	ID:            "split.id",
	TransactionID: "split.tx_id",
	AccountID:     "split.account_id",
	ValueNum:      "split.value_num",
	ValueDenom:    "split.value_denom",
}

// Generated where

var SplitWhere = struct {
	ID            whereHelperint64
	TransactionID whereHelperint64
	AccountID     whereHelperint64
	ValueNum      whereHelperint64
	ValueDenom    whereHelperint64
}{
	ID:            whereHelperint64{field: "\"split\".\"id\""},
	TransactionID: whereHelperint64{field: "\"split\".\"tx_id\""},
	AccountID:     whereHelperint64{field: "\"split\".\"account_id\""},
	ValueNum:      whereHelperint64{field: "\"split\".\"value_num\""},
	ValueDenom:    whereHelperint64{field: "\"split\".\"value_denom\""},
}

// SplitRels is where relationship names are stored.
var SplitRels = struct {
	Account string
	TX      string
}{
	Account: "Account",
	TX:      "TX",
}

// splitR is where relationships are stored.
type splitR struct {
	Account *Account     `boil:"Account" json:"Account" toml:"Account" yaml:"Account"`
	TX      *Transaction `boil:"TX" json:"TX" toml:"TX" yaml:"TX"`
}

// NewStruct creates a new relationship struct
func (*splitR) NewStruct() *splitR {
	return &splitR{}
}

func (r *splitR) GetAccount() *Account {
	if r == nil {
		return nil
	}
	return r.Account
}

func (r *splitR) GetTX() *Transaction {
	if r == nil {
		return nil
	}
	return r.TX
}

// splitL is where Load methods for each relationship are stored.
type splitL struct{}

var (
	splitAllColumns            = []string{"id", "tx_id", "account_id", "value_num", "value_denom"}
	splitColumnsWithoutDefault = []string{"tx_id", "account_id", "value_num", "value_denom"}
	splitColumnsWithDefault    = []string{"id"}
	splitPrimaryKeyColumns     = []string{"id"}
	splitGeneratedColumns      = []string{"id"}
)

type (
	// SplitSlice is an alias for a slice of pointers to Split.
	// This should almost always be used instead of []Split.
	SplitSlice []*Split
	// SplitHook is the signature for custom Split hook methods
	SplitHook func(context.Context, boil.ContextExecutor, *Split) error

	splitQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	splitType                 = reflect.TypeOf(&Split{})
	splitMapping              = queries.MakeStructMapping(splitType)
	splitPrimaryKeyMapping, _ = queries.BindMapping(splitType, splitMapping, splitPrimaryKeyColumns)
	splitInsertCacheMut       sync.RWMutex
	splitInsertCache          = make(map[string]insertCache)
	splitUpdateCacheMut       sync.RWMutex
	splitUpdateCache          = make(map[string]updateCache)
	splitUpsertCacheMut       sync.RWMutex
	splitUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var splitAfterSelectHooks []SplitHook

var splitBeforeInsertHooks []SplitHook
var splitAfterInsertHooks []SplitHook

var splitBeforeUpdateHooks []SplitHook
var splitAfterUpdateHooks []SplitHook

var splitBeforeDeleteHooks []SplitHook
var splitAfterDeleteHooks []SplitHook

var splitBeforeUpsertHooks []SplitHook
var splitAfterUpsertHooks []SplitHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Split) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range splitAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Split) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range splitBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Split) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range splitAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Split) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range splitBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Split) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range splitAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Split) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range splitBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Split) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range splitAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Split) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range splitBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Split) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range splitAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSplitHook registers your hook function for all future operations.
func AddSplitHook(hookPoint boil.HookPoint, splitHook SplitHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		splitAfterSelectHooks = append(splitAfterSelectHooks, splitHook)
	case boil.BeforeInsertHook:
		splitBeforeInsertHooks = append(splitBeforeInsertHooks, splitHook)
	case boil.AfterInsertHook:
		splitAfterInsertHooks = append(splitAfterInsertHooks, splitHook)
	case boil.BeforeUpdateHook:
		splitBeforeUpdateHooks = append(splitBeforeUpdateHooks, splitHook)
	case boil.AfterUpdateHook:
		splitAfterUpdateHooks = append(splitAfterUpdateHooks, splitHook)
	case boil.BeforeDeleteHook:
		splitBeforeDeleteHooks = append(splitBeforeDeleteHooks, splitHook)
	case boil.AfterDeleteHook:
		splitAfterDeleteHooks = append(splitAfterDeleteHooks, splitHook)
	case boil.BeforeUpsertHook:
		splitBeforeUpsertHooks = append(splitBeforeUpsertHooks, splitHook)
	case boil.AfterUpsertHook:
		splitAfterUpsertHooks = append(splitAfterUpsertHooks, splitHook)
	}
}

// One returns a single split record from the query.
func (q splitQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Split, error) {
	o := &Split{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for split")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Split records from the query.
func (q splitQuery) All(ctx context.Context, exec boil.ContextExecutor) (SplitSlice, error) {
	var o []*Split

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to Split slice")
	}

	if len(splitAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Split records in the query.
func (q splitQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count split rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q splitQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if split exists")
	}

	return count > 0, nil
}

// Account pointed to by the foreign key.
func (o *Split) Account(mods ...qm.QueryMod) accountQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.AccountID),
	}

	queryMods = append(queryMods, mods...)

	return Accounts(queryMods...)
}

// TX pointed to by the foreign key.
func (o *Split) TX(mods ...qm.QueryMod) transactionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.TransactionID),
	}

	queryMods = append(queryMods, mods...)

	return Transactions(queryMods...)
}

// LoadAccount allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (splitL) LoadAccount(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSplit interface{}, mods queries.Applicator) error {
	var slice []*Split
	var object *Split

	if singular {
		var ok bool
		object, ok = maybeSplit.(*Split)
		if !ok {
			object = new(Split)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeSplit)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeSplit))
			}
		}
	} else {
		s, ok := maybeSplit.(*[]*Split)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeSplit)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeSplit))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &splitR{}
		}
		args = append(args, object.AccountID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &splitR{}
			}

			for _, a := range args {
				if a == obj.AccountID {
					continue Outer
				}
			}

			args = append(args, obj.AccountID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`account`),
		qm.WhereIn(`account.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Account")
	}

	var resultSlice []*Account
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Account")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for account")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for account")
	}

	if len(accountAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Account = foreign
		if foreign.R == nil {
			foreign.R = &accountR{}
		}
		foreign.R.Splits = append(foreign.R.Splits, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.AccountID == foreign.ID {
				local.R.Account = foreign
				if foreign.R == nil {
					foreign.R = &accountR{}
				}
				foreign.R.Splits = append(foreign.R.Splits, local)
				break
			}
		}
	}

	return nil
}

// LoadTX allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (splitL) LoadTX(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSplit interface{}, mods queries.Applicator) error {
	var slice []*Split
	var object *Split

	if singular {
		var ok bool
		object, ok = maybeSplit.(*Split)
		if !ok {
			object = new(Split)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeSplit)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeSplit))
			}
		}
	} else {
		s, ok := maybeSplit.(*[]*Split)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeSplit)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeSplit))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &splitR{}
		}
		args = append(args, object.TransactionID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &splitR{}
			}

			for _, a := range args {
				if a == obj.TransactionID {
					continue Outer
				}
			}

			args = append(args, obj.TransactionID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`tx`),
		qm.WhereIn(`tx.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Transaction")
	}

	var resultSlice []*Transaction
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Transaction")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for tx")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for tx")
	}

	if len(transactionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.TX = foreign
		if foreign.R == nil {
			foreign.R = &transactionR{}
		}
		foreign.R.Splits = append(foreign.R.Splits, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TransactionID == foreign.ID {
				local.R.TX = foreign
				if foreign.R == nil {
					foreign.R = &transactionR{}
				}
				foreign.R.Splits = append(foreign.R.Splits, local)
				break
			}
		}
	}

	return nil
}

// SetAccount of the split to the related item.
// Sets o.R.Account to related.
// Adds o to related.R.Splits.
func (o *Split) SetAccount(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Account) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"split\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"account_id"}),
		strmangle.WhereClause("\"", "\"", 0, splitPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AccountID = related.ID
	if o.R == nil {
		o.R = &splitR{
			Account: related,
		}
	} else {
		o.R.Account = related
	}

	if related.R == nil {
		related.R = &accountR{
			Splits: SplitSlice{o},
		}
	} else {
		related.R.Splits = append(related.R.Splits, o)
	}

	return nil
}

// SetTX of the split to the related item.
// Sets o.R.TX to related.
// Adds o to related.R.Splits.
func (o *Split) SetTX(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Transaction) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"split\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"tx_id"}),
		strmangle.WhereClause("\"", "\"", 0, splitPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TransactionID = related.ID
	if o.R == nil {
		o.R = &splitR{
			TX: related,
		}
	} else {
		o.R.TX = related
	}

	if related.R == nil {
		related.R = &transactionR{
			Splits: SplitSlice{o},
		}
	} else {
		related.R.Splits = append(related.R.Splits, o)
	}

	return nil
}

// Splits retrieves all the records using an executor.
func Splits(mods ...qm.QueryMod) splitQuery {
	mods = append(mods, qm.From("\"split\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"split\".*"})
	}

	return splitQuery{q}
}

// FindSplit retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSplit(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Split, error) {
	splitObj := &Split{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"split\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, splitObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from split")
	}

	if err = splitObj.doAfterSelectHooks(ctx, exec); err != nil {
		return splitObj, err
	}

	return splitObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Split) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no split provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(splitColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	splitInsertCacheMut.RLock()
	cache, cached := splitInsertCache[key]
	splitInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			splitAllColumns,
			splitColumnsWithDefault,
			splitColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, splitGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(splitType, splitMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(splitType, splitMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"split\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"split\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "model: unable to insert into split")
	}

	if !cached {
		splitInsertCacheMut.Lock()
		splitInsertCache[key] = cache
		splitInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Split.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Split) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	splitUpdateCacheMut.RLock()
	cache, cached := splitUpdateCache[key]
	splitUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			splitAllColumns,
			splitPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, splitGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update split, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"split\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, splitPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(splitType, splitMapping, append(wl, splitPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "model: unable to update split row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for split")
	}

	if !cached {
		splitUpdateCacheMut.Lock()
		splitUpdateCache[key] = cache
		splitUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q splitQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for split")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for split")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SplitSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("model: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), splitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"split\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, splitPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in split slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all split")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Split) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no split provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(splitColumnsWithDefault, o)

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

	splitUpsertCacheMut.RLock()
	cache, cached := splitUpsertCache[key]
	splitUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			splitAllColumns,
			splitColumnsWithDefault,
			splitColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			splitAllColumns,
			splitPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("model: unable to upsert split, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(splitPrimaryKeyColumns))
			copy(conflict, splitPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"split\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(splitType, splitMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(splitType, splitMapping, ret)
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
		return errors.Wrap(err, "model: unable to upsert split")
	}

	if !cached {
		splitUpsertCacheMut.Lock()
		splitUpsertCache[key] = cache
		splitUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Split record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Split) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no Split provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), splitPrimaryKeyMapping)
	sql := "DELETE FROM \"split\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from split")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for split")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q splitQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no splitQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from split")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for split")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SplitSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(splitBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), splitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"split\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, splitPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from split slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for split")
	}

	if len(splitAfterDeleteHooks) != 0 {
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
func (o *Split) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSplit(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SplitSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SplitSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), splitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"split\".* FROM \"split\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, splitPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in SplitSlice")
	}

	*o = slice

	return nil
}

// SplitExists checks if the Split row exists.
func SplitExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"split\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if split exists")
	}

	return exists, nil
}

// Exists checks if the Split row exists.
func (o *Split) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return SplitExists(ctx, exec, o.ID)
}
