package configuration

import (
	"flag"
)

var (
	configPath string
)

func Init() string {
	flag.StringVar(&configPath, "config", "./provider", "setuping plugins path")
	flag.Parse()
	return configPath
}
