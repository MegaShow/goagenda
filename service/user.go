package service

import (
	"errors"

	"github.com/MegaShow/goagenda/lib/hash"
	"github.com/MegaShow/goagenda/lib/log"
)

type UserService interface {
	DeleteUser(name string, password string) error
	GetAllUsers() string
	GetUserDetail(name string) (string, error)
	SetUser(name, password string, setPassword bool, email string, setEmail bool, telephone string, setTel bool) error
}

func (s *Service) SetUser(name, password string, setPassword bool, email string, setEmail bool, telephone string, setTel bool) error {
	log.Verbose("set logged user")
	salt := ""
	if setPassword {
		password, salt = hash.Encrypt(password)
	}
	s.DB.User().SetUser(name, password, salt, setPassword, email, setEmail, telephone, setTel)
	return nil
}

func (s *Service) DeleteUser(name string, password string) error {
	user := s.DB.User().GetUserByName(name)
	if user.Name == "" {
		return errors.New("invalid username or password")
	}
	checkPassword := hash.EncryptWithSalt(password, user.Salt)
	if checkPassword != user.Password {
		return errors.New("invalid user name or password")
	}
	s.DB.User().DeleteUser(name)
	return nil
}

func (s *Manager) User() UserService {
	return s.GetService()
}

func (s *Service) GetAllUsers() string {
	users := s.DB.User().GetAllUser()
	output := "\n  ---------------------------\n"
	for i := 0; i < len(users); i++ {
		user := "  Name: " + users[i].Name + "\n" + "  Email: " + users[i].Email + "\n" + "  Tel: " + users[i].Telephone + "\n"
		user += "  ---------------------------\n"
		output += user
	}
	return output
}

func (s *Service) GetUserDetail(name string) (string, error) {
	user := s.DB.User().GetUserByName(name)
	if user.Name == "" {
		return "", errors.New("can't find this user")
	}
	return "\n  ---------------------------\n" + "  Name: " + user.Name + "\n" + "  Email: " + user.Email + "\n" + "  Tel: " + user.Telephone + "\n" + "  ---------------------------\n", nil
}
