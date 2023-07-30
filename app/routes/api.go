package routes

import (
	"github.com/Firhan384/restfull-api-golang/app/controllers"
	"github.com/Firhan384/restfull-api-golang/pkg"
	"github.com/gofiber/fiber/v2"
)

func RegisterRouteAPI(app *fiber.App) {
	// set prefix api
	groupApi := app.Group("/api")

	// register route in controller
	user := controllers.NewUserController()
	user.Route(groupApi)

	// list all routes
	pkg.ListRouters(app)
}
