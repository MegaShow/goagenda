package cmd

import (
	"github.com/MegaShow/goagenda/controller"
	"github.com/spf13/cobra"
)

var userRootCmd = &cobra.Command{
	Use:     "user",
	Aliases: []string{"u"},
	Short:   "User management",
	Long:    "User management",
}

var userDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "Delete your account",
	Long:    "Delete your account",
	Run:     wrapper(controller.GetUserCtrl().Delete),
}

var userListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List users",
	Long:    "List users",
	Run:     wrapper(controller.GetUserCtrl().List),
}

var userSetCmd = &cobra.Command{
	Use:     "set",
	Aliases: []string{"s"},
	Short:   "Set user's profile",
	Long:    "Set user's profile",
	Run:     wrapper(controller.GetUserCtrl().Set),
}

func init() {
	rootCmd.AddCommand(userRootCmd)
	userRootCmd.AddCommand(userDeleteCmd)
	userRootCmd.AddCommand(userListCmd)
	userRootCmd.AddCommand(userSetCmd)

	userListCmd.Flags().StringP("user", "u", "", "the username searched")

	userSetCmd.Flags().StringP("password", "p", "", "new password")
	userSetCmd.Flags().StringP("email", "e", "", "new email")
	userSetCmd.Flags().StringP("telephone", "t", "", "new telephone")
}