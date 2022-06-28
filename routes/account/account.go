package account

import (
	// "log"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	accountClient "github.com/ridwankustanto/family-tree-tracker/clients/account"
	accountRepo "github.com/ridwankustanto/family-tree-tracker/repository/account"
	accountService "github.com/ridwankustanto/family-tree-tracker/services/account"
)

func Routes(api fiber.Router, db *sql.DB) {
	//NewPostgresRepository digunakan untuk initizialie datbase biar di repo gk usah initDB
	repo := accountRepo.NewPostgresRepository(db)

	//taking in repo because service is equivalent to controller while repo is equivallent to model 
	srv := accountService.NewService(repo)
	// Clients is the body parser, the upper part | Service is the controller | Repo is the model that connects themselves to database
	account := api.Group("account")

	account.Post("/register", func(c *fiber.Ctx) error {
		return accountClient.CreateAccount(c, srv)
	})

	account.Post("/login", func(c *fiber.Ctx) error {
		return accountClient.Authenticate(c, srv)
	})

	account.Post("/location", func(c * fiber.Ctx) error {
		return nil
	})
}
