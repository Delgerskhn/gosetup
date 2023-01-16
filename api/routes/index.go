package routes

import "github.com/gofiber/fiber/v2"

func RegisterApis(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	BindBookApis(v1)
}
