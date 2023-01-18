package book

import (
	"github.com/gofiber/fiber/v2"
)

func HandleChangeAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	ChangeAuthor(id, "new author")
	return c.SendStatus(200)
}

func HandleGetBooks(c *fiber.Ctx) error {
	books, err := GetBooks()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(books)
}

func HandleCreateBook(c *fiber.Ctx) error {
	book := c.Locals("validated_struct").(*CreateBookInput)

	createdBook, err := CreateBook(book.ToEntity())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(createdBook)
}
