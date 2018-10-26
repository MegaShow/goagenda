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
	Args:    cobra.NoArgs,
	Run:     wrapper(controller.GetUserCtrl().UserDelete),
}

var userListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List users",
	Long:    "List users",
	Args:    cobra.NoArgs,
	Run:     wrapper(controller.GetUserCtrl().UserList),
}

var userSetCmd = &cobra.Command{
	Use:     "set",
	Aliases: []string{"s"},
	Short:   "Set user's profile",
	Long:    "Set user's profile",
	Args:    cobra.NoArgs,
	Run:     wrapper(controller.GetUserCtrl().UserSet),
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

	userDeleteCmd.Flags().StringP("username", "u", "", "user name")
	userDeleteCmd.Flags().StringP("password", "p", "", "the user password")
	userDeleteCmd.MarkFlagRequired("username")
	userDeleteCmd.MarkFlagRequired("password")
}
