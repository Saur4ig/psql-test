package db

import (
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// RunMigrations applies all database migrations
func RunMigrations(dsn string) error {
	log.Printf("Running migrations with DSN: %s", dsn)

	// Initialize migrator
	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		return err
	}

	// Run migrations
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	log.Println("Migrations completed successfully")
	return nil
}
