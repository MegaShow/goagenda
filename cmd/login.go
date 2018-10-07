package cmd

import (
	"github.com/MegaShow/goagenda/controller"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:     "login",
	Aliases: []string{"l"},
	Short:   "Log in agenda",
	Long:    "Log in agenda with username and password.",
	Run:     controller.GetLoginCtrl().Login,
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("user", "u", "", "user of your account")
	loginCmd.Flags().StringP("password", "p", "", "password of your account")

	controller.GetLoginCtrl().Ctx.BindPFlags(loginCmd.Flags())
}
