package controller

import (
	"github.com/MegaShow/goagenda/lib/log"
	"github.com/MegaShow/goagenda/lib/verify"
	"github.com/MegaShow/goagenda/service"
	"github.com/spf13/cobra"
)

var registerCtrl RegisterCtrl

func (c *RegisterCtrl) Register(cmd *cobra.Command, args []string) {
	user := c.Ctx.GetString("user")
	password := c.Ctx.GetString("password")
	email := c.Ctx.GetString("email")
	telephone := c.Ctx.GetString("telephone")

	verify.AssertNil(user, "user name required")
	verify.AssertNil(password, "password required")
	verifyUser(user)
	verifyPassword(password)
	verifyEmail(email)
	verifyTelephone(telephone)

	log.SetUser(user)
	err := service.Register(user, password, email, telephone)
	if err != nil {
		log.Error(err.Error())
	}
	log.Info("register account successfully")
}

func GetRegisterCtrl() *RegisterCtrl {
	return (*RegisterCtrl)(initController((*Controller)(&registerCtrl)))
}
