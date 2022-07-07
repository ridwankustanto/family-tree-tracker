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
	GetCountry(ctx context.Context, id string) (*models.CountryReturn, error)
	GetAllCountry(ctx context.Context)(*[]models.Country, error)
	GetProvince(ctx context.Context, id string) (*models.ProvinceReturn, error)
	GetCity(ctx context.Context, id string) (*models.CityReturn, error)
	GetDistrict(ctx context.Context, id string) (*models.DistrictReturn, error)
	GetSubdistrict(ctx context.Context, id string) (*models.Subdistrict, error)
	UpdateLocation(ctx context.Context, input models.LocationInput) (*models.LocationInput, error)
	DeleteLocation(ctx context.Context, input models.LocationInput) (*models.LocationInput, error)
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
		return nil, "", err
	}
	return &input, message, nil
}

func (s service) GetCountry(ctx context.Context, id string) (*models.CountryReturn, error){
	result, err := s.repository.GetCountry(ctx, id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s service) GetAllCountry(ctx context.Context)(*[]models.Country, error){
	result, err := s.repository.GetAllCountry(ctx)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s service) GetProvince(ctx context.Context, id string) (*models.ProvinceReturn, error){
	result, err := s.repository.GetProvince(ctx, id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s service) GetCity(ctx context.Context, id string) (*models.CityReturn, error){
	result, err := s.repository.GetCity(ctx, id)
	if err != nil {
		return nil, err
	}
	return &result, nil	
}

func (s service) GetDistrict(ctx context.Context, id string) (*models.DistrictReturn, error){
	result, err := s.repository.GetDistrict(ctx, id)
	if err != nil {
		return nil, err
	}
	return &result, nil	
}

func (s service) GetSubdistrict(ctx context.Context, id string) (*models.Subdistrict, error){
	result, err := s.repository.GetSubdistrict(ctx, id)
	if err != nil {
		return nil, err
	}
	return &result, nil	
}

func (s service) UpdateLocation(ctx context.Context, input models.LocationInput) (*models.LocationInput, error){
	layout := "2006-01-02T15:04:05-0700"
	input.UpdatedAt = time.Now().Format(layout)

	_, err := s.repository.UpdateLocation(ctx, input)
	if err != nil {
		return nil, err
	}
	return &input, nil
}

func (s service) DeleteLocation(ctx context.Context, input models.LocationInput) (*models.LocationInput, error){
	_, err:= s.repository.DeleteLocation(ctx, input)
	if err != nil {
		return nil, err
	}
	return &input, nil
}