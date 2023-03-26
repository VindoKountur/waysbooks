package routes

import (
	"waysbooks/handlers"
	"waysbooks/pkg/middleware"
	"waysbooks/pkg/mysql"
	"waysbooks/repositories"

	"github.com/labstack/echo/v4"
)

func BookRoutes(e *echo.Group) {
	bookRepository := repositories.RepositoryBook(mysql.DB)
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)

	h := handlers.HandlerBook(bookRepository, transactionRepository)

	e.GET("/books", h.FindBooks)
	e.GET("/book/:id", h.GetBook)
	e.GET("/book-best", h.FindBestBooks)
	e.GET("/books-user", middleware.Auth(h.GetUserBooks))
	e.POST("/book", middleware.Auth(middleware.IsAdmin(middleware.UploadThumbnail(middleware.UploadPDF(h.CreateBook)))))
	e.DELETE("/book/:id", middleware.Auth(middleware.IsAdmin(h.DeleteBook)))
	e.PATCH("/book/:id", middleware.Auth(middleware.IsAdmin(middleware.UploadThumbnail(middleware.UploadPDF(h.UpdateBook)))))
}
