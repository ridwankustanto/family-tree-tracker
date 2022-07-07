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
	"github.com/ridwankustanto/family-tree-tracker/utils/middlewares"

	
)

func CreateLocation(c *fiber.Ctx, srv locationService.Service) error {
	role := middlewares.Authorize(c); 
	if role !=nil{
		log.Println("Error You are not Authorized: ", role)
		return middlewares.GetOut(c, role.Error())
	}
	ctx := context.Background()
	location := new(models.LocationInput)
	location.Type = c.Params("type")
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
		Message: fmt.Sprintf("%v%v Created!", strings.ToUpper(string(location.Type[0])), string(location.Type[1:])),
		Data:    *location,
	})

}

func GetLocationByID(c * fiber.Ctx, srv locationService.Service)error{
	ctx := context.Background()
	id := c.Params("id")
	request_type := c.Params("type")

	result, err := srv.GetLocationByID(ctx, id, request_type)
	if err != nil {
		log.Println("srv.GetLocationByID()", err)
		return c.Status(http.StatusBadGateway).JSON(clients.Response{
			Error:        true,
			DebugMessage: err.Error(),
			Message:      clients.ErrBadGateway,
		})
	}

	return c.Status(http.StatusOK).JSON(clients.Response{
		Error: false,
		// Message: message,
		Message: fmt.Sprintf("Showing %v with ID: %v", strings.Title(string(result.Type)), result.ID),
		Data:    *&result,
	})
}

func GetAllLocation(c *fiber.Ctx, srv locationService.Service) error {
	ctx := context.Background()
	request_type := c.Params("type")
	result, err := srv.GetAllLocation(ctx, request_type)
	if err != nil {
		log.Println("srv.GetAllLocation()", err)
		return c.Status(http.StatusBadGateway).JSON(clients.Response{
			Error:        true,
			DebugMessage: err.Error(),
			Message:      clients.ErrBadGateway,
		})
	}

	return c.Status(http.StatusOK).JSON(clients.Response{
		Error: false,
		// Message: message,
		Message: fmt.Sprintf("Showing %v", strings.Title(string(request_type))),
		Data:    *&result,
	})
}

// func GetAllCountry(c *fiber.Ctx, srv locationService.Service) error {
// 	ctx := context.Background()

// 	result, err := srv.GetAllCountry(ctx)
// 	if err != nil {
// 		log.Println("srv.GetAllCountry()", err)
// 		return c.Status(http.StatusBadGateway).JSON(clients.Response{
// 			Error:        true,
// 			DebugMessage: err.Error(),
// 			Message:      clients.ErrBadGateway,
// 		})
// 	}

// 	return c.Status(http.StatusOK).JSON(clients.Response{
// 		Error: false,
// 		// Message: message,
// 		Message: fmt.Sprintf("Showing All Country"),
// 		Data:    *&result,
// 	})
// }

func UpdateLocation(c *fiber.Ctx, srv locationService.Service) error {
	role := middlewares.Authorize(c); 
	if role !=nil{
		log.Println("Error You are not Authorized: ", role)
		return middlewares.GetOut(c, role.Error())
	}
	ctx := context.Background()
	location := new(models.LocationInput)
	location.ID = c.Params("id")
	location.Type = c.Params("type")

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
	location, err = srv.UpdateLocation(ctx, *location)
	if err != nil {
		log.Println("srv.UpdateLocation()", err)
		return c.Status(http.StatusBadGateway).JSON(clients.Response{
			Error:        true,
			DebugMessage: err.Error(),
			Message:      clients.ErrBadGateway,
		})
	}

	return c.Status(http.StatusOK).JSON(clients.Response{
		Error: false,
		// Message: message,
		Message: fmt.Sprintf("%v%v Updated!", strings.ToUpper(string(location.Type[0])), string(location.Type[1:])),
		Data:    *location,
	})
}

func DeleteLocation(c *fiber.Ctx, srv locationService.Service) error {
	role := middlewares.Authorize(c); 
	if role !=nil{
		log.Println("Error You are not Authorized: ", role)
		return middlewares.GetOut(c, role.Error())
	}
	ctx := context.Background()
	location := new(models.LocationInput)
	location.ID = c.Params("id")
	location.Type = c.Params("type")

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
	location, err = srv.DeleteLocation(ctx, *location)
	if err != nil {
		log.Println("srv.DeleteLocation()", err)
		return c.Status(http.StatusBadGateway).JSON(clients.Response{
			Error:        true,
			DebugMessage: err.Error(),
			Message:      clients.ErrBadGateway,
		})
	}

	return c.Status(http.StatusOK).JSON(clients.Response{
		Error: false,
		// Message: message,
		Message: fmt.Sprintf("%v%v Deleted!", strings.ToUpper(string(location.Type[0])), string(location.Type[1:])),
		Data:    *location,
	})
}