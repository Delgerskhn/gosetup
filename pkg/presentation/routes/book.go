package routes

import (
	"github.com/delgerskhn/gosetup/pkg/application/book"
	"github.com/delgerskhn/gosetup/pkg/presentation/middlewares"
	"github.com/gofiber/fiber/v2"
)

func BindBookApis(app fiber.Router) {
	group := app.Group("/books")
	group.Get("/", book.HandleGetBooks)
	group.Post("/", middlewares.ValidationMiddleware(&book.CreateBookInput{}), book.HandleCreateBook)
	group.Put("/:id", book.HandleChangeAuthor)
}
