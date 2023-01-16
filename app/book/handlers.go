package book

import (
	"github.com/delgerskhn/gosetup/app/book/inputs"
	"github.com/delgerskhn/gosetup/app/book/repository"
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	books, err := repository.GetBooks()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(books)
}

func CreateBook(c *fiber.Ctx) error {
	var book inputs.CreateBookInput
	if err := c.BodyParser(&book); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	createdBook, err := repository.CreateBook(book.ToEntity())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(createdBook)
}
