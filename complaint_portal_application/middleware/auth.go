package middleware

import (
	"complaint_portal/models"
	"complaint_portal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(u usecase.UserUsecase) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			secret := c.Request().Header.Get("secrete code")
			if secret == "" {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "secrete code header required"})
			}
			user, err := u.LoginWithSecret(secret)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid secret code"})
			}
			c.Set("user", user)
			return next(c)
		}
	}
}
func AdminOnlyMiddleware(u usecase.UserUsecase) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			v := c.Get("user")
			if v == nil {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "not authenticated"})
			}
			user := v.(*models.UserModel)
			if !user.IsAdmin {
				return c.JSON(http.StatusForbidden, echo.Map{"error": "admin access required"})
			}
			return next(c)
		}
	}
}
