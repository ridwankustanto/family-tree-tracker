package account

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ridwankustanto/family-tree-tracker/models"
	accountService "github.com/ridwankustanto/family-tree-tracker/services/account"
)

func CreateAccount(c *fiber.Ctx, svr accountService.Service) error {
	ctx := context.Background()

	// Query param
	account := new(models.Account)

	// Parse body request
	if err := c.BodyParser(account); err != nil {
		return c.Status(http.StatusBadGateway).JSON(map[string]interface{}{
			"error": true,
			"data":  nil,
		})
	}

	_, err := svr.CreateAccount(ctx, account.Username)
	if err != nil {
		return c.Status(http.StatusBadGateway).JSON(map[string]interface{}{
			"error": true,
			"data":  nil,
		})
	}

	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"error": false,
		"data":  "account created",
	})
}
