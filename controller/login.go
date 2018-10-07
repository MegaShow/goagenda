package controller

import (
	"github.com/MegaShow/goagenda/lib/log"
	"github.com/MegaShow/goagenda/lib/verify"
	"github.com/MegaShow/goagenda/service"
	"github.com/spf13/cobra"
)

var loginCtrl LoginCtrl

func (c *LoginCtrl) Login(cmd *cobra.Command, args []string) {
	user := c.Ctx.GetString("user")
	password := c.Ctx.GetString("password")

	verify.AssertNil(user, "user name required")
	verify.AssertNil(password, "password required")
	verifyUser(user)
	verifyPassword(password)

	log.SetUser(user)
	err := service.Login(user, password)
	if err != nil {
		log.Error(err.Error())
	}
	log.Info("login successfully")
}

func GetLoginCtrl() *LoginCtrl {
	return (*LoginCtrl)(initController((*Controller)(&loginCtrl)))
}
