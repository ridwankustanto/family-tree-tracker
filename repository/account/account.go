package account

import (
	"context"
	"database/sql"
	"errors"
	"log"
	_ "log"

	"github.com/ridwankustanto/family-tree-tracker/models"
)

type Repository interface {
	Close()
	Ping() error
	CreateAccount(ctx context.Context, a models.Account) error
	Authenticate(ctx context.Context, a models.AccountLogin) (models.AccountLogin ,error)
}

type postgresRepository struct {
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

// create acc 2nd step
func (r postgresRepository) CreateAccount(ctx context.Context, a models.Account) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO accounts(id, username, password, role, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6)",
		a.ID, a.Username, a.Password, a.Role, a.CreatedAt, a.UpdatedAt)
	return err
}

func (r postgresRepository) Authenticate(ctx context.Context, a models.AccountLogin) (models.AccountLogin, error) {
	err := r.db.QueryRow("SELECT id, username, password, role FROM accounts WHERE username=$1", a.Username).Scan(&a.ID, &a.Username, &a.Password, &a.Role)
	if(err != nil){
		return a, err
	}
	return a, nil
}

func (r postgresRepository) Location(ctx context.Context, a models.LocationInput) (string, error){
	// kalo pake switch case
	switch a.RequestType {
	case "country":
		_, err:= r.db.ExecContext(ctx, "INSERT INTO country(id, name, code, created_at, updated_at) VALUES($1, $2, $3, $4, $5)", 
		a.ID, a.Name, a.Code, a.CreatedAt, a.UpdatedAt)
		return "Country Inserted", err
	case "provinces":
		_, err:= r.db.ExecContext(ctx, "INSERT INTO provinces(id, country_id,  name, code, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6)", 
		a.ID, a.ParentID, a.Name, a.Code, a.CreatedAt, a.UpdatedAt)
		log.Println("provinces")
		return "Province Inserted", err
	case "city":
		_, err:= r.db.ExecContext(ctx, "INSERT INTO city(id, name, code, created_at, updated_at) VALUES($1, $2, $3, $4, $5)", 
		a.ID, a.ParentID, a.Name, a.Code, a.CreatedAt, a.UpdatedAt)
		log.Println("City")
		return "City Inserted", err
	case "districts":
		_, err:= r.db.ExecContext(ctx, "INSERT INTO districts(id, name, code, created_at, updated_at) VALUES($1, $2, $3, $4, $5)", 
		a.ID, a.ParentID, a.Name, a.Code, a.CreatedAt, a.UpdatedAt)
		log.Println("district")
		return "Districts Inserted", err
	case "subdistricts":
		_, err:= r.db.ExecContext(ctx, "INSERT INTO subdistricts(id, name, code, created_at, updated_at) VALUES($1, $2, $3, $4, $5)", 
		a.ID, a.ParentID, a.Name, a.Code, a.CreatedAt, a.UpdatedAt)
		log.Println("subdistricts")
		return "Sub Districts Inserted", err
	default:
		log.Println("Enter Location Type")
		return "", errors.New("Enter Location Type / 'type'")
	}


}


