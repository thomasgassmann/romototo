package config

import (
	"github.com/goccy/go-yaml"
	"io/ioutil"
)

type Config struct {
	Mail struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"mail"`
}

func ParseConfig(configPath string) (*Config, error) {
	config := &Config{}

	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(bytes, config); err != nil {
		return nil, err
	}

	return config, nil
}
