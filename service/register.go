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
