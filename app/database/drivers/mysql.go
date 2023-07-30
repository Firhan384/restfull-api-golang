package drivers

import (
	"fmt"

	"github.com/Firhan384/restfull-api-golang/app/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DriverMysql() (*gorm.DB, error) {
	cfg := configs.LoadConfigDatabase()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
