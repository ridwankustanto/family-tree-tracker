package location

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/ridwankustanto/family-tree-tracker/models"
	"github.com/ridwankustanto/family-tree-tracker/repository/location"
	"github.com/ridwankustanto/family-tree-tracker/utils"
)
type Service interface{
	CreateLocation(ctx context.Context, input models.LocationInput)(*models.LocationInput, string, error)
	GetCountry(ctx context.Context, id string) (models.CountryReturn, error)
}

type service struct{
	repository location.Repository
}

func NewService(r location.Repository) Service {
	return &service{r}
}

func (s service) CreateLocation(ctx context.Context, input models.LocationInput) (*models.LocationInput, string, error){
	layout := "2006-01-02T15:04:05-0700"

	input.ID = utils.FormatUUID(uuid.New().String())
	input.CreatedAt = time.Now().Format(layout)
	input.UpdatedAt = time.Now().Format(layout)

	message, err := s.repository.CreateLocation(ctx, input)
	if err != nil {
		return &input, "", err
	}
	return &input, message, nil
}

func (s service) GetCountry(ctx context.Context, id string) (models.CountryReturn, error){
	result, err := s.repository.GetCountry(ctx, id)
	if err != nil {
		return result, err
	}
	return result, nil
}