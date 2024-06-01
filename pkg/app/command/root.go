package command

import (
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const VERSION = "sprune/v7.0.0"

var (
	homePath    string
	defaultHome = os.ExpandEnv("$HOME/.sprune")
	appName     = "Sprune"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   appName,
		Short: "sprune is meant to prune data base history from a cosmos application, avoiding needing to state sync every couple amount of weeks",
	}

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, _ []string) error {
		if err := initConfig(rootCmd, appConfig); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		return nil
	}

	rootCmd.PersistentFlags().StringVar(&homePath, flags.FlagHome, defaultHome, "set home directory")
	if err := viper.BindPFlag(flags.FlagHome, rootCmd.PersistentFlags().Lookup(flags.FlagHome)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rootCmd.AddCommand(
		configCmd(),
		pruneCmd(),
	)

	return rootCmd
}
