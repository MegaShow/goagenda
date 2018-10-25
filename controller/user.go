package controller

import (
	"fmt"
	"github.com/MegaShow/goagenda/lib/log"
)

type UserCtrl interface {
	UserDelete()
	UserList()
	UserSet()
}

func (c *Controller) UserDelete() {
	currentUser := c.Ctx.User.Get()
	if currentUser == "" {
		fmt.Println("you should login")
		return
	}

	c.Srv.User().DeleteUser(currentUser)
	c.Ctx.User.Set("")
	log.Info("Delete account successfully")
}

func (c *Controller) UserList() {
	// TODO
}

func (c *Controller) UserSet() {
	password, setP := c.Ctx.GetSecretString("password")
	email, setE := c.Ctx.GetString("email")
	telephone, setT := c.Ctx.GetString("telephone")

	if setP {
		verifyNonNilPassword(password)
	}
	verifyEmail(email)
	verifyTelephone(telephone)

	log.Verbose("check status")
	currentUser := c.Ctx.User.Get()
	if currentUser == "" {
		fmt.Println("you should login")
		return
	}

	if !setP && !setE && !setT {
		fmt.Println("set nothing")
		return
	}
	err := c.Srv.User().SetUser(currentUser, password, setP, email, setE, telephone, setT)
	if err != nil {
		log.Error(err.Error())
	}
	log.Info("set user successfully")
}

func GetUserCtrl() UserCtrl {
	return &ctrl
}
