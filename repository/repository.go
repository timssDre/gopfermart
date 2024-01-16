package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"time"
)

type StoreDB struct {
	db *sql.DB
}

func InitDatabase(DatabasePath string) (*StoreDB, error) {
	db, err := sql.Open("pgx", DatabasePath)
	if err != nil {
		return nil, fmt.Errorf("error opening db: %w", err)
	}

	storeDB := new(StoreDB)
	storeDB.db = db

	err = createTableUsers(db)
	if err != nil {
		return nil, fmt.Errorf("error creae table db: %w", err)
	}

	return storeDB, nil
}

func (s *StoreDB) CreateUser(id, login, password string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func() {
		if err != nil {
			if err = tx.Rollback(); err != nil {
				err = fmt.Errorf("failed to rollback to transaction %w", err)
			}
			return
		}
		if err = tx.Commit(); err != nil {
			err = fmt.Errorf("failed to commit transaction %w", err)
		}
	}()
	queryUsers := `INSERT INTO users (login, password) VALUES ($1,$2) RETURNING ID`
	var userID int
	err = s.db.QueryRow(queryUsers, login, password).Scan(&userID)
	if err != nil {
		return fmt.Errorf("requers execution error: %w", err)
	}

	query := `INSERT INTO tokens (user_id, token, expiration_time) VALUES ($1,$2, $3)`
	expirationTime := time.Now().Add(time.Hour * 24)
	_, err = s.db.Exec(query, userID, id, expirationTime)
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
