package controller

import (
	"github.com/MegaShow/goagenda/lib/log"
)

type AdminCtrl interface {
	GetStatus()
	Login()
	Logout()
	Register()
}

func (c *Controller) Register() {
	user := c.Ctx.GetString("user")
	password := c.Ctx.GetString("password")
	email := c.Ctx.GetString("email")
	telephone := c.Ctx.GetString("telephone")

	verifyUser(user)
	verifyPassword(password)
	verifyEmail(email)
	verifyTelephone(telephone)

	log.SetUser(user)
	log.AddParams("email", email)
	log.AddParams("telephone", telephone)
	err := c.Srv.Admin().Register(user, password, email, telephone)
	if err != nil {
		log.Error(err.Error())
	}
	log.Info("register account successfully")
}

func (c *Controller) Login() {
	user := c.Ctx.GetString("user")
	password := c.Ctx.GetString("password")

	verifyUser(user)
	verifyPassword(password)

	log.SetUser(user)
	err := c.Srv.Admin().Login(user, password)
	if err != nil {
		log.Error(err.Error())
	}
	c.Srv.Admin().SetCurrentUserName(user)
	log.Info("login successfully")
}

func (c *Controller) Logout() {
	user := c.Srv.Admin().GetCurrentUserName()
	if user == "" {
		log.Show("not logged user")
		return
	}
	log.SetUser(user)
	c.Srv.Admin().SetCurrentUserName("")
	log.Info("user '" + user + "' logged out")
}

func (c *Controller) GetStatus() {
	user := c.Srv.Admin().GetCurrentUserName()
	if user == "" {
		log.Show("not logged user")
	} else {
		log.Show("user '" + user + "' logged in")
	}
}

func GetAdminCtrl() AdminCtrl {
	return &ctrl
}
