package controller

import (
	"github.com/MegaShow/goagenda/lib/log"
	"github.com/spf13/cobra"
)

var statusCtrl StatusCtrl

func (c *StatusCtrl) GetStatus(cmd *cobra.Command, args []string) {
	user := c.User.GetUser()
	if user == "" {
		log.Show("not logged user")
	} else {
		log.Show("user '" + user + "' logged in")
	}
}

func GetStatusCtrl() *StatusCtrl {
	// because not used of Ctx, it doesn't need to init.
	// return (*StatusCtrl)(initController((*Controller)(&statusCtrl)))
	return &statusCtrl
}
