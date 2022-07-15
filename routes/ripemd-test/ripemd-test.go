package ripemdtest

import (
	"github.com/gofiber/fiber/v2"
	ripemdService"github.com/ridwankustanto/family-tree-tracker/services/ripemd-test"
	ripemdClient "github.com/ridwankustanto/family-tree-tracker/clients/ripemd-test"
)

func Routes(api fiber.Router){
	ripemd := api.Group("req")
	srv := ripemdService.Services{}

	ripemd.Get("/request-client-secret/:client_id", func(c *fiber.Ctx) error {
		return ripemdClient.RequestClientSecret(c, srv)
	})

	ripemd.Post("/validate", func(c *fiber.Ctx) error {
		return ripemdClient.Validate(c, srv)
	})

}