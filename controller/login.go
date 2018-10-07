package controller

import (
	"fmt"
	"github.com/MegaShow/goagenda/lib/log"
	"github.com/spf13/cobra"
)

var loginCtrl LoginCtrl

func (c *LoginCtrl) Login(cmd *cobra.Command, args []string) {
	log.Info("login called")
	fmt.Println(c.Ctx.GetString("user"), c.Ctx.GetString("password"))
}

func GetLoginCtrl() *LoginCtrl {
	return (*LoginCtrl)(initController((*Controller)(&loginCtrl)))
}
