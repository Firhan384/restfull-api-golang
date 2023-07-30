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
	app := fiber.New(configs.NewConfigFiber(configApp))

	wrap.Bootstrap(app)

	log.Fatal(app.Listen(":" + configApp.AppPort))
}
