package cli

import (
	"github.com/redcuckoo/go-labs-SE/lab3/assets"
	"github.com/redcuckoo/go-labs-SE/lab3/config"
	migrate "github.com/rubenv/sql-migrate"
	"log"
)

var migrations = &migrate.PackrMigrationSource{
	Box: assets.Migrations,
}

func MigrateUp(cfg config.Config) error {
	applied, err := migrate.Exec(cfg.DB, "postgres", migrations, migrate.Up)

	if err != nil {
		log.Fatalf("failed to apply migrations, err: %v", err)
		return err
	}

	log.Printf("migrations applied: %v", applied)
	return nil
}

func MigrateDown(cfg config.Config) error {
	applied, err := migrate.Exec(cfg.DB, "postgres", migrations, migrate.Down)

	if err != nil {
		log.Fatalf("failed to apply migrations, err: %v", err)
		return err

	}

	log.Printf("migrations applied: %v", applied)
	return nil
}
