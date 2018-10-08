package service

import (
	"errors"
	"github.com/MegaShow/goagenda/lib/hash"
	"github.com/MegaShow/goagenda/lib/log"
	"github.com/MegaShow/goagenda/model"
)

func Register(name, password, email, telephone string) error {
	log.Verbose("check if username exists")
	user := model.UserModel.GetUserByName(name)
	if user.Name != "" {
		return errors.New("user name already exists")
	}
	log.Verbose("add new user into database")
	password, salt := hash.Encrypt(password)
	model.UserModel.AddUser(model.User{
		Name:      name,
		Email:     email,
		Telephone: telephone,
		Password:  password,
		Salt:      salt,
	})
	return nil
}

func Login(name, password string) error {
	log.Verbose("check if username and password is correct")
	user := model.UserModel.GetUserByName(name)
	if user.Name == "" {
		return errors.New("invalid username or password")
	}
	checkPassword := hash.EncryptWithSalt(password, user.Salt)
	if checkPassword != user.Password {
		return errors.New("invalid user name or password")
	}
	log.Verbose("check status")
	status := model.StatusModel.GetStatus()
	if status.Name == user.Name {
		return errors.New("you are already logged in with this account")
	} else if status.Name != "" {
		return errors.New("you are already logged in with user '" + status.Name + "', please logout first")
	}
	return nil
}

func GetCurrentUserName() string {
	return model.StatusModel.GetStatus().Name
}

func SetCurrentUserName(name string) {
	model.StatusModel.SetStatus(model.Status{Name: name})
}
