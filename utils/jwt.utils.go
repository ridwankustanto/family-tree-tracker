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