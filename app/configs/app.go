package configs

import (
	"strconv"

	"github.com/Firhan384/restfull-api-golang/app/utils"
)

type ConfigApp struct {
	AppName  string
	AppPort  string
	AppDebug bool
	AppEnv   string
}

func LoadConfigApp() ConfigApp {
	config := ConfigApp{}

	// convert
	convAppDebug, _ := strconv.ParseBool(utils.GetEnv("APP_DEBUG"))

	config.AppName = utils.GetEnv("APP_NAME")
	config.AppDebug = convAppDebug
	config.AppPort = utils.GetEnv("APP_PORT")
	config.AppEnv = utils.GetEnv("APP_ENV")

	return config
}
