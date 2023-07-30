package app

import (
	"github.com/Firhan384/restfull-api-golang/app/database"
	"github.com/Firhan384/restfull-api-golang/app/routes"
	"github.com/gofiber/fiber/v2"
)

// wrapping all functionality in folder app
func Bootstrap(app *fiber.App) {
	// connect DB
	database.ConnectionDB()

	// migrations
	database.Migrations()

	// seeders
	// database.Seeder()

	// register route
	routes.RegisterRouteAPI(app)
}
