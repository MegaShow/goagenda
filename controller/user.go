package controller

import (
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
	userName := c.Srv.Admin().GetCurrentUserName()
	if userName == "" {
		log.Show("you should login")
		return
	}
	log.SetUser(userName)
	c.Srv.User().DeleteUser(userName)
	log.Info("Delete account successfully")
	c.Srv.Admin().SetCurrentUserName("")
}

func (c *Controller) List() {
	fmt.Println("user:", c.Ctx.GetString("user"))
	// TODO
}

func (c *Controller) Set() {
	fmt.Println("password:", c.Ctx.GetString("password"))
	fmt.Println("email:", c.Ctx.GetString("email"))
	fmt.Println("telephone:", c.Ctx.GetString("telephone"))
	// TODO
}

func GetUserCtrl() UserCtrl {
	return &ctrl
}
