package services

import (
	"fmt"
	"mBoxMini/internal/users"
)

type Repo interface {
	CreateUser(login, password string) error
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
	err := b.db.CreateUser(user.Login, user.Password)
	if err != nil {
		return fmt.Errorf("failed registration user in db  %w", err)
	}
	return nil
}
