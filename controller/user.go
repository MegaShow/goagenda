package controller

import (
	"fmt"
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
	fmt.Println("password:", c.Ctx.GetString("password"))
	fmt.Println("email:", c.Ctx.GetString("email"))
	fmt.Println("telephone:", c.Ctx.GetString("telephone"))
	// TODO
}

func GetUserCtrl() UserCtrl {
	return &ctrl
}
