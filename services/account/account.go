package account

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/ridwankustanto/family-tree-tracker/models"
	"github.com/ridwankustanto/family-tree-tracker/repository/account"
	"github.com/ridwankustanto/family-tree-tracker/utils"
)

type Service interface {
	CreateAccount(ctx context.Context, account models.Account) (*models.Account, error)
}

type service struct {
	repository account.Repository
}

func NewService(r account.Repository) Service {
	return &service{r}
}

func (s service) CreateAccount(ctx context.Context, account models.Account) (*models.Account, error) {
	layout := "2006-01-02T15:04:05-0700"

	account.ID = utils.FormatUUID(uuid.New().String())
	account.PeopleID = "dbf6d29d25144d3aa54d44ad36c27b60"
	account.CreatedAt = time.Now().Format(layout)
	account.UpdatedAt = time.Now().Format(layout)

	log.Println(account)

	if err := s.repository.CreateAccount(ctx, account); err != nil {
		return nil, err
	}
	return &account, nil
}
