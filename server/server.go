package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ridwankustanto/family-tree-tracker/server/routes"
)

func Run() {
	app := fiber.New()
	app.Static("/", "./public")
	routes.Routes(app)
	app.Listen(":3000")
}
