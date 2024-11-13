package handlers

import (
	"github.com/NuttayotSukkum/acn/acn-authentication/configs"
	"github.com/NuttayotSukkum/acn/acn-authentication/internal/constants"
	"github.com/NuttayotSukkum/acn/acn-authentication/internal/models"
	"github.com/NuttayotSukkum/acn/acn-authentication/internal/models/response"
	"github.com/NuttayotSukkum/acn/acn-authentication/internal/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserHandler struct {
	service *services.UserRepo
}

func NewUserHandler(service *services.UserRepo) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) RegisterUser(e echo.Context) error {
	var user *models.User
	if err := e.Bind(&user); err != nil {
		return e.JSON(http.StatusBadRequest, response.MessageResponse{
			HttpStatus: strconv.Itoa(http.StatusBadRequest),
			Time:       constants.TIME_NOW,
			Message:    err.Error(),
		})
	}

	if err := h.service.CreateUser(user); err != nil {
		return e.JSON(http.StatusInternalServerError, response.MessageResponse{
			HttpStatus: strconv.Itoa(http.StatusInternalServerError),
			Time:       constants.TIME_NOW,
			Message:    err.Error(),
		})
	}

	return e.JSON(http.StatusOK, response.MessageResponse{
		HttpStatus: strconv.Itoa(http.StatusOK),
		Time:       constants.TIME_NOW,
		Message:    "User registered successfully",
	})
}

func (h *UserHandler) LoginHandler(e echo.Context) error {
	var user models.User
	cfg := configs.InitConfig()
	if err := e.Bind(&user); err != nil {
		return e.JSON(http.StatusBadRequest, response.MessageResponse{
			HttpStatus: strconv.Itoa(http.StatusBadRequest),
			Time:       constants.TIME_NOW,
			Message:    err.Error(),
		})
	}
	token, err := h.service.LoginUser(user.Email, user.Password, *cfg)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, response.MessageResponse{
			HttpStatus: strconv.Itoa(http.StatusInternalServerError),
			Time:       constants.TIME_NOW,
			Message:    err.Error(),
		})
	}

	return e.JSON(http.StatusOK, response.ResponseData{
		HttpStatus: strconv.Itoa(http.StatusOK),
		Time:       constants.TIME_NOW,
		Data:       *token,
	})

}

func (h *UserHandler) VerifyUserHandler(e echo.Context) error {
	ID := e.Param("id")
	if ID == "" || ID == "0" {
		return e.JSON(http.StatusBadRequest, response.MessageResponse{
			HttpStatus: strconv.Itoa(http.StatusBadRequest),
			Time:       constants.TIME_NOW,
			Message:    "User ID is empty",
		})
	}
	res, err := h.service.VerifyUser(ID)
	if err != nil {
		return e.JSON(http.StatusNotFound, response.MessageResponse{
			HttpStatus: strconv.Itoa(http.StatusNotFound),
			Time:       constants.TIME_NOW,
			Message:    err.Error(),
		})
	}

	return e.JSON(http.StatusOK, response.ResponseUser{
		HTTPStatus: strconv.Itoa(http.StatusOK),
		Time:       constants.TIME_NOW,
		Data:       res,
	})
}
