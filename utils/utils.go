package utils

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ridwankustanto/family-tree-tracker/models"
)

func ForeverSleep(d time.Duration, f func(int) error) {
	for i := 0; ; i++ {
		err := f(i)
		if err == nil {
			return
		}
		time.Sleep(d)
	}
}

func FormatUUID(s string) string {
	return strings.ReplaceAll(s, "-", "")
}

func GenerateToken(input *models.AccountLogin, key string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = input.ID
	claims["username"] = input.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	
	signedToken, err := token.SignedString([]byte(key))

	if(err != nil){
		return "", err
	}
	return signedToken, nil
}
