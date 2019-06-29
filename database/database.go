package database

import (
	"database/sql"
	"os"
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	return db, err
}

// InitTable if not exists
func InitTable() {
	db, err := connectDB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

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

// CreateCustomer record in table
func CreateCustomer(name string, email string, status string) (*sql.Row, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
		INSERT INTO customers (name, email, status) VALUES ($1, $2, $3) RETURNING id
	`
	row := db.QueryRow(query, name, email, status)

	return row, nil
}
