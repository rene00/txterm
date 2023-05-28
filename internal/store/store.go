package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"time"
	"txterm/internal/db"
	"txterm/internal/db/query"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/mattn/go-sqlite3"
)

const (
	DBName = "txterm.db"
)

type Store struct {
	db      *sql.DB
	dir     string
	queries *query.Queries
}

func (s *Store) DBPath() string {
	return filepath.Join(s.dir, DBName)
}

func (s *Store) init() error {
	if err := ensureDir(s.dir); err != nil {
		return fmt.Errorf("ensuredir: %w", err)
	}

	if err := s.migrate(); err != nil {
		return fmt.Errorf("migrate: %w", err)
	}

	if err := s.initDB(); err != nil {
		return fmt.Errorf("initDB: %w", err)
	}

	s.queries = query.New(s.db)

	return nil
}

func (s *Store) initDB() error {
	db, err := sql.Open("sqlite3", s.DBPath())
	if err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Store) migrate() error {
	sourceDriver, err := iofs.New(db.FSMigrations, "migrations")
	if err != nil {
		return err
	}

	path := "sqlite3://" + s.DBPath()
	mg, err := migrate.NewWithSourceInstance("iofs", sourceDriver, path)
	if err != nil {
		return err
	}

	err = mg.Up()
	switch {
	case errors.Is(err, migrate.ErrNoChange):
	case err != nil:
		return fmt.Errorf("up: %w", err)
	}
	return nil
}

func New() (*Store, error) {
	store := &Store{}
	if store.dir == "" {
		dir, err := defaultStorePath()
		if err != nil {
			return nil, err
		}
		store.dir = dir
	}

	if err := store.init(); err != nil {
		return nil, fmt.Errorf("init: %w", err)
	}

	return store, nil
}

func defaultStorePath() (string, error) {
	hd, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	res := filepath.Join(hd, ".config", "txterm")
	return res, nil
}

func ensureDir(dir string) error {
	info, err := os.Stat(dir)
	switch {
	case os.IsNotExist(err):
		if err := os.Mkdir(dir, 0755); err != nil {
			return err
		}
	case err != nil:
		return err
	case info.IsDir():
	case !info.IsDir():
		return fmt.Errorf("%q exists but is a file", dir)
	}
	return nil
}

// SaveImport persists the import "run" to storage.
func (s *Store) SaveImport(ctx context.Context, filename string, balanceAmount big.Rat, dateAsOf time.Time) (*Import, error) {
	imprt, err := s.queries.CreateImport(ctx, query.CreateImportParams{
		DateCreated:      time.Now(),
		Filename:         filename,
		BalanceAmountNum: sql.NullInt64{Int64: balanceAmount.Num().Int64(), Valid: true},
		BalanceAmountDen: sql.NullInt64{Int64: balanceAmount.Denom().Int64(), Valid: true},
		DateAsOf:         sql.NullTime{Time: dateAsOf, Valid: true},
	})
	if err != nil {
		return nil, fmt.Errorf("save import: %w", err)
	}
	return &Import{imprt}, nil
}

type Import struct {
	query.Import
}

// SaveTransaction persists the transaction to storage.
//
// Duplicates can occur within the same import. Duplicates that occur outside
// of a single will not be persisted.
func (s *Store) SaveTransaction(ctx context.Context, datePosted time.Time, memo string, amount big.Rat, imprt Import) (*Transaction, error) {

	duplicates, err := s.queries.GetDuplicateTx(ctx, query.GetDuplicateTxParams{DatePosted: datePosted, Memo: memo, AmountNum: amount.Num().Int64(), AmountDen: amount.Denom().Int64()})
	if err != nil {
		return nil, fmt.Errorf("check duplicates: %w", err)
	}

	duplicateExists := false
	for _, duplicate := range duplicates {
		if duplicate.ImportID != imprt.ID {
			duplicateExists = true
			break
		}
	}
	if duplicateExists {
		return nil, TransactionDuplicateError()
	}

	tx, err := s.queries.CreateTx(ctx, query.CreateTxParams{
		DateCreated: time.Now(),
		DatePosted:  datePosted,
		Memo:        memo,
		AmountNum:   amount.Num().Int64(),
		AmountDen:   amount.Denom().Int64(),
		ImportID:    imprt.ID,
	})
	if err != nil {
		return nil, fmt.Errorf("save transaction: %w", err)
	}
	return &Transaction{tx}, nil
}

// GetAmountFromNumDenom takes a numerator & denominator and prints a human readable dollar amount.
func GetAmountFromNumDenom(n, d int64) string {
	return strings.TrimRight(strings.TrimRight(big.NewRat(n, d).FloatString(100), "0"), ".")
}
