package config

import (
	"github.com/goccy/go-yaml"
	"os"
)

type Config struct {
	Mail struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

func ParseConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
