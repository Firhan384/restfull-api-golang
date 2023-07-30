package pkg

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func ListRouters(app *fiber.App) {
	routes := app.GetRoutes(true)

	for _, router := range routes {
		log.Info(router.Method, " ", router.Path)
	}
}
