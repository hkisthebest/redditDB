package db

import (
	"log"
	"os"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
  _ "github.com/golang-migrate/migrate/v4/database/postgres"
)

func Migrate() {
	m, err := migrate.New(
		"file://migrations",
		os.Getenv("DB_URL"))
	if err != nil {
    log.Fatal("error opening connection: ", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("error applying migration: ", err)
	}
}
