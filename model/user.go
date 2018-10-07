package model

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"os"
	"strings"
)

type UserDB struct {
	Data   []User
	isInit bool
	isDirty bool
	path string
	file string
}

type User struct {
	Name      string `json:"user"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
	Salt      string `json:"salt"`
}

var UserModel UserDB

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
	if UserModel.isInit == false {
		UserModel.path = viper.GetString("Database.Path")
		UserModel.file = viper.GetString("Database.UserFile")
		if _, err := os.Stat(UserModel.path); err != nil {
			err := os.MkdirAll(UserModel.path, 0777)
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
		}
		f, err := os.OpenFile(UserModel.path + string(os.PathSeparator) + UserModel.file, os.O_CREATE|os.O_RDONLY, 0666)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		defer f.Close()
		decoder := json.NewDecoder(f)
		err = decoder.Decode(&UserModel.Data)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			os.Exit(2)
		}
		UserModel.isInit = true
	}
}

func ReleaseUserModel() {
	if UserModel.isDirty == true {
		f, err := os.OpenFile(UserModel.path + string(os.PathSeparator) + UserModel.file, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		defer f.Close()
		encoder := json.NewEncoder(f)
		err = encoder.Encode(UserModel.Data)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		UserModel.isDirty = false
	}
}
