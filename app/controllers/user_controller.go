package controllers

import (
	"github.com/Firhan384/restfull-api-golang/app/database"
	dto "github.com/Firhan384/restfull-api-golang/app/dto/user"
	"github.com/Firhan384/restfull-api-golang/app/repository"
	"github.com/Firhan384/restfull-api-golang/app/schemas"
	services "github.com/Firhan384/restfull-api-golang/app/services/user"
	"github.com/Firhan384/restfull-api-golang/app/utils"
	"github.com/gofiber/fiber/v2"
)

type userController struct {
	service services.UserService
}

func NewUserController() userController {
	repository := repository.NewUserRepo(database.DB)
	service := services.NewUserService(repository)

	return userController{service: service}
}

func (c userController) Route(app fiber.Router) {
	app.Get("/user", c.getUsers)
	app.Post("/user", c.createUser)
}

func (c userController) getUsers(ctx *fiber.Ctx) error {
	resp := schemas.NewResponseApi()
	resp.StatusCode = fiber.StatusOK
	resp.Message = "I'm a GET request!"
	resp.Error = false
	resp.Data = map[string]interface{}{
		"email": "arjun",
	}

	return ctx.Status(resp.StatusCode).JSON(resp)
}

func (c userController) createUser(ctx *fiber.Ctx) error {
	resp := schemas.NewResponseApi()
	request := new(dto.CreateUserDTO)

	// check json
	if err := ctx.BodyParser(&request); err != nil {
		resp.StatusCode = fiber.StatusBadRequest
		resp.Error = true
		resp.Message = utils.StatusText(fiber.StatusBadRequest)

		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	validate := utils.NewValidator()

	// check validator
	if err := validate.Struct(request); err != nil {
		resp.StatusCode = fiber.StatusBadRequest
		resp.Error = true
		resp.DetailMessage = utils.ValidatorErrors(err)
		resp.Message = fiber.ErrBadRequest.Error()

		return ctx.Status(fiber.StatusBadRequest).JSON(resp)
	}

	// call the business logic
	service := c.service.CreateUser(request)

	return ctx.Status(service.StatusCode).JSON(service)
}
