package account

import (
	"context"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ridwankustanto/family-tree-tracker/clients"
	"github.com/ridwankustanto/family-tree-tracker/models"
	"github.com/ridwankustanto/family-tree-tracker/utils/middlewares"
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
		log.Println("svr.CreateAccount()", err)
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
	data, token, err := srv.Authenticate(ctx, *account)
	if err != nil {
		log.Println("svr.Authenticate", err)
		return c.Status(http.StatusBadGateway).JSON(clients.Response{
			Error:        true,
			DebugMessage: err.Error(),
			Message:      clients.ErrBadGateway,
		})
	}
	data.Password = ""
	if data.Role == "1" {
		return c.Status(http.StatusOK).JSON(clients.ResponseLogin{
			Message: "logged in",
			Token: token,
		})
	}
	// Call generate token function
	return c.Status(http.StatusOK).JSON(clients.ResponseLogin{
		Message: "logged in",
		Token: token,
		Data:    &data,
	})
}

func BestowAccount(c *fiber.Ctx, srv accountService.Service) error {
	role := middlewares.SuperAdmin(c);
	if role != nil {
		log.Println("Error You are not Authorized: ", role)
		return middlewares.GetOut(c, role.Error())
	}
	ctx := context.Background()
	
	account := new(models.Account)
	if err := c.BodyParser(account); err != nil {
		log.Println("c.BodyParser(account)", err)
		return c.Status(http.StatusBadGateway).JSON(clients.Response{
			Error:        true,
			DebugMessage: err.Error(),
			Message:      clients.ErrSomethingWentWrong,
		})
	}

	var err error
	account, err = srv.BestowAccount(ctx, *account)
	if err != nil {
		log.Println("svr.BestowAccount()", err)
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


