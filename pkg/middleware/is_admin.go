package middleware

import (
	"net/http"
	dto "waysbooks/dto/result"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userLogin := c.Get("userLogin")

		userRole := userLogin.(jwt.MapClaims)["role"].(string)

		if userRole != "admin" {
			return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Code: http.StatusUnauthorized, Message: "Unauthorize, only admin can access this"})
		}
		return next(c)
	}
}
