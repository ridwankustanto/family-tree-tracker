package ripemdtest

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ridwankustanto/family-tree-tracker/clients"

	ripemdService "github.com/ridwankustanto/family-tree-tracker/services/ripemd-test"
)

func RequestClientSecret(c *fiber.Ctx, srv ripemdService.Service) error {
	client_id := c.Params("client_id")
	token := srv.Encrypt(client_id)
	return c.Status(http.StatusOK).JSON(clients.Response{
		Error: false,
		Message: fmt.Sprintf("Here's your token"),
		Data:    *&token,
	})
}