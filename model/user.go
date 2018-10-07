package model

type UserModel struct {
	DB UserDB
}

type UserDB struct {
	Data []User
	ModelDB
}

type User struct {
	User      string `json:"user"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
	Salt      string `json:"salt"`
}
