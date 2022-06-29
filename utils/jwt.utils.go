package utils

import (
	// "errors"
	// "log"
	"os"
	"time"

	// "github.com/gofiber/fiber/v2"
	// jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt"
	"github.com/ridwankustanto/family-tree-tracker/models"
)

// func Restrict() fiber.Handler {
// 	return jwtware.New(jwtware.Config{
// 		SigningKey:   []byte(os.Getenv("SECRET_KEY")),
// 		ErrorHandler: jwtError,
// 	})
// }

// func Authorize(c *fiber.Ctx) error {
// 	user := c.Locals("user").(*jwt.Token)
// 	claims := user.Claims.(jwt.MapClaims)
// 	role := claims["role"].(string)
// 	if(role == "3"){
// 		return c.Status(fiber.StatusUnauthorized).
// 	JSON(fiber.Map{"status": "Unauthorized", "message": "You are not allowed to enter this area"})
// 	}

// 	// switch permission {
// 	// case "superadmin":
// 	// 	if(role != "1"){
// 	// 		return c.Status(fiber.StatusUnauthorized).
// 	// 	JSON(fiber.Map{"status": "Unauthorized", "message": "You are not allowed to enter this area"})
// 	// 	}
// 	// case "admin":
// 	// 	if(role == "3"){
// 	// 		return c.Status(fiber.StatusUnauthorized).
// 	// 	JSON(fiber.Map{"status": "Unauthorized", "message": "You are not allowed to enter this area"})
// 	// 	}
// 	// }

// 	return c.Next()
	
// }

// func jwtError(c *fiber.Ctx, err error) error {
// 	if err.Error() == "Missing or malformed JWT" {
// 		return c.Status(fiber.StatusBadRequest).
// 			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
// 	}
// 	return c.Status(fiber.StatusUnauthorized).
// 		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
// }

func GenerateToken(input *models.AccountLogin) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = input.ID
	claims["username"] = input.Username
	claims["role"] = input.Role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	
	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if(err != nil){
		return "", err
	}
	return signedToken, nil
}