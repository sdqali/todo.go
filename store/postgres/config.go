package postgres

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func GetDb() *sql.DB {
	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	return db
}
