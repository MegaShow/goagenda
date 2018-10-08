package cmd

import (
	"github.com/MegaShow/goagenda/controller"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:     "status",
	Aliases: []string{"s"},
	Short:   "Display username of the logged account",
	Long:    "Display username of the logged account",
	Run:     controller.GetStatusCtrl().GetStatus,
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
