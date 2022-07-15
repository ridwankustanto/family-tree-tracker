package ripemdtest

import (
	"encoding/hex"
	"errors"
	"log"
	"strings"

	"golang.org/x/crypto/ripemd160"
)
type Service interface{
	Encrypt(tokenString string) string
	Compare(Authorization string, b string) error
	Sanitize(bearer string, filter string) string
}

type Services struct {
	Path string
}


func (s Services) Encrypt(tokenString string) string{
	h := ripemd160.New()
	h.Write([]byte(tokenString))
	hashBytes := h.Sum(nil)

	return hex.EncodeToString(hashBytes)
}

func (s Services) Sanitize(bearer string, filter string) string{
	str := strings.Replace(bearer, filter, "", -1)
	return str
}

func (s Services) Compare(Authorization string, b string) error {
	log.Println(Authorization, b)
	if Authorization == b{
		return nil
	}else{
		return errors.New("String doesn't match!")
	}
}