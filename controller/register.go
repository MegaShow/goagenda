package controller

import (
	"fmt"
	"github.com/MegaShow/goagenda/lib/log"
	"github.com/MegaShow/goagenda/service"
	"github.com/spf13/cobra"
)

var registerCtrl RegisterCtrl

func (c *RegisterCtrl) Register(cmd *cobra.Command, args []string) {
	log.Info("register called")
	fmt.Println(c.Ctx.GetString("user"), c.Ctx.GetString("password"), c.Ctx.GetString("email"), c.Ctx.GetString("telephone"))
	service.Register("", "", "", "")
}

func GetRegisterCtrl() *RegisterCtrl {
	return (*RegisterCtrl)(initController((*Controller)(&registerCtrl)))
}
