package services

import "github.com/Firhan384/restfull-api-golang/app/repository"

type userService struct {
	repo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) UserService {
	return userService{repo: repo}
}
