package middleware

import (
	"github.com/G-rillei/go-simple-server/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateBody(dto interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := c.BodyParser(dto); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
		}

		validate := validator.New(validator.WithRequiredStructEnabled())

		if err := validate.Struct(dto); err != nil {
			errMsg := utils.ValidateBodyErrMsg(err)

			return c.Status(400).JSON(fiber.Map{
				"errors":  errMsg,
				"message": "Body Validation Error",
			})
		}

		return c.Next()
	}
}
