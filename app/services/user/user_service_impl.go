package services

import (
	dto "github.com/Firhan384/restfull-api-golang/app/dto/user"
	"github.com/Firhan384/restfull-api-golang/app/repository"
	"github.com/Firhan384/restfull-api-golang/app/schemas"
	"github.com/gofiber/fiber/v2"
)

type userService struct {
	repo repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) UserService {
	return userService{repo: *repo}
}

func (r userService) CreateUser(data *dto.CreateUserDTO) schemas.ResponseApi {
	resp := schemas.NewResponseApi()
	resp.StatusCode = fiber.StatusOK
	resp.Message = "OK"

	return resp
}
