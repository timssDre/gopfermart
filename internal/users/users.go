package users

import (
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
)

func (u *User) PasswordStringToHash() (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		// Обработка ошибки, если генерация хэша не удалась
		return "", err
	}
	hashPasswordString := base64.StdEncoding.EncodeToString(hashedPassword)
	return hashPasswordString, nil
}

func (u *User) PasswordHashToString(hashedPasswordString string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(hashedPasswordString)
}

func (u *User) ComparedPass(hash []byte, password string) error {
	bytesPassword := []byte(password)
	err := bcrypt.CompareHashAndPassword(hash, bytesPassword)
	if err != nil {
		return err
	}
	return nil
}
