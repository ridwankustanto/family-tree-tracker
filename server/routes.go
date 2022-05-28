package server

import "github.com/gofiber/fiber/v2"

func routes(app *fiber.App) {
	api := app.Group("api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
