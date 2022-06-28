package location

import(
	"context"
	"database/sql"
	"errors"
	"log"
	"github.com/ridwankustanto/family-tree-tracker/models"


)

type Repository interface {
	Close()
	Ping() error
	CreateLocation(ctx context.Context, a models.LocationInput) (string, error)
}

type postgresRepository struct{
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &postgresRepository{db}
}

func (r *postgresRepository) Close() {
	r.db.Close()
}

func (r *postgresRepository) Ping() error {
	return r.db.Ping()
}

func (r postgresRepository) CreateLocation(ctx context.Context, a models.LocationInput) (string, error){
	// kalo pake switch case
	switch a.RequestType {
	case "country":
		result, err:= r.db.ExecContext(ctx, "INSERT INTO country(id, name, code, created_at, updated_at) VALUES($1, $2, $3, $4, $5)", 
		a.ID, a.Name, a.Code, a.CreatedAt, a.UpdatedAt)
		log.Println(result)
		return "Country Created!", err
	case "provinces":
		result, err:= r.db.ExecContext(ctx, "INSERT INTO provinces(id, country_id,  name, code, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6)", 
		a.ID, a.ParentID, a.Name, a.Code, a.CreatedAt, a.UpdatedAt)
		log.Println(result)
		return "Provinces Created!", err
	case "city":
		result, err:= r.db.ExecContext(ctx, "INSERT INTO city(id, provinces_id, name, code, created_at, updated_at) VALUES($1, $2, $3, $4, $5)", 
		a.ID, a.ParentID, a.Name, a.Code, a.CreatedAt, a.UpdatedAt)
		log.Println(result)
		return "City Created!", err
	case "districts":
		result, err:= r.db.ExecContext(ctx, "INSERT INTO districts(id, city_id,  name, code, created_at, updated_at) VALUES($1, $2, $3, $4, $5)", 
		a.ID, a.ParentID, a.Name, a.Code, a.CreatedAt, a.UpdatedAt)
		log.Println(result)
		return "District Created", err
	case "subdistricts":
		result, err:= r.db.ExecContext(ctx, "INSERT INTO subdistricts(id, districts_id, name, code, created_at, updated_at) VALUES($1, $2, $3, $4, $5)", 
		a.ID, a.ParentID, a.Name, a.Code, a.CreatedAt, a.UpdatedAt)
		log.Println(result)
		return "Subdistricts Created", err
	default:
		log.Println("Enter Location Type")
		return "", errors.New("Enter Location Type / 'type'")
	}
}