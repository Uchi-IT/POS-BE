package handler

import (
	"fmt"
	"net/http"
	handler "uchiiParfume/features/users/dto"
	"uchiiParfume/features/users/entity"
	middleware "uchiiParfume/utils/jwt"

	"github.com/google/uuid"
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

func (user *userHandler) CreateUser(e echo.Context) error {
	input := handler.UserRequest{}
	errBind := e.Bind(&input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	fmt.Println("input handler :")
	fmt.Println(input)

	userInput := handler.UserRequestToUserCore(input)

	fmt.Println("input handler sesudah mapping:")
	fmt.Println(userInput)

	_, errUser := user.userService.CreateUser(userInput)
	if errUser != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error create user",
			"error":   errUser.Error(),
		})
	}
	

	return e.JSON(http.StatusOK, map[string]any{
		"message": "succes create user",
	})
}

func (user *userHandler) GetAllUser(e echo.Context) error {
	_, role, err := middleware.ExtractToken(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	if role != "admin" {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "access denied",
		})
	}

	data, err := user.userService.GetAllUser()
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all user",
		})
	}

	dataList := handler.ListUserCoreToListUserResponse(data)

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get all user",
		"data":    dataList,
	})
}

func (user *userHandler) GetSpecificUser(e echo.Context) error {
	idParamstr := e.Param("id")

	idParams, err := uuid.Parse(idParamstr)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "user not found",
		})
	}

	data, err := user.userService.GetById(idParams.String())
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get specific user",
		})
	}

	response := handler.UserCoreToUserResponse(data)
	
	return e.JSON(http.StatusOK, map[string]any{
		"message": "get user",
		"data":    response,
	})
}

func (user *userHandler) UpdateUser(e echo.Context) error {
	idParams := e.Param("id")

	data := handler.UserRequest{}
	if errBind := e.Bind(data); errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error binding data",
		})
	}

	userData := handler.UserRequestToUserCore(data)

	err := user.userService.UpdateUser(idParams, userData)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error updating user",
			"error":   err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "user updated successfully",
		"data":    userData,
	})
}

func (user *userHandler) DeleteUser(e echo.Context) error {
	idParams := e.Param("id")
	err := user.userService.DeleteUser(idParams)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error deleting user",
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "User deleted successfully",
	})
}

func (user *userHandler) Login(e echo.Context) error {
	input := new(handler.UserRequest)
	errBind := e.Bind(&input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := entity.UsersCore{
		Email:    input.Email,
		Password: input.Password,
	}

	data, token, err := user.userService.Login(data.Email, data.Password)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error login",
			"error":   err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "login success",
		"email":   data.Email,
		"token":   token,
	})
}
