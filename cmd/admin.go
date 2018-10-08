package cmd

import (
	"github.com/MegaShow/goagenda/controller"
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:     "register",
	Aliases: []string{"r", "reg"},
	Short:   "Register an account",
	Long: `Register an account with username and password necessarily,
and with email and telephone optionally.`,
	PreRun: adminPreRun,
	Run:    controller.GetAdminCtrl().Register,
}

var loginCmd = &cobra.Command{
	Use:     "login",
	Aliases: []string{"l"},
	Short:   "Log in agenda",
	Long:    "Log in agenda with username and password.",
	PreRun:  adminPreRun,
	Run:     controller.GetAdminCtrl().Login,
}

var logoutCmd = &cobra.Command{
	Use:    "logout",
	Short:  "Log out agenda",
	Long:   "Log out agenda.",
	PreRun: adminPreRun,
	Run:    controller.GetAdminCtrl().Logout,
}

var statusCmd = &cobra.Command{
	Use:     "status",
	Aliases: []string{"s"},
	Short:   "Display username of the logged account",
	Long:    "Display username of the logged account",
	PreRun:  adminPreRun,
	Run:     controller.GetAdminCtrl().GetStatus,
}

func adminPreRun(cmd *cobra.Command, args []string) {
	controller.GetAdminCtrl().Ctx.BindPFlags(cmd.Flags())
}

func init() {
	rootCmd.AddCommand(registerCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)
	rootCmd.AddCommand(statusCmd)

	registerCmd.Flags().StringP("user", "u", "", "username of your new account")
	registerCmd.Flags().StringP("password", "p", "", "password of your new account")
	registerCmd.Flags().StringP("email", "e", "", "email of your new account")
	registerCmd.Flags().StringP("telephone", "t", "", "telephone of your new account")
	registerCmd.MarkFlagRequired("user")
	registerCmd.MarkFlagRequired("password")

	loginCmd.Flags().StringP("user", "u", "", "user of your account")
	loginCmd.Flags().StringP("password", "p", "", "password of your account")
	loginCmd.MarkFlagRequired("user")
	loginCmd.MarkFlagRequired("password")
}
