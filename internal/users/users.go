package users

import (
	"golang.org/x/crypto/bcrypt"
)

func NewUser(id string, new bool, token string) *User {
	return &User{
		ID:    id,
		New:   new,
		Token: token,
	}
}
func (u *User) PasswordStringToHash() ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
}
