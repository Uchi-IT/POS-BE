package handler

import (
	"net/http"
	dto "uchiiParfume/features/users/dto"
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

func (handler *userHandler) CreateUser(e echo.Context) error {
	input := new(dto.UserRequest)
	errBind := e.Bind(&input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := entity.UsersCore{
		Email:    input.Email,
		Password: input.Password,
		Cabang:   input.Cabang,
	}

	row, errUser := handler.userService.CreateUser(data)
	if errUser != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error create user",
			"error":   errUser.Error(),
		})
	}

	resp := dto.UserResponse{
		Id: row.Id,
		Email: row.Email,
		Cabang: row.Cabang,
		Role: row.Role,
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "succes create user",
		"data":    resp,
	})
}

func (handler *userHandler) GetAllUser(e echo.Context) error {
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

	data, err := handler.userService.GetAllUser()
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all user",
		})
	}

	dataList := []dto.UserResponse{}
	for _, v := range data {
		result := dto.UserResponse{
			Id:     v.Id,
			Email:  v.Email,
			Cabang: v.Cabang,
			Role:   v.Role,
		}
		dataList = append(dataList, result)
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get all user",
		"data":    dataList,
	})
}

func (handler *userHandler) GetSpecificUser(e echo.Context) error {
	idParamstr := e.Param("id")

	idParams, err := uuid.Parse(idParamstr)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "user not found",
		})
	}

	data, err := handler.userService.GetById(idParams.String())
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get specific user",
		})
	}

	response := dto.UserResponse{
		Id:     data.Id,
		Email:  data.Email,
		Cabang: data.Cabang,
		Role:   data.Role,
	}
	return e.JSON(http.StatusOK, map[string]any{
		"message": "get user",
		"data":    response,
	})
}

func (handler *userHandler) UpdateUser(e echo.Context) error {
	idParams := e.Param("id")

	data := new(dto.UserRequest)
	if errBind := e.Bind(data); errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error binding data",
		})
	}

	userData := entity.UsersCore{
		Id:       idParams,
		Email:    data.Email,
		Cabang:   data.Cabang,
		Password: data.Password,
	}

	err := handler.userService.UpdateUser(idParams, userData)
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

func (handler *userHandler) DeleteUser(e echo.Context) error {
	idParams := e.Param("id")
	err := handler.userService.DeleteUser(idParams)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error deleting user",
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "User deleted successfully",
	})
}

func (handler *userHandler) Login(e echo.Context) error {
	input := new(dto.UserRequest)
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

	data, token, err := handler.userService.Login(data.Email, data.Password)
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
