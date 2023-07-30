package configs

import "github.com/Firhan384/restfull-api-golang/app/utils"

type configDatabase struct {
	DBHost   string
	DBPort   string
	DBEngine string
	DBName   string
	DBUser   string
	DBPass   string
}

func LoadConfigDatabase() configDatabase {
	return configDatabase{
		DBHost:   utils.GetEnv("DB_HOST"),
		DBPort:   utils.GetEnv("DB_PORT"),
		DBEngine: utils.GetEnv("DB_ENGINE"),
		DBName:   utils.GetEnv("DB_NAME"),
		DBUser:   utils.GetEnv("DB_USER"),
		DBPass:   utils.GetEnv("DB_PASS"),
	}
}
