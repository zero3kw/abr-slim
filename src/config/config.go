// abr-slim/src/config/config.go

package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	BaseURL  string `yaml:"baseURL"`
	Datasets []struct {
		ID   string `yaml:"id"`
		File string `yaml:"file"`
	} `yaml:"datasets"`
	WorkingDir string `yaml:"workingDir"`
}

func LoadConfig(configPath string) (*Config, error) {
	var config Config
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(configFile, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
