package database

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
