package model

import (
	"strings"
)

type UserModel interface {
	AddUser(user User)
	GetUserByName(name string) User
	SetUser(password string, salt string, email string, setEmail bool, telephone string, setTel bool)
}

type UserDB struct {
	Data []User
	Database
}

type User struct {
	Name      string `json:"user"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
	Salt      string `json:"salt"`
}

var userDB = UserDB{Database: Database{schema: "User"}}

func (m *UserDB) GetUserByName(name string) User {
	for _, item := range m.Data {
		if strings.ToLower(item.Name) == strings.ToLower(name) {
			return item
		}
	}
	return User{}
}

func (m *UserDB) AddUser(user User) {
	m.isDirty = true
	m.Data = append(m.Data, user)
}

func (m *UserDB) SetUser(password string, salt string, email string, setEmail bool, telephone string, setTel bool) {
	m.isDirty = true
	name := statusDB.GetStatus().Name
	for _, item := range m.Data {
		if strings.ToLower(item.Name) == strings.ToLower(name) {
			item.Password = password
			item.Salt = salt
			if setEmail {
				item.Email = email
			}
			if setTel {
				item.Telephone = telephone
			}
			break
		}
	}
}

func ReleaseUserModel() {
	userDB.releaseModel(&userDB.Data)
}

func (m *Manager) User() UserModel {
	if userDB.isInit == false {
		userDB.initModel(&userDB.Data)
	}
	return &userDB
}
