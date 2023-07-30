package database

import "github.com/Firhan384/restfull-api-golang/app/database/models"

func Migrations() {
	DB.AutoMigrate(&models.User{})
}
