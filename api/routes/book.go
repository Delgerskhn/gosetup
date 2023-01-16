package routes

import (
	"github.com/delgerskhn/gosetup/app/book"
	"github.com/gofiber/fiber/v2"
)

func BindBookApis(app fiber.Router) {
	app.Group("/books").Get("/", book.GetBooks).Post("/", book.CreateBook)
}
