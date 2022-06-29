package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	// "github.com/golang-jwt/jwt/v4"
	"github.com/ridwankustanto/family-tree-tracker/routes/account"
	"github.com/ridwankustanto/family-tree-tracker/routes/location"
	"github.com/ridwankustanto/family-tree-tracker/utils"
	"github.com/ridwankustanto/family-tree-tracker/utils/database"
	"github.com/ridwankustanto/family-tree-tracker/utils/middlewares"

)

func Routes(app *fiber.App) {
	api := app.Group("api")
	// c := &fiber.Ctx{}
	test := app.Group("test")
	api.Get("/", utils.Restrict(), func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	test.Use(middlewares.Restrict())
	test.Get("/", func(c *fiber.Ctx) error {
		err:=middlewares.Authorize(c)
		log.Println(err)
		if err != nil {
			return middlewares.GetOut(c)
		}
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
