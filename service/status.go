package service

import (
	"github.com/MegaShow/goagenda/model"
)

func GetCurrentUserName() string  {
	return model.StatusModel.GetStatus().Name
}

func SetCurrentUserName(name string) {
	model.StatusModel.SetStatus(model.Status{ Name: name })
}
