package account

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/ridwankustanto/family-tree-tracker/models"
	"github.com/ridwankustanto/family-tree-tracker/repository/account"
	"github.com/ridwankustanto/family-tree-tracker/utils"
	"golang.org/x/crypto/bcrypt"
)

// interface 
type Service interface {
	Authenticate(ctx context.Context, account models.AccountLogin) (string, error)
	CreateAccount(ctx context.Context, account models.Account) (*models.Account, error)
}

type service struct {
	repository account.Repository
}
//New service tied to Service interface karena method dibawah bind ke service
func NewService(r account.Repository) Service {
	return &service{r}
}
//Create acc 1st step
func (s service) CreateAccount(ctx context.Context, account models.Account) (*models.Account, error) {
	layout := "2006-01-02T15:04:05-0700"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), 14)

	account.ID = utils.FormatUUID(uuid.New().String())
	account.Password = string(hashedPassword)
	account.PeopleID = "dbf6d29d25144d3aa54d44ad36c27b60"
	account.CreatedAt = time.Now().Format(layout)
	account.UpdatedAt = time.Now().Format(layout)

	log.Println(account)

	if err := s.repository.CreateAccount(ctx, account); err != nil {
		return nil, err
	}
	return &account, nil
}

func (s service) Authenticate(ctx context.Context, account models.AccountLogin) (string, error){
	log.Println("input: ", account)

	x, err := s.repository.Authenticate(ctx, account); 
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(x.Password), []byte(account.Password))
	if err != nil{
		return "", errors.New("Invalid username or password")
	}
	
	token, err := utils.GenerateToken(&x, "SECRET_TUNNEL")
	// log.Println(err)
	
	// log.Println(&account)
	return token, nil
}



