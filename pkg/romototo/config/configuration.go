package config

import (
	"github.com/goccy/go-yaml"
	"io/ioutil"
)

type MailConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	Mail MailConfig `yaml:"mail"`
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
