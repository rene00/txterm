package store

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
)

const (
	DBName = "txterm.db"
)

type Store struct {
	db  *sql.DB
	dir string
}

func (s *Store) init() error {
	if err := ensureDir(s.dir); err != nil {
		return err
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
