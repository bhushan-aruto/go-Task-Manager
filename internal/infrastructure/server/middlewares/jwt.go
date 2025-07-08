package middlewares

import (
	"net/http"

	"github.com/bhushan-aruto/go-task-manager/internal/utils/jwtutil"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := jwtutil.ExtractUserID(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
		}
		c.Set("user_id", userId)
		return next(c)
	}
}
