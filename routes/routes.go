package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ridwankustanto/family-tree-tracker/routes/account"
	"github.com/ridwankustanto/family-tree-tracker/routes/location"
	"github.com/ridwankustanto/family-tree-tracker/utils"
	"github.com/ridwankustanto/family-tree-tracker/utils/database"
)

func Routes(app *fiber.App) {
	api := app.Group("api")
	api.Get("/", utils.Restrict(), func(c *fiber.Ctx) error {
		utils.Authorize(c)
		return c.SendString("Hello, World!")
	})

	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Regis your routes
	account.Routes(api, db)
	location.Routes(api, db)

}
