package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func createTableUsers(db *sql.DB) error {
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

func (s *StoreDB) CreateUser(login, password string) error {
	query := `INSERT INTO users (login, password) VALUES ($1,$2)`
	_, err := s.db.Exec(query, login, password)
	if err != nil {
		return fmt.Errorf("requers execution error: %w", err)
	}
	return nil
}

func (s *StoreDB) GetUser(login string) (int, error) {
	query := `SELECT id FROM users WHERE login=$1`
	var answer int
	err := s.db.QueryRow(query, login).Scan(&answer)
	if err != nil {
		return 0, err
	}
	return answer, nil
}
