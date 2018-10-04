package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any sub-commands
var rootCmd = &cobra.Command{
	Use:   "agenda",
	Short: "An meeting management system",
	Long: `Agenda is a meetings management system.
This application is a perfect and essential tool
to be well organized in your work.`,
	Run: func(cmd *cobra.Command, args []string) { cmd.Usage() },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
