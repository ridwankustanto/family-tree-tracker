package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ridwankustanto/family-tree-tracker/routes/account"
	"github.com/ridwankustanto/family-tree-tracker/utils/database"
)

func Routes(app *fiber.App) {
	api := app.Group("api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Regis your routes
	account.Routes(api, db)
}
