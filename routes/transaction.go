package routes

import (
	"waysbooks/handlers"
	"waysbooks/pkg/middleware"
	"waysbooks/pkg/mysql"
	"waysbooks/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	// cartRepository := repositories.RepositoryCart(mysql.DB)
	// productRepository := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerTransaction(
		transactionRepository,
		// cartRepository,
		// productRepository
	)

	e.POST("/transaction", middleware.Auth(h.CreateTransaction))
	e.GET("/transaction/:id", h.GetTransaction)
	e.GET("/transactions-user", middleware.Auth(h.GetUserTransaction))
	e.GET("/transactions", h.FindTransaction)
	e.POST("/notification", h.Notification)
}
