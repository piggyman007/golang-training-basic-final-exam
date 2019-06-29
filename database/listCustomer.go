package database

import (
	"database/sql"
)

// ListCustomer - list all customers in table
func ListCustomer() (*sql.Rows, error) {
	stmt, err := db.Prepare(`
		SELECT id, name, email, status 
		FROM customers;
	`)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()

	return rows, err
}
