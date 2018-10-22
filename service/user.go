package service

import (
	"errors"
	"github.com/MegaShow/goagenda/lib/hash"
	"github.com/MegaShow/goagenda/lib/log"
)

type UserService interface {
	Set(password string, setPassword bool, email string, setEmail bool, telephone string, setTel bool) error
}

func (s *Service) Set(password string, setPassword bool, email string, setEmail bool, telephone string, setTel bool) error {
	log.Verbose("check status")
	status := s.DB.Status().GetStatus()
	if status.Name == "" {
		return errors.New("you are not logged")
	}

	log.Verbose("set logged user")
	salt := ""
	if setPassword {
		password, salt = hash.Encrypt(password)
	}
	s.DB.User().SetUser(password, salt, setPassword, email, setEmail, telephone, setTel)
	return nil
}

func (s *Manager) User() UserService {
	return s.GetService()
}
