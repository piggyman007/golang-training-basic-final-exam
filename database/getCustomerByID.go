package database

import (
	"database/sql"
)

// GetCustomerByID - get customer by ID
func GetCustomerByID(id int) (*sql.Row, error) {
	stmt, err := db.Prepare(`
		SELECT id, name, email, status 
		FROM customers
		WHERE id = $1;
	`)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(id)

	return row, nil
}
