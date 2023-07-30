package dto

type CreateUserDTO struct {
	Fullname string `json:"fullname" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
