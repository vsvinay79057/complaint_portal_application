package controller

import (
	"complaint_portal/models"
	"complaint_portal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUC usecase.UserUsecase
}

func NewUserController(u usecase.UserUsecase) *UserController {
	return &UserController{userUC: u}
}

func (c *UserController) Register(ctx echo.Context) error {
	var req models.RegisterRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid payload"})
	}
	if req.Name == "" || req.Email == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "name and email are required"})
	}
	user, err := c.userUC.Register(req.Name, req.Email)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) Login(ctx echo.Context) error {
	var req models.LoginRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid payload"})
	}
	user, err := c.userUC.LoginWithSecret(req.SecretCode)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, user)
}
func (c *UserController) CreateAdmin(ctx echo.Context) error {
	type req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		AdminKey string `json:"admin_key"`
	}
	var r req
	if err := ctx.Bind(&r); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid payload"})
	}
	if r.Name == "" || r.Email == "" || r.AdminKey == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "name, email and admin_key required"})
	}
	u, err := c.userUC.CreateAdmin(r.Name, r.Email, r.AdminKey)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusCreated, u)
}
