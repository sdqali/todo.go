package main

import (
	"fmt"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"github.com/sdqali/todo/db"
)

func main() {
	db := db.GetDb()
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
