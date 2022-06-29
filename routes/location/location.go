package location

import(
	"database/sql"
	"github.com/gofiber/fiber/v2"
	locationClient "github.com/ridwankustanto/family-tree-tracker/clients/location"
	locationRepo "github.com/ridwankustanto/family-tree-tracker/repository/location"
	locationService "github.com/ridwankustanto/family-tree-tracker/services/location"
)

func Routes(api fiber.Router, db *sql.DB) {
	repo:= locationRepo.NewPostgresRepository(db)
	
	srv:= locationService.NewService(repo)

	location := api.Group("location")

	location.Post("/add", func(c *fiber.Ctx) error {
		return locationClient.CreateLocation(c, srv)
	})

	location.Get("/locations", func(c *fiber.Ctx) error {
		return nil
	})
	location.Post("/edit", func(c *fiber.Ctx) error {
		return nil
	})
}