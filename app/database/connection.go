package database

import (
	"log"

	"github.com/Firhan384/restfull-api-golang/app/configs"
	"github.com/Firhan384/restfull-api-golang/app/database/drivers"
	"gorm.io/gorm"
)

// global variable DB
var DB *gorm.DB

func ConnectionDB() {
	var err error
	cnf := configs.LoadConfigDatabase()

	if cnf.DBEngine == "mysql" {
		DB, err = drivers.DriverMysql()
	}

	if err != nil {
		log.Fatalln(err)
	}
}
