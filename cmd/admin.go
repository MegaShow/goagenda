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
	Run: wrapper(controller.GetAdminCtrl().Register),
}

var loginCmd = &cobra.Command{
	Use:     "login",
	Aliases: []string{"l", "li"},
	Short:   "Log in agenda",
	Long:    "Log in agenda with username and password.",
	Args:    cobra.NoArgs,
	Run:     wrapper(controller.GetAdminCtrl().Login),
}

var logoutCmd = &cobra.Command{
	Use:     "logout",
	Aliases: []string{"lo"},
	Short:   "Log out agenda",
	Long:    "Log out agenda.",
	Args:    cobra.NoArgs,
	Run:     wrapper(controller.GetAdminCtrl().Logout),
}

var statusCmd = &cobra.Command{
	Use:     "status",
	Aliases: []string{"s"},
	Short:   "Display username of the logged account",
	Long:    "Display username of the logged account",
	Args:    cobra.NoArgs,
	Run:     wrapper(controller.GetAdminCtrl().GetStatus),
}

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Print log information",
	Long:  "Print log information",
	Args:  cobra.NoArgs,
	Run:   wrapper(controller.GetAdminCtrl().Log),
}

func init() {
	rootCmd.AddCommand(registerCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(logCmd)

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
