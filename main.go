package main

import (
	"flag"

	"github.com/thomasgassmann/romototo/cmd"
	"github.com/thomasgassmann/romototo/pkg/romototo/config"
)

func main() {
	configPath := parseConfigPath()
	config, err := config.ParseConfig(configPath)
	if err != nil {
		panic(err)
	}

	cmd.Execute(config)
}

func parseConfigPath() string {
	var configPath string
	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")
	flag.Parse()
	return configPath
}
