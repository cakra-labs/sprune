package config

import "gopkg.in/yaml.v2"

type Config struct {
	HomePath string `yaml:"home_path" json:"home_path"`

	DataDir  string `yaml:"data_dir" json:"data_dir"`
	Backend  string `yaml:"backend" json:"backend"`
	App      string `yaml:"app" json:"app"`
	LogLevel string `yaml:"log_level" json:"log_level"`

	CosmosSdk  bool   `yaml:"cosmos_sdk" json:"cosmos_sdk"`
	Tendermint bool   `yaml:"tendermint" json:"tendermint"`
	Blocks     uint64 `yaml:"blocks" json:"blocks"`
	Version    uint64 `yaml:"version" json:"version"`
	AppName    string
}

func (c Config) MustYAML() []byte {
	out, err := yaml.Marshal(c)
	if err != nil {
		panic(err)
	}
	return out
}
