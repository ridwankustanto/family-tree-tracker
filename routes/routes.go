package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	// "github.com/golang-jwt/jwt/v4"
	"github.com/ridwankustanto/family-tree-tracker/routes/account"
	"github.com/ridwankustanto/family-tree-tracker/routes/location"
	// "github.com/ridwankustanto/family-tree-tracker/utils"
	"github.com/ridwankustanto/family-tree-tracker/utils/database"
	"github.com/ridwankustanto/family-tree-tracker/utils/middlewares"
)

func Routes(app *fiber.App) {
	api := app.Group("api")
	// c := &fiber.Ctx{}
	api.Get("/", middlewares.Restrict(), func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// api.Get("/test/:id", func(c *fiber.Ctx) error {
	// 	log.Println(c.Params("id"))
	// 	return nil
	// })

	api.Get("/test/", func(c *fiber.Ctx) error {
		log.Println(c.Get("Authorization"))
		return nil
	})

	// Regis your routes
	account.Routes(api, db)
	location.Routes(api, db)

}
