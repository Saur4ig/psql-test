package main

import (
	"log"

	"github.com/Saur4ig/psql-test/config"
	"github.com/Saur4ig/psql-test/db"
)

func main() {
	// Load configuration
	cfg := config.New()

	// Run database migrations
	if err := db.RunMigrations(cfg.DSN()); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}
