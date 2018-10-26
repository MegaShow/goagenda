package controller

import (
	"fmt"
	"github.com/MegaShow/goagenda/lib/log"
)

type AdminCtrl interface {
	GetStatus()
	Login()
	Logout()
	Register()
}

func (c *Controller) Register() {
	user, _ := c.Ctx.GetString("user")
	password, _ := c.Ctx.GetSecretString("password")
	email, _ := c.Ctx.GetString("email")
	telephone, _ := c.Ctx.GetString("telephone")

	verifyNonNilUser(user)
	verifyNonNilPassword(password)
	verifyEmail(email)
	verifyTelephone(telephone)

	err := c.Srv.Admin().Register(user, password, email, telephone)
	if err != nil {
		log.Error(err.Error())
	}
	log.Info("register account successfully")
}

func (c *Controller) Login() {
	user, _ := c.Ctx.GetString("user")
	password, _ := c.Ctx.GetSecretString("password")

	verifyNonNilUser(user)
	verifyNonNilPassword(password)

	log.Verbose("check status")
	currentUser := c.Ctx.User.Get()
	if currentUser == user {
		log.Error("you are already logged in with this account")
	} else if currentUser != "" {
		log.Error("you are already logged in with user '" + currentUser + "', please logout first")
	}

	err := c.Srv.Admin().Login(user, password)
	if err != nil {
		log.Error(err.Error())
	}
	err = c.Ctx.User.Set(user)
	if err != nil {
		log.Error(err.Error())
	}
	log.Info("login successfully")
}

func (c *Controller) Logout() {
	currentUser := c.Ctx.User.Get()
	if currentUser == "" {
		fmt.Println("not logged user")
		return
	}
	c.Ctx.User.Set("")
	log.Info("user '" + currentUser + "' logged out")
}

func (c *Controller) GetStatus() {
	currentUser := c.Ctx.User.Get()
	if currentUser == "" {
		fmt.Println("not logged user")
	} else {
		fmt.Println("user '" + currentUser + "' logged in")
	}
}

func GetAdminCtrl() AdminCtrl {
	return &ctrl
}
