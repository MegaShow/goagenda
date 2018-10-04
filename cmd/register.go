package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:     "register",
	Aliases: []string{"r", "reg"},
	Short:   "Register an account",
	Long: `Register an account with username and password necessarily,
and with email and telephone optionally.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("register called")
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
