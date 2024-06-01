package config

import (
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type Config struct {
	AppName string `yaml:"-" json:"-"`

	ChainDir string `yaml:"chain_dir" json:"chain_dir"`
	DataDir  string `yaml:"data_dir" json:"data_dir"`
	Chain    string `yaml:"chain" json:"chain"`
	LogLevel string `yaml:"log_level" json:"log_level"`

	AppState     bool   `yaml:"app_state" json:"app_state"`
	BlockState   bool   `yaml:"block_state" json:"block_state"`
	BlocksToKeep uint64 `yaml:"blocks_to_keep" json:"blocks_to_keep"`
}

func CreateConfig(home string) error {
	cfgPath := path.Join(home, "config.yaml")

	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		if _, err := os.Stat(home); os.IsNotExist(err) {
			if err = os.Mkdir(home, os.ModePerm); err != nil {
				return err
			}
		}
	}

	f, err := os.Create(cfgPath)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.Write(defaultConfig()); err != nil {
		return err
	}
	return nil
}

func (c Config) MustYAML() []byte {
	out, err := yaml.Marshal(c)
	if err != nil {
		panic(err)
	}
	return out
}

func ValidateConfig(c *Config) error {
	return nil
}

func defaultConfig() []byte {
	return Config{
		ChainDir:     "/home/user/.gaia",
		DataDir:      "data",
		Chain:        "",
		LogLevel:     "info",
		AppState:     true,
		BlockState:   true,
		BlocksToKeep: 10,
	}.MustYAML()
}
