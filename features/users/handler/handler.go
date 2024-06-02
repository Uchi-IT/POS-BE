package handler

import (
	"uchiiParfume/features/users/entity"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService entity.UsersServiceInterface
}

func NewUserHandler(user entity.UsersServiceInterface) *userHandler {
	return &userHandler{
		userService: user,
	}
}

func (user *userHandler) CreateUesr(e echo.Context) error {
	return nil
}

func (user *userHandler) GetAllUser(e echo.Context) error {
	return nil
}

func (user *userHandler) GetSpecificUser(e echo.Context) error {
	return nil
}

func (user *userHandler) UpdateUser(e echo.Context) error {
	return nil
}

func (user *userHandler) DeleteUser(e echo.Context) error {
	return nil
}
