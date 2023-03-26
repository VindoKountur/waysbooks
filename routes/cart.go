package routes

import (
	"waysbooks/handlers"
	"waysbooks/pkg/middleware"
	"waysbooks/pkg/mysql"
	"waysbooks/repositories"

	"github.com/labstack/echo/v4"
)

func CartRoutes(e *echo.Group) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(cartRepository)

	e.GET("/cart-user", middleware.Auth(h.GetUserCartList))
	e.POST("/cart/:id", middleware.Auth(h.AddToCart))
	e.DELETE("/cart/:id", middleware.Auth(h.RemoveBookFromCart))
}
