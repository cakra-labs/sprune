package app

import (
	"github.com/cakra-labs/sprune/v7/pkg/app/command"
	"github.com/spf13/cobra"
)

func newCLI() *cobra.Command {
	cobra.EnableCommandSorting = false

	rootCmd := command.NewRootCmd()
	rootCmd.SilenceUsage = true
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	return rootCmd
}
