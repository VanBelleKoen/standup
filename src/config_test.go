package main

import (
	"os"
	"os/user"
	"path/filepath"
	"testing"
)

func TestLoadConfig_NoConfigFile(t *testing.T) {
	usr, err := user.Current()
	if err != nil {
		t.Fatalf("Failed to get current user: %v", err)
	}
	configPath := filepath.Join(usr.HomeDir, ".standupconfig")

	os.Remove(configPath)

	config, err := loadConfig()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(config.Workspaces) != 0 {
		t.Errorf("Expected no workspaces, got %v", config.Workspaces)
	}
}

func TestCreateDefaultConfig(t *testing.T) {
	usr, err := user.Current()
	if err != nil {
		t.Fatalf("Failed to get current user: %v", err)
	}
	configPath := filepath.Join(usr.HomeDir, ".standupconfig")

	os.Remove(configPath)

	err = createDefaultConfig()
	if err != nil {
		t.Fatalf("Failed to create default config: %v", err)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Fatalf("Expected config file to be created, but it does not exist")
	}
	
	config, err := loadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if len(config.Workspaces) != 1 || config.Workspaces[0] != "/path/to/default/workspace" {
		t.Errorf("Default config does not match expected values: %v", config.Workspaces)
	}
}