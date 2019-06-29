package database

import (
	"database/sql"
	"os"
)

var db *sql.DB

// New - init db connection and table
func New() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err.Error())
	}

	createTb := `
		CREATE TABLE IF NOT EXISTS customers(
			id		SERIAL PRIMARY KEY,
			name	TEXT,
			email	TEXT,
			status	TEXT
		);
	`
	_, err = db.Exec(createTb)
	if err != nil {
		panic(err.Error())
	}
}

// Close - close db connection
func Close() {
	db.Close()
}
