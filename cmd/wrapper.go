package cmd

import (
	"github.com/MegaShow/goagenda/controller"
	"github.com/spf13/cobra"
)

func wrapper(fn func()) func(*cobra.Command, []string) {
	return controller.WrapperRun(fn)
}
