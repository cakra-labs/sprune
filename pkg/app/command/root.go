package command

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const VERSION = "sprune/v7.0.0"

var (
	blocks      uint64
	versions    uint64
	backend     string
	app         string
	cosmosSdk   bool
	tendermint  bool
	defaultHome = os.ExpandEnv("$HOME/.sprune")
	appName     = "sprune"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   appName,
		Short: "sprune is meant to prune data base history from a cosmos application, avoiding needing to state sync every couple amount of weeks",
	}

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, _ []string) error {
		// reads `homeDir/config.yaml` into `var config *Config` before each command
		// if err := initConfig(rootCmd); err != nil {
		// 	return err
		// }

		return nil
	}

	rootCmd.PersistentFlags().Uint64VarP(&blocks, "blocks", "b", 10, "set the amount of blocks to keep (default=10)")
	if err := viper.BindPFlag("blocks", rootCmd.PersistentFlags().Lookup("blocks")); err != nil {
		panic(err)
	}

	rootCmd.PersistentFlags().Uint64VarP(&versions, "versions", "v", 10, "set the amount of versions to keep in the application store (default=10)")
	if err := viper.BindPFlag("versions", rootCmd.PersistentFlags().Lookup("versions")); err != nil {
		panic(err)
	}

	rootCmd.PersistentFlags().StringVar(&backend, "backend", "goleveldb", "set the type of db being used(default=goleveldb)") //todo add list of dbs to comment
	if err := viper.BindPFlag("backend", rootCmd.PersistentFlags().Lookup("backend")); err != nil {
		panic(err)
	}

	rootCmd.PersistentFlags().StringVar(&app, "app", "", "set the app you are pruning (supported apps: osmosis)")
	if err := viper.BindPFlag("app", rootCmd.PersistentFlags().Lookup("app")); err != nil {
		panic(err)
	}

	rootCmd.PersistentFlags().BoolVar(&cosmosSdk, "cosmos-sdk", true, "set to false if using only with tendermint (default true)")
	if err := viper.BindPFlag("cosmos-sdk", rootCmd.PersistentFlags().Lookup("cosmos-sdk")); err != nil {
		panic(err)
	}

	rootCmd.PersistentFlags().BoolVar(&tendermint, "tendermint", true, "set to false you dont want to prune tendermint data(default true)")
	if err := viper.BindPFlag("tendermint", rootCmd.PersistentFlags().Lookup("tendermint")); err != nil {
		panic(err)
	}

	rootCmd.AddCommand()

	return rootCmd
}
