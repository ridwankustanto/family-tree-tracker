package location

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ridwankustanto/family-tree-tracker/clients"
	"github.com/ridwankustanto/family-tree-tracker/models"
	locationService "github.com/ridwankustanto/family-tree-tracker/services/location"
)

func CreateLocation(c *fiber.Ctx, srv locationService.Service) error {
	ctx := context.Background()

	location := new(models.LocationInput)

	if err := c.BodyParser(location); err != nil {
		log.Println("c.BodyParser(location)", err)
		return c.Status(http.StatusBadGateway).JSON(clients.Response{
			Error:        true,
			DebugMessage: err.Error(),
			Message:      clients.ErrSomethingWentWrong,
		})
	}
	var err error
	// var message string
	location, _, err = srv.CreateLocation(ctx, *location)
	if err != nil {
		log.Println("srv.CreateLocation()", err)
		return c.Status(http.StatusBadGateway).JSON(clients.Response{
			Error:        true,
			DebugMessage: err.Error(),
			Message:      clients.ErrBadGateway,
		})
	}

	return c.Status(http.StatusOK).JSON(clients.Response{
		Error: false,
		// Message: message,
		Message: fmt.Sprintf("%v%v Created!", strings.ToUpper(string(location.RequestType[0])), string(location.RequestType[1:])),
		Data:    *location,
	})

}
