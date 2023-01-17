package middlewares

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidationMiddleware(validateStruct interface{}) fiber.Handler {
	validate := validator.New()
	return func(c *fiber.Ctx) error {
		if err := json.Unmarshal(c.Body(), validateStruct); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid request body",
			})
		}

		if err := validate.Struct(validateStruct); err != nil {
			errors := make(map[string]string)
			for _, err := range err.(validator.ValidationErrors) {
				errors[err.Field()] = fmt.Sprintf("%s, %s", err.Tag(), err.Error())
			}

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Validation failed",
				"errors":  errors,
			})
		}
		// bind the validated struct to the context, so it can be accessed by next middleware or route handler
		c.Locals("validated_struct", validateStruct)
		return c.Next()
	}
}
