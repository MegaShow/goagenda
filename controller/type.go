package controller

import "github.com/spf13/viper"

type Controller struct {
	Ctx *viper.Viper
}

type LoginCtrl Controller
type RegisterCtrl Controller

func initController(ctrl *Controller) *Controller {
	if ctrl.Ctx == nil {
		ctrl.Ctx = viper.New()
	}
	return ctrl
}
