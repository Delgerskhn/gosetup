package main

import (
	//the config initialization must exec before other packages
	_ "github.com/delgerskhn/gosetup/config"

	"github.com/delgerskhn/gosetup/pkg/presentation/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.RegisterV1Apis(app)

	app.Listen(":3000")
}
