package store

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"txterm/db"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/mattn/go-sqlite3"
)

const (
	DBName = "txterm.db"
)

type Store struct {
	DB  *sql.DB
	dir string
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

	return nil
}

func (s *Store) initDB() error {
	db, err := sql.Open("sqlite3", s.DBPath())
	if err != nil {
		return err
	}
	s.DB = db
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

func New() (Store, error) {
	store := Store{}
	if store.dir == "" {
		dir, err := defaultStorePath()
		if err != nil {
			return store, err
		}
		store.dir = dir
	}

	if err := store.init(); err != nil {
		return store, err
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
