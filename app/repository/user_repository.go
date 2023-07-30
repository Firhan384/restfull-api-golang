package repository

import (
	"github.com/Firhan384/restfull-api-golang/app/database/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (model *UserRepo) List() ([]models.User, error) {
	var users []models.User
	if err := model.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (model *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	var user = models.User{Email: email}
	if err := model.DB.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (model *UserRepo) CreateUser(data *models.User) error {
	return model.DB.Create(data).Error
}
