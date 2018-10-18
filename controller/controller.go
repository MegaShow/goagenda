package controller

import (
	"github.com/MegaShow/goagenda/lib/log"
	"github.com/MegaShow/goagenda/model"
	"github.com/MegaShow/goagenda/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ctrl Controller

type Controller struct {
	Args    []string
	Cmd     *cobra.Command
	Ctx     *viper.Viper
	Srv		service.Manager
}

func CtrlRelease(cmd *cobra.Command, args []string) {
	log.Release()
	model.ReleaseUserModel()
	model.ReleaseMeetingModel()
	model.ReleaseStatusModel()
}

func WrapperRun(fn func()) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		cmdStr := cmd.Name()
		cmd.VisitParents(func(pcmd *cobra.Command) { cmdStr = pcmd.Name() + "." + cmdStr })
		log.SetCommand(cmdStr)
		ctrl.Args = args
		ctrl.Cmd = cmd
		ctrl.Ctx = viper.New()
		ctrl.Ctx.BindPFlags(cmd.Flags())
		fn()
	}
}
