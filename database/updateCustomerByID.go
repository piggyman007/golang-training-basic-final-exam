package database

// UpdateCustomerByID - update customer by ID
func UpdateCustomerByID(id int, name string, email string, status string) error {
	stmt, err := db.Prepare(`
		UPDATE customers
		SET
			name = $2,
			email = $3,
			status = $4
		WHERE id = $1;
	`)

	if err != nil {
		return err
	}
	_, err = stmt.Exec(id, name, email, status)

	return err
}
