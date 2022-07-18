package ripemdtest

import (
	"encoding/hex"
	"errors"
	"fmt"
	
	"strings"
	"time"

	"golang.org/x/crypto/ripemd160"
)
type Service interface{
	Prepare(tokenString string, date string) Services
	Encrypt() string
	Compare(Authorization string, b string) error
	Sanitize(bearer string, filter string) string
}

type Services struct {
	Result string
}

func (s Services) Prepare(tokenString string, date string) Services {
	format := "20060102"
	if date == "" {
		date = time.Now().Format(format)
	}else {
		tempDate, _ := time.Parse(format, date)
		date = tempDate.Format(format)
	}

	token := fmt.Sprintf("RIPEMD160(%v%v)", tokenString, date)
	s.Result = token
	return s
}

func (s Services) Encrypt() string{
	h := ripemd160.New()
	h.Write([]byte(s.Result))
	hashBytes := h.Sum(nil)

	return hex.EncodeToString(hashBytes)
}

func (s Services) Sanitize(bearer string, filter string) string{
	str := strings.Replace(bearer, filter, "", -1)
	return str
}

func (s Services) Compare(Authorization string, b string) error {
	// log.Println(Authorization, b)
	if Authorization == b{
		return nil
	}else{
		return errors.New("String doesn't match!")
	}
}