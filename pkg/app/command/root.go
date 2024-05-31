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
	blocks      uint64
	versions    uint64
	chain       string
	appState    bool
	blockState  bool
	defaultHome = os.ExpandEnv("$HOME/.sprune")
	appName     = "sprune"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   appName,
		Short: "sprune is meant to prune data base history from a cosmos application, avoiding needing to state sync every couple amount of weeks",
	}

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, _ []string) error {
		if err := initConfig(rootCmd, appConfig); err != nil {
			return err
		}

		return nil
	}

	rootCmd.PersistentFlags().StringVar(&homePath, flags.FlagHome, defaultHome, "set home directory")
	if err := viper.BindPFlag(flags.FlagHome, rootCmd.PersistentFlags().Lookup(flags.FlagHome)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// rootCmd.PersistentFlags().Uint64VarP(&blocks, "blocks", "b", 10, "set the amount of blocks to keep (default=10)")
	// if err := viper.BindPFlag("blocks", rootCmd.PersistentFlags().Lookup("blocks")); err != nil {
	// 	panic(err)
	// }

	// rootCmd.PersistentFlags().StringVar(&chain, "chain", "", "set the app you are pruning (supported apps: osmosis)")
	// if err := viper.BindPFlag("chain", rootCmd.PersistentFlags().Lookup("chain")); err != nil {
	// 	panic(err)
	// }

	// rootCmd.PersistentFlags().BoolVar(&appState, "app-state", true, "set to false if using only with tendermint (default true)")
	// if err := viper.BindPFlag("app-state", rootCmd.PersistentFlags().Lookup("app-state")); err != nil {
	// 	panic(err)
	// }

	// rootCmd.PersistentFlags().BoolVar(&blockState, "block-state", true, "set to false you dont want to prune tendermint data(default true)")
	// if err := viper.BindPFlag("block-state", rootCmd.PersistentFlags().Lookup("block-state")); err != nil {
	// 	panic(err)
	// }

	rootCmd.AddCommand(
	// pruneCmd(),
	)

	return rootCmd
}
