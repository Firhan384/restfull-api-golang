package controllers

import (
	"github.com/Firhan384/restfull-api-golang/app/schemas"
	services "github.com/Firhan384/restfull-api-golang/app/services/user"
	"github.com/gofiber/fiber/v2"
)

type userController struct {
	service services.UserService
}

func NewUserController(service services.UserService) userController {
	return userController{service: service}
}

func (c userController) Route(app *fiber.App) {
	app.Get("/api/user", getUsers)
}

func getUsers(c *fiber.Ctx) error {
	resp := schemas.NewResponseApi()
	resp.StatusCode = fiber.StatusOK
	resp.Message = "I'm a GET request!"
	resp.Error = false
	resp.Data = map[string]interface{}{
		"email": "arjun",
	}

	return c.Status(resp.StatusCode).JSON(resp)
}
