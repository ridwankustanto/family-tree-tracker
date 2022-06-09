package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ridwankustanto/family-tree-tracker/server/routes/account"
)

func Routes(app *fiber.App) {
	api := app.Group("api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Regis your routes
	account.Routes(api)
}
