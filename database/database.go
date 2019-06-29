package database

import (
	"database/sql"
	"os"
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	return db, err
}
