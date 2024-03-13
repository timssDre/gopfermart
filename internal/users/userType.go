package users

type User struct {
	ID       int    `json:"id"`
	New      bool   `json:"new"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserInterface interface {
	PasswordStringToHash() ([]byte, error)
}
