package account

import (
	"context"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	accountService "github.com/ridwankustanto/family-tree-tracker/services/account"
)

func CreateAccount(c *fiber.Ctx, svr accountService.Service) error {
	ctx := context.Background()

	// Query param
	// account := new(models.Account)

	// // Parse body request
	// if err := c.BodyParser(account); err != nil {
	// 	log.Println("c.BodyParser(account)", err)
	// 	return c.Status(http.StatusBadGateway).JSON(map[string]interface{}{
	// 		"error": true,
	// 		"debug": err.Error(),
	// 		"data":  nil,
	// 	})
	// }

	_, err := svr.CreateAccount(ctx, "account.Username")
	if err != nil {
		log.Println("svr.CreateAccount", err)
		return c.Status(http.StatusBadGateway).JSON(map[string]interface{}{
			"error": true,
			"debug": err.Error(),
			"data":  nil,
		})
	}

	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"error": false,
		"data":  "account created",
	})
}
