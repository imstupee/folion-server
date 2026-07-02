package system

import (
	"folion-server/src/system/errors"
	"folion-server/src/system/paths"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DeveloperMode bool   `yaml:"devmode"`
	Address       string `yaml:"address"`
	Port          int    `yaml:"port"`
}

func LoadConfig() (*Config, error) {
	data, err := os.ReadFile(paths.ConfigPath)

	if err != nil {
		return nil, errors.ErrConfigLoadFailed
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, errors.ErrConfigLoadFailed
	}

	return &config, nil
}
