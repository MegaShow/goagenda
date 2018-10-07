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
	// TODO
	return nil
}
