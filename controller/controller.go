package controller

import (
	"github.com/MegaShow/goagenda/lib/log"
	"github.com/MegaShow/goagenda/model"
	"github.com/MegaShow/goagenda/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Controller struct {
	Ctx *viper.Viper
	User Status
}

type UserCtrl Controller
type LoginCtrl Controller
type RegisterCtrl Controller
type StatusCtrl Controller

type Status interface {
	GetUser() string
	SetUser(string)
}

func (c *Controller) GetUser() string {
	return service.GetCurrentUserName()
}

func (c *Controller) SetUser(user string) {
	service.SetCurrentUserName(user)
}

func CtrlRelease(cmd *cobra.Command, args []string) {
	log.Release()
	model.ReleaseUserModel()
	model.ReleaseStatusModel()
}

func initController(ctrl *Controller) *Controller {
	if ctrl.Ctx == nil {
		ctrl.Ctx = viper.New()
	}
	return ctrl
}
