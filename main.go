package main

import (
	"github.com/G-rillei/go-simple-server/constants"
	"github.com/G-rillei/go-simple-server/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	constants := constants.GetConstants()

	app := fiber.New()
	api := app.Group("/api")

	route.PostRoutes(api)

	app.Listen(constants.PORT)
}
