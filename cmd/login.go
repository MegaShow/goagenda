package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:     "login",
	Aliases: []string{"l"},
	Short:   "Log in agenda",
	Long:    "Log in agenda with username and password.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
