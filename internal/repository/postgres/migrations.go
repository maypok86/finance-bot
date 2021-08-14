package postgres

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
)

func runMigrations(migrationsPath string, dsn string) error {
	if migrationsPath == "" {
		return nil
	}
	if dsn == "" {
		return errors.New("No postgres dsn provided")
	}
	m, err := migrate.New(migrationsPath, dsn)
	if err != nil {
		return errors.Wrap(err, "[migrate.New] failed")
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return errors.Wrap(err, "migrate up failed")
	}
	return nil
}