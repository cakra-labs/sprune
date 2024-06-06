package command

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/cakra-labs/sprune/pkg/app/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var appConfig = &config.Config{
	AppName: appName,
}

func initConfig(cfg *config.Config) error {
	// Check config path
	cfgPath := path.Join(homePath, "config.yaml")
	_, err := os.Stat(cfgPath)
	if err != nil {
		err = config.CreateConfig(homePath)
		if err != nil {
			return err
		}
	}
	viper.SetConfigFile(cfgPath)
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Failed to read in config:", err)
		os.Exit(1)
	}

	// Read config file
	file, err := os.ReadFile(viper.ConfigFileUsed())
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	if err := yaml.Unmarshal(file, cfg); err != nil {
		fmt.Println("Error unmarshalling config:", err)
		os.Exit(1)
	}

	// Validate configuration
	if err := config.ValidateConfig(cfg); err != nil {
		fmt.Println("Error parsing chain config:", err)
		os.Exit(1)
	}
	return nil
}

func configCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Validate and show config",
		RunE: func(cmd *cobra.Command, args []string) error {
			// print to json
			b, err := json.MarshalIndent(appConfig, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(b))

			return nil
		},
	}
	return cmd
}
