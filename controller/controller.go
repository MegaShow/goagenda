package controller

import (
	"github.com/MegaShow/goagenda/lib/log"
	"github.com/MegaShow/goagenda/model"
	"github.com/MegaShow/goagenda/service"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var ctrl Controller

type Controller struct {
	Args []string
	Cmd  *cobra.Command
	Ctx  Ctx
	Srv  service.Manager
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
		log.Init()
		log.SetCommand(cmdStr)
		if len(args) != 0 {
			log.AddParams("args", args)
		}
		ctrl.Args = args
		ctrl.Cmd = cmd
		ctrl.Ctx.User = &user{
			get: func() string {
				name := ctrl.Srv.Admin().GetCurrentUserName()
				log.SetUser(name)
				return name
			},
			set: func(name string) error {
				err := ctrl.Srv.Admin().SetCurrentUserName(name)
				if err == nil && name != "" {
					log.SetUser(name)
				}
				return err
			},
		}
		ctrl.Ctx.User.Get()
		ctrl.Ctx.Value = viper.New()
		ctrl.Ctx.Value.BindPFlags(cmd.Flags())
		ctrl.Ctx.Visit = make(map[string]bool)
		cmd.Flags().Visit(func(flag *pflag.Flag) { ctrl.Ctx.Visit[flag.Name] = true })
		fn()
	}
}
