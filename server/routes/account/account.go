package account

import (
	"log"

	"github.com/gofiber/fiber/v2"
	accountClient "github.com/ridwankustanto/family-tree-tracker/clients/account"
	accountRepo "github.com/ridwankustanto/family-tree-tracker/repository/account"
	accountService "github.com/ridwankustanto/family-tree-tracker/services/account"
	"github.com/ridwankustanto/family-tree-tracker/utils/database"
)

func Routes(api fiber.Router) {

	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	repo := accountRepo.NewPostgresRepository(db)
	srv := accountService.NewService(repo)

	api.Post("/account", func(c *fiber.Ctx) error {
		return accountClient.CreateAccount(c, srv)
	})
}
