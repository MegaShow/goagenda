package service

import (
	"errors"
	"github.com/MegaShow/goagenda/lib/hash"
	"github.com/MegaShow/goagenda/lib/log"
	"github.com/MegaShow/goagenda/model"
)

type AdminService interface {
	GetCurrentUserName() string
	Login(name, password string) error
	Register(name, password, email, telephone string) error
	SetCurrentUserName(name string) error
}

func (s *Service) Register(name, password, email, telephone string) error {
	log.Verbose("check if username exists")
	user := s.DB.User().GetUserByName(name)
	if user.Name != "" {
		return errors.New("user name already exists")
	}
	log.Verbose("add new user into database")
	password, salt := hash.Encrypt(password)
	s.DB.User().AddUser(model.User{
		Name:      name,
		Email:     email,
		Telephone: telephone,
		Password:  password,
		Salt:      salt,
	})
	return nil
}

func (s *Service) Login(name, password string) error {
	log.Verbose("check if username and password is correct")
	user := s.DB.User().GetUserByName(name)
	if user.Name == "" {
		return errors.New("invalid username or password")
	}
	checkPassword := hash.EncryptWithSalt(password, user.Salt)
	if checkPassword != user.Password {
		return errors.New("invalid user name or password")
	}
	log.Verbose("check status")
	status := s.DB.Status().GetStatus()
	if status.Name == user.Name {
		return errors.New("you are already logged in with this account")
	} else if status.Name != "" {
		return errors.New("you are already logged in with user '" + status.Name + "', please logout first")
	}
	return nil
}

func (s *Service) GetCurrentUserName() string {
	return s.DB.Status().GetStatus().Name
}

func (s *Service) SetCurrentUserName(name string) error {
	s.DB.Status().SetStatus(model.Status{Name: name})
	return nil
}

func (s *Manager) Admin() AdminService {
	return s.GetService()
}
