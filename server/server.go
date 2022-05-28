package server

import "github.com/gofiber/fiber/v2"

func Run() {
	app := fiber.New()
	app.Static("/", "./public")
	routes(app)
	app.Listen(":3000")
}
