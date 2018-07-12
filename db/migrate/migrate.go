package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {
	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	driver, _ := postgres.WithInstance(db, &postgres.Config{})
	m, error := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver)
	if error == nil {
		m.Up()
	} else {
		fmt.Println("Error initialiaing migrations", error)
	}
}
