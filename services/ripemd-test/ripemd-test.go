package ripemdtest

import (
	"errors"

	"golang.org/x/crypto/ripemd160"
)
type Service interface{
	Encrypt(tokenString string) string
	Compare(a string, b string) error
}

type Services struct {
	Path string
}


func (s Services) Encrypt(tokenString string) string{
	h := ripemd160.New()
	h.Write([]byte(tokenString))

	return ""
}

func (s Services) Compare(a string, b string) error {
	if (a == b){
		return nil
	}else{
		return errors.New("String doesn't match!")
	}
}