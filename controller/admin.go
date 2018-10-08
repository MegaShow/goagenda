package controller

import (
	"github.com/MegaShow/goagenda/lib/log"
	"github.com/MegaShow/goagenda/service"
	"github.com/spf13/cobra"
)

type AdminCtrl Controller

var adminCtrl AdminCtrl

func (c *AdminCtrl) Register(cmd *cobra.Command, args []string) {
	user := c.Ctx.GetString("user")
	password := c.Ctx.GetString("password")
	email := c.Ctx.GetString("email")
	telephone := c.Ctx.GetString("telephone")

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

func (c *AdminCtrl) Login(cmd *cobra.Command, args []string) {
	user := c.Ctx.GetString("user")
	password := c.Ctx.GetString("password")

	verifyUser(user)
	verifyPassword(password)

	log.SetUser(user)
	err := service.Login(user, password)
	if err != nil {
		log.Error(err.Error())
	}
	service.SetCurrentUserName(user)
	log.Info("login successfully")
}

func (c *AdminCtrl) Logout(cmd *cobra.Command, args []string) {
	user := service.GetCurrentUserName()
	if user == "" {
		log.Show("not logged user")
		return
	}
	log.SetUser(user)
	service.SetCurrentUserName("")
	log.Info("user '" + user + "' logged out")
}

func (c *AdminCtrl) GetStatus(cmd *cobra.Command, args []string) {
	user := service.GetCurrentUserName()
	if user == "" {
		log.Show("not logged user")
	} else {
		log.Show("user '" + user + "' logged in")
	}
}

func GetAdminCtrl() *AdminCtrl {
	return (*AdminCtrl)(initController((*Controller)(&adminCtrl)))
}
