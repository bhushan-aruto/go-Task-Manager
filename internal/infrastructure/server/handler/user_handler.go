package handler

import (
	"net/http"

	"github.com/bhushan-aruto/go-task-manager/internal/entity"
	"github.com/bhushan-aruto/go-task-manager/internal/infrastructure/server/model"
	"github.com/bhushan-aruto/go-task-manager/internal/usecase"
	"github.com/bhushan-aruto/go-task-manager/internal/utils/jwtutil"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	usecase *usecase.UserUsecase
}

func NewUserHandler(u *usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

func (h *UserHandler) Register(c echo.Context) error {
	var req model.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	user := &entity.User{
		Email:    req.Email,
		Password: req.Password,
	}
	err := h.usecase.Register(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "could not register user"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "user registered"})
}

func (h *UserHandler) Login(c echo.Context) error {
	var req model.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	user, err := h.usecase.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid credentials"})
	}

	token, err := jwtutil.GenerateJWT(user.ID.Hex())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "could not generate token"})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}
