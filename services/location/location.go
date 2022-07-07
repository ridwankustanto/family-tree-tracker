package location

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/ridwankustanto/family-tree-tracker/models"
	"github.com/ridwankustanto/family-tree-tracker/repository/location"
	"github.com/ridwankustanto/family-tree-tracker/utils"
)
type Service interface{
	CreateLocation(ctx context.Context, input models.LocationInput)(*models.LocationInput, string, error)
	GetLocationByID(ctx context.Context, id string, request_type string) (*models.LocationReturn, error)
	GetAllLocation(ctx context.Context, request_type string) (*[]models.LocationReturn, error)
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

func(s service) GetLocationByID(ctx context.Context, id string, request_type string) (*models.LocationReturn, error) {
	switch request_type{
	case "country":
		result, err := s.repository.GetCountryByID(ctx, id)
		if err!= nil {
			return nil, err
		}
		return &result, nil

	case "provinces":
		result, err := s.repository.GetProvinceByID(ctx, id)
		if err!= nil {
			return nil, err
		}
		return &result, nil

	case "city":
		result, err := s.repository.GetCityByID(ctx, id)
		if err!= nil {
			return nil, err
		}
		return &result, nil
	
	case "districts":
		result, err := s.repository.GetDistrictByID(ctx, id)
		if err!= nil {
			return nil, err
		}
		return &result, nil
	
	case "subdistricts":
		result, err := s.repository.GetSubdistrictByID(ctx, id)
		if err!= nil {
			return nil, err
		}
		return &result, nil

	default:
		return nil, errors.New("Enter Location Type / 'type'")
	}
} 

func (s service) GetAllLocation(ctx context.Context, request_type string) (*[]models.LocationReturn, error) {
	switch request_type{
	case "country":
		result, err := s.repository.GetAllCountry(ctx)
		if err!= nil {
			return nil, err
		}
		return &result, nil

	case "provinces":
		result, err := s.repository.GetAllProvince(ctx)
		if err!= nil {
			return nil, err
		}
		return &result, nil

	case "city":
		result, err := s.repository.GetAllCity(ctx)
		if err!= nil {
			return nil, err
		}
		return &result, nil
	
	case "districts":
		result, err := s.repository.GetAllDistrict(ctx)
		if err!= nil {
			return nil, err
		}
		return &result, nil
	
	case "subdistricts":
		result, err := s.repository.GetAllSubdistrict(ctx)
		if err!= nil {
			return nil, err
		}
		return &result, nil

	default:
		return nil, errors.New("Enter Location Type / 'type'")
	}
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