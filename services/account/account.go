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
	Authenticate(ctx context.Context, account models.AccountLogin) (models.AccountLogin, string, error)
	CreateAccount(ctx context.Context, account models.Account) (*models.Account, error)
	BestowAccount(ctx context.Context, input models.Account)(*models.Account, error)
}
//sambungin ke Repository yang ada di account yang ngontrakin function yang ada
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
	account.CreatedAt = time.Now().Format(layout)
	account.UpdatedAt = time.Now().Format(layout)
	exist, err := s.repository.CheckUserExist(ctx, account.Username)
	if err != nil {
		return nil, err
	}else if exist == true {
		return nil, errors.New("This username already existed!")
	}

	if account.Role == "1" {
		exist, err := s.repository.CheckSuperAdmin(ctx);
		if err != nil{
			return nil, err
		}else if exist == true {
			return nil, errors.New("Role number one already existed!")
		}
	}else if account.Role == "2" {
		return nil, errors.New("Super admin's permission required!")
	}

	log.Println(account)

	if err := s.repository.CreateAccount(ctx, account); err != nil {
		return nil, err
	}
	return &account, nil
}

func (s service) Authenticate(ctx context.Context, account models.AccountLogin) (models.AccountLogin, string,  error){
	log.Println("input: ", account)

	acc, err := s.repository.Authenticate(ctx, account); 
	if err != nil {
		return acc, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(account.Password))
	if err != nil{
		return acc, "", errors.New("Invalid username or password")
	}
	
	token, err := utils.GenerateToken(&acc)
	// log.Println(err)
	
	// log.Println(&account)
	return acc, token, nil
}

func (s service) BestowAccount(ctx context.Context, input models.Account)(*models.Account, error){
	layout := "2006-01-02T15:04:05-0700"
	input.UpdatedAt = time.Now().Format(layout)
	// if input.Role == "1" {
	// 	return nil, errors.New("There may be only one superadmin")
	// }
	_, err := s.repository.BestowAccount(ctx, input)
	if err != nil {
		return nil, err
	}
	return &input, nil
}




