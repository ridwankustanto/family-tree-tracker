package location

import (
	"database/sql"
	// "log"

	"github.com/gofiber/fiber/v2"
	locationClient "github.com/ridwankustanto/family-tree-tracker/clients/location"
	locationRepo "github.com/ridwankustanto/family-tree-tracker/repository/location"
	locationService "github.com/ridwankustanto/family-tree-tracker/services/location"
	"github.com/ridwankustanto/family-tree-tracker/utils/middlewares"
)

func Routes(api fiber.Router, db *sql.DB) {
	repo:= locationRepo.NewPostgresRepository(db)
	
	srv:= locationService.NewService(repo)

	location := api.Group("location")
	location.Use(middlewares.Restrict())
	location.Post("/:type", func(c *fiber.Ctx) error {
		return locationClient.CreateLocation(c, srv)
	})

	location.Get("/country", func(c *fiber.Ctx) error {
		return locationClient.GetAllCountry(c, srv)
	})

	location.Get("/country/:id", func(c *fiber.Ctx) error {
		return locationClient.GetCountry(c, srv)
	})
	
	location.Get("/province/:id", func(c *fiber.Ctx) error {
		return locationClient.GetProvince(c, srv)
	})
	
	location.Get("/city/:id", func(c *fiber.Ctx) error {
		return locationClient.GetCity(c, srv)
	})

	location.Get("/district/:id", func(c *fiber.Ctx) error {
		return locationClient.GetDistrict(c, srv)
	})

	location.Get("/subdistrict/:id", func(c *fiber.Ctx) error {
		return locationClient.GetSubdistrict(c, srv)
	})
	
	location.Put("/:type/id", func(c *fiber.Ctx) error {
		return locationClient.UpdateLocation(c, srv)
	})

	location.Delete("/:type/:id", func(c *fiber.Ctx) error {
		return locationClient.DeleteLocation(c, srv)
	})
	
}