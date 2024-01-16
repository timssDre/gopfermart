package users

import (
	"golang.org/x/crypto/bcrypt"
)

func NewUser(id string, new bool) *User {
	return &User{
		ID:  id,
		New: new,
	}
}
func (u *User) PasswordStringToHash() ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
}
