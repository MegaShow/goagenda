package controller

import (
	"errors"
	"fmt"
	"github.com/MegaShow/goagenda/lib/log"
)

type UserCtrl interface {
	Delete()
	List()
	Set()
}

func (c *Controller) Delete() {
	// TODO
}

func (c *Controller) List() {
	fmt.Println("user:", c.Ctx.GetString("user"))
	// TODO
}

func (c *Controller) Set() {
	password := c.Ctx.GetString("password")
	email := c.Ctx.GetString("email")
	telephone := c.Ctx.GetString("telephone")
	_, setP := c.Visit["password"]
	_, setE := c.Visit["email"]
	_, setT := c.Visit["telephone"]

	if setP && password == "" {
		err := errors.New("password empty")
		log.Error(err.Error())
	}
	verifyPassword(password)
	verifyEmail(email)
	verifyTelephone(telephone)

	err := c.Srv.User().Set(password, setP, email, setE, telephone, setT)
	if err != nil {
		log.Error(err.Error())
	}
	log.SetUser(c.Srv.Admin().GetCurrentUserName())
	log.Info("set user successfully")
}

func GetUserCtrl() UserCtrl {
	return &ctrl
}
