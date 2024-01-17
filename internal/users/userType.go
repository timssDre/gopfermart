package users

type User struct {
	ID       string `json:"id"`
	New      bool   `json:"new"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Token    string
}

type UserInterface interface {
	PasswordStringToHash() ([]byte, error)
}
