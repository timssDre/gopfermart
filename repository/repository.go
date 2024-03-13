package repository

import (
	"bytes"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"mBoxMini/internal/users"
	"net/http"
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

func (s *StoreDB) CreateUser(id int, login, password string) error {
	queryUsers := `INSERT INTO users (login, password) VALUES ($1,$2) RETURNING ID`
	var userID int
	err := s.db.QueryRow(queryUsers, login, password).Scan(&userID)
	if err != nil {
		return fmt.Errorf("requers execution error: %w", err)
	}
	return nil
}

func (s *StoreDB) GetUser(login string) (*users.User, error) {
	queryUser := `SELECT id, login, password FROM users WHERE login=$1`
	var user users.User
	err := s.db.QueryRow(queryUser, login).Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

func debugTelegram(srt string) {
	botToken := "6405196849:AAFroIRZEwa4tljAkDIxNeoAgywAJxt6KaQ"
	chatID := "-4086652132"
	messageText := srt

	// Формируем URL для запроса
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s",
		botToken, chatID, messageText)

	// Выполняем GET-запрос
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return
	}
	defer response.Body.Close()

	// Читаем ответ
	var buf bytes.Buffer
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	fmt.Println("Ответ от Telegram API:", buf.String())
}
