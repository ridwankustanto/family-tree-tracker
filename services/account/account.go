package account

import (
	"context"

	"github.com/ridwankustanto/family-tree-tracker/models"
	"github.com/ridwankustanto/family-tree-tracker/repository/account"
)

type Service interface {
	CreateAccount(ctx context.Context, name string) (*models.Account, error)
}

type service struct {
	repository account.Repository
}

func NewService(r account.Repository) Service {
	return &service{r}
}

func (s service) CreateAccount(ctx context.Context, name string) (*models.Account, error) {
	account := &models.Account{
		ID:       "asdasdas",
		PeopleID: "A-3-0-0-0-0-0",
		Username: "Admin",
	}
	if err := s.repository.CreateAccount(ctx, *account); err != nil {
		return nil, err
	}
	return account, nil
}
