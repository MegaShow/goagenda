package service

import (
	"errors"
	"github.com/MegaShow/goagenda/lib/hash"
	"github.com/MegaShow/goagenda/lib/log"
	"github.com/MegaShow/goagenda/model"
)

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
		return errors.New("you are already logged in with user '"+status.Name+"', please logout first")
	}
	model.StatusModel.SetStatus(model.Status{ Name: user.Name })
	return nil
}
