package db

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"embed"

	migrate "github.com/rubenv/sql-migrate"
)

//go:embed sql/migrations/*
var migrationsFS embed.FS

func RunMigrations(pgConnectionString string) error {
	tries := 10
	db, err := sql.Open("postgres", pgConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()
	for {
		if tries < 0 {
			return errors.New("ran out of retries for migrations")
		}
		err = db.Ping()
		if err != nil {
			tries = tries - 1
			time.Sleep(2 * time.Second)
			continue
		}
		err := runMigrations(db)
		if err != nil {
			fmt.Println("issue running migrations", "error", err, "tries_left", tries)
			return err
		}
		return nil
	}
}

func runMigrations( db *sql.DB) error {
	migrations := migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrationsFS,
		Root:       "sql/migrations",
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}

	fmt.Printf("Applied %d successful migrations!\n", n)

	return nil
}
