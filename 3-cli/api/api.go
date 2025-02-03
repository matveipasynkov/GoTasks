package api

import "3-cli/app/config"

func GetConfig() *config.Config {
	return config.NewConfig()
}
