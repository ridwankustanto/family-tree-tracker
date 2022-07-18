package ripemdtest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ridwankustanto/family-tree-tracker/clients"
	"github.com/ridwankustanto/family-tree-tracker/models"

	ripemdService "github.com/ridwankustanto/family-tree-tracker/services/ripemd-test"
)

func RequestClientSecret(c *fiber.Ctx, srv ripemdService.Service) error {
	client_id := c.Params("client_id")
	token := srv.Prepare(client_id, "").Encrypt()
	return c.Status(http.StatusOK).JSON(clients.Response{
		Error: false,
		Message: fmt.Sprintf("Here's your token"),
		Data:    *&token,
	})
}

func Validate(c *fiber.Ctx, srv ripemdService.Service) error {
	Authorization := c.Get("Authorization")
	date := c.Params("date")
	input := new(models.ClientInput)
	if err := c.BodyParser(input); err != nil {
		log.Println("c.BodyParser(input)", err)
		return c.Status(http.StatusBadGateway).JSON(clients.Response{
			Error:        true,
			DebugMessage: err.Error(),
			Message:      clients.ErrSomethingWentWrong,
		})
	}
	token := srv.Prepare(input.ClientID, date).Encrypt()
	Auth := srv.Sanitize(Authorization, "Bearer ")
	err := srv.Compare(Auth, token)
	if err != nil {
		log.Println("srv.Compare()", err)
		return c.Status(http.StatusBadGateway).JSON(clients.Response{
			Error:        true,
			DebugMessage: err.Error(),
			Message:      "Your Client ID doesn't match Authorization",
		})
	}
	return c.Status(http.StatusOK).JSON(clients.Response{
		Error: false,
		Message: fmt.Sprintf("ACCESS GRANTED!"),
	})
	
}