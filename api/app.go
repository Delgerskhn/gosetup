package main

import (
	"github.com/delgerskhn/gosetup/api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.RegisterApis(app)

	app.Listen(":3000")
}
