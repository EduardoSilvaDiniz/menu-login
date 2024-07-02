package entities

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUser() *User {
	user := User{}
	return &user
}
