package service

import "github.com/MegaShow/goagenda/model"

type Service struct {
	DB model.Manager
}

type Manager Service

func (s *Manager) GetService() *Service {
	return (*Service)(s)
}
