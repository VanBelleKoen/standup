package main

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Workspaces []string `toml:"workspaces"`
}

func loadConfig() (Config, error) {
	usr, err := user.Current()
	if err != nil {
		return Config{}, err
	}
	configPath := filepath.Join(usr.HomeDir, ".standupconfig")
	var config Config
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := createDefaultConfig(); err != nil {
			return Config{}, err
		}
		return config, nil
	}
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		return config, err
	}
	return config, nil
}

func createDefaultConfig() error {
	usr, err := user.Current()
	if err != nil {
		return err
	}
	configPath := filepath.Join(usr.HomeDir, ".standupconfig")
	defaultConfig := Config{
		Workspaces: []string{"/path/to/default/workspace"},
	}
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()
	return toml.NewEncoder(file).Encode(defaultConfig)
}