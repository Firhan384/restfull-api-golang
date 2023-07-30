package services

import (
	dto "github.com/Firhan384/restfull-api-golang/app/dto/user"
	"github.com/Firhan384/restfull-api-golang/app/schemas"
)

type UserService interface {
	CreateUser(data *dto.CreateUserDTO) schemas.ResponseApi
}
