package controller

import (
	"github.com/MegaShow/goagenda/lib/log"
	"github.com/MegaShow/goagenda/model"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Controller struct {
	Ctx *viper.Viper
}

func CtrlRelease(cmd *cobra.Command, args []string) {
	log.Release()
	model.ReleaseUserModel()
	model.ReleaseMeetingModel()
	model.ReleaseStatusModel()
}

func initController(ctrl *Controller) *Controller {
	if ctrl.Ctx == nil {
		ctrl.Ctx = viper.New()
	}
	return ctrl
}
