package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func createTableUsers(db *sql.DB) error {
	err := userTable(db)
	if err != nil {
		return err
	}
	return nil
}

func userTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users (
		ID SERIAL PRIMARY KEY,
		Login VARCHAR(256) NOT NULL UNIQUE,
		Password VARCHAR(256) NOT NULL
	);`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
