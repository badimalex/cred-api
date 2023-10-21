package db

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateDB() error {
	driver, err := postgres.WithInstance(DB.DB(), &postgres.Config{})
	if err != nil {
		return err
	}

	migrationsDir := "file://migrations"
	m, err := migrate.NewWithDatabaseInstance(
		migrationsDir,
		"postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
