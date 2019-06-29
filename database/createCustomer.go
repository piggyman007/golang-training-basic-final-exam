package database

import "database/sql"

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
