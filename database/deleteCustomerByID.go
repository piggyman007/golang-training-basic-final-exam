package database

// DeleteCustomerByID - delete customer by ID
func DeleteCustomerByID(id int) error {
	stmt, err := db.Prepare(`
		DELETE FROM customers
		WHERE id = $1;
	`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)

	return err
}
