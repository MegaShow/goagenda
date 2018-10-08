package controller

import (
	"fmt"
	"github.com/spf13/cobra"
)

type UserCtrl Controller

var userCtrl UserCtrl

func (c *UserCtrl) Delete(cmd *cobra.Command, args []string) {

}

func (c *UserCtrl) List(cmd *cobra.Command, args []string) {
	fmt.Println("user:", c.Ctx.GetString("user"))
}

func (c *UserCtrl) Set(cmd *cobra.Command, args []string) {
	fmt.Println("password:", c.Ctx.GetString("password"))
	fmt.Println("email:", c.Ctx.GetString("email"))
	fmt.Println("telephone:", c.Ctx.GetString("telephone"))
}

func GetUserCtrl() *UserCtrl {
	return (*UserCtrl)(initController((*Controller)(&userCtrl)))
}
