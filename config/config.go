package config

import "github.com/RichardKnop/machinery/v1/config"

var Config *config.Config

func LoadConfig(configFile string) {
	Config = config.NewFromYaml(configFile, true, true)
}
