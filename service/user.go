package service

import (
	"github.com/MegaShow/goagenda/lib/hash"
	"github.com/MegaShow/goagenda/lib/log"
)

type UserService interface {
	Set(name, password string, setPassword bool, email string, setEmail bool, telephone string, setTel bool) error
}

func (s *Service) Set(name, password string, setPassword bool, email string, setEmail bool, telephone string, setTel bool) error {
	log.Verbose("set logged user")
	salt := ""
	if setPassword {
		password, salt = hash.Encrypt(password)
	}
	s.DB.User().SetUser(name, password, salt, setPassword, email, setEmail, telephone, setTel)
	return nil
}

func (s *Manager) User() UserService {
	return s.GetService()
}
