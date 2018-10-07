package cmd

import (
	"github.com/MegaShow/goagenda/controller"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:     "register",
	Aliases: []string{"r", "reg"},
	Short:   "Register an account",
	Long: `Register an account with username and password necessarily,
and with email and telephone optionally.`,
	Run: controller.GetRegisterCtrl().Register,
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringP("user", "u", "", "username of your new account")
	registerCmd.Flags().StringP("password", "p", "", "password of your new account")
	registerCmd.Flags().StringP("email", "e", "", "email of your new account")
	registerCmd.Flags().StringP("telephone", "t", "", "telephone of your new account")

	controller.GetRegisterCtrl().Ctx.BindPFlags(registerCmd.Flags())
}
