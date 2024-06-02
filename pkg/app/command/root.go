package command

import (
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	homePath    string
	defaultHome = os.ExpandEnv("$HOME/.sprune")
	appName     = "sprune"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   appName,
		Short: "Sprune prunes database history in a Cosmos application, eliminating the need for frequent state synchronizations",
	}

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, _ []string) error {
		if err := initConfig(appConfig); err != nil {
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
