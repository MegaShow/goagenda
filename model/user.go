package model

import (
	"strings"
)

type UserDB struct {
	Data   []User
	Database
}

type User struct {
	Name      string `json:"user"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
	Salt      string `json:"salt"`
}

var UserModel = UserDB{  Database: Database{ schema: "User" } }

func (m *UserDB) GetUserByName(name string) User {
	initUserModel()
	for _, item := range m.Data {
		if strings.ToLower(item.Name) == strings.ToLower(name) {
			return item
		}
	}
	return User{}
}

func (m *UserDB) AddUser(user User) {
	initUserModel()
	m.isDirty = true
	m.Data = append(m.Data, user)
}

func initUserModel() {
	UserModel.initModel(&UserModel.Data)
}

func ReleaseUserModel() {
	UserModel.releaseModel(&UserModel.Data)
}
