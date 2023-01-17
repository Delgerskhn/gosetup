package routes

import "github.com/gofiber/fiber/v2"

func RegisterV1Apis(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	BindBookApis(v1)
}
