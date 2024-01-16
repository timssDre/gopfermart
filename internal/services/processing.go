package services

import (
	"fmt"
	"mBoxMini/internal/users"
)

type Repo interface {
	CreateUser(id, login, password string) error
	GetUser(login string) (int, error)
}

type BoxService struct {
	db Repo
}

func NewBoxService(db Repo) *BoxService {
	return &BoxService{
		db,
	}
}

func (b *BoxService) GetUser(login string) (int, error) {
	return b.db.GetUser(login)
}

func (b *BoxService) CreateUser(user *users.User) error {
	hashedPassword, err := user.PasswordStringToHash()
	if err != nil {
		return err
	}
	err = b.db.CreateUser(user.ID, user.Login, string(hashedPassword))
	if err != nil {
		return fmt.Errorf("failed registration user in db  %w", err)
	}
	return nil
}
