package config

import (
	"errors"
	"folion-server/src/system/paths"
	"folion-server/src/system/x"
	"os"

	"gopkg.in/yaml.v3"
)

var instance *Config = nil

type Config struct {
	DeveloperMode bool   `yaml:"devmode"`
	Address       string `yaml:"address"`
	Port          string `yaml:"port"`
	MaxWorkers    int    `yaml:"max_workers"`
}

var defaultConfig Config = Config{
	DeveloperMode: false,
	Address:       "127.0.0.1",
	Port:          "8080",
	MaxWorkers:    5,
}

func GetInstance() *Config {
	if instance != nil {
		return instance
	} else {
		return nil
	}
}

func LoadConfig() error {
	if exists, _ := configExists(); exists == false {
		createDefaultConfig()
	}

	data, err := os.ReadFile(paths.ConfigPath)

	if err != nil {
		return x.ErrConfigLoadFailed.Wrap(err)
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return x.ErrConfigLoadFailed.Wrap(err)
	}

	instance = &config
	return nil
}

func createDefaultConfig() {}
func configExists() (bool, error) {
	_, err := os.Stat(paths.ConfigPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, x.ErrConfigNotFound
		}
	}
	return true, nil
}
