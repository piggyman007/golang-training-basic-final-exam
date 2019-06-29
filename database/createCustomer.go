package database

import (
	"database/sql"
)

// CreateCustomer - create customer in table
func CreateCustomer(name string, email string, status string) (*sql.Row, error) {
	query := `
		INSERT INTO customers (name, email, status) VALUES ($1, $2, $3) RETURNING id
	`
	row := db.QueryRow(query, name, email, status)

	return row, nil
}
