package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	AuthRoutes(e)
	ProfileRoutes(e)
	BookRoutes(e)
	TransactionRoutes(e)
	CartRoutes(e)
	UserRoutes(e)
	// AddressRoutes(e)
	// CartRoutes(e)
}
