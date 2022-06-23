package account

import (
	"context"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ridwankustanto/family-tree-tracker/clients"
	"github.com/ridwankustanto/family-tree-tracker/models"
	accountService "github.com/ridwankustanto/family-tree-tracker/services/account"
)
// Body Parser nya 
func CreateAccount(c *fiber.Ctx, srv accountService.Service) error {
	ctx := context.Background()

	// json raw
	account := new(models.Account)

	// // Parse body request
	if err := c.BodyParser(account); err != nil {
		log.Println("c.BodyParser(account)", err)
		return c.Status(http.StatusBadGateway).JSON(clients.Response{
			Error:        true,
			DebugMessage: err.Error(),
			Message:      clients.ErrSomethingWentWrong,
		})
	}

	// Validate

	var err error
	account, err = srv.CreateAccount(ctx, *account)
	if err != nil {
		log.Println("svr.CreateAccount", err)
		return c.Status(http.StatusBadGateway).JSON(clients.Response{
			Error:        true,
			DebugMessage: err.Error(),
			Message:      clients.ErrBadGateway,
		})
	}

	return c.Status(http.StatusOK).JSON(clients.Response{
		Error:   false,
		Message: clients.CreateSuccess,
		Data:    account,
	})
}

func Authenticate(c *fiber.Ctx, srv accountService.Service)error{
	ctx := context.Background()

	// json raw
	account := new(models.AccountLogin)

	// // Parse body request
	if err := c.BodyParser(account); err != nil {
		log.Println("c.BodyParser(account)", err)
		return c.Status(http.StatusBadGateway).JSON(clients.Response{
			Error:        true,
			DebugMessage: err.Error(),
			Message:      clients.ErrSomethingWentWrong,
		})
	}

	// Validate

	var err error
	account, err = srv.Authenticate(ctx, *account)
	if err != nil {
		log.Println("svr.Authenticate", err)
		return c.Status(http.StatusBadGateway).JSON(clients.Response{
			Error:        true,
			DebugMessage: err.Error(),
			Message:      clients.ErrBadGateway,
		})
	}
	// Call generate token function
	return c.Status(http.StatusOK).JSON(clients.Response{
		Error:   false,
		Message: "logged in",
		Data:    account,
	})
}
