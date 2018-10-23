package model

import (
	"strings"
)

type UserModel interface {
	AddUser(user User)
	GetUserByName(name string) User
	SetUser(name, password, salt string, setPassword bool, email string, setEmail bool, telephone string, setTel bool)
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

func (m *UserDB) SetUser(name, password, salt string, setPassword bool, email string, setEmail bool, telephone string, setTel bool) {
	m.isDirty = true
	for index, item := range m.Data {
		if strings.ToLower(item.Name) == strings.ToLower(name) {
			if setPassword {
				m.Data[index].Password = password
				m.Data[index].Salt = salt
			}
			if setEmail {
				m.Data[index].Email = email
			}
			if setTel {
				m.Data[index].Telephone = telephone
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
