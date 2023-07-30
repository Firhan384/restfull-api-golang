package main

import (
	"fmt"
	"log"

	wrap "github.com/Firhan384/restfull-api-golang/app"
	"github.com/Firhan384/restfull-api-golang/app/configs"
	"github.com/gofiber/fiber/v2"
	"moul.io/banner"
)

func main() {
	// load configs app
	configApp := configs.LoadConfigApp()

	// print banner app_name
	fmt.Println(banner.Inline(configApp.AppName))

	// config fiber
	app := fiber.New(fiber.Config{
		AppName:               configApp.AppName,
		DisableStartupMessage: true,
	})

	wrap.Bootstrap(app)

	log.Fatal(app.Listen(":" + configApp.AppPort))
}
