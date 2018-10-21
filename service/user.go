package service

type UserService interface {
	DeleteUser(name string)
}

func (s *Service) DeleteUser(name string) {
	s.DB.User().DeleteUser(name)

}

func (s *Manager) User() UserService {
	return s.GetService()
}
