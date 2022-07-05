package account

import (
	"context"
	"database/sql"

	"github.com/ridwankustanto/family-tree-tracker/models"
)

type Repository interface {
	Close()
	Ping() error
	CheckUserExist(ctx context.Context, username string)(bool, error)
	CreateAccount(ctx context.Context, a models.Account) error
	Authenticate(ctx context.Context, a models.AccountLogin) (models.AccountLogin ,error)
	CheckSuperAdmin(ctx context.Context) (bool, error)
}
//tied contract by returning postgresRepository as repository and call postgresRepository on each struct
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

func (r postgresRepository) CheckUserExist(ctx context.Context, username string)(bool, error){
	err := r.db.QueryRow("SELECT id FROM accounts WHERE username=$1", username).Scan(&username)
	if err != nil {
		if err != sql.ErrNoRows{
			return false, err
		}
		return false, nil
	}
	return true, nil
}
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

func (r postgresRepository) CheckSuperAdmin(ctx context.Context) (bool, error) {
	var username string
	err := r.db.QueryRow("SELECT id FROM accounts WHERE role=$1", "1").Scan(&username)
	if err != nil {
		if err != sql.ErrNoRows{
			return false, err
		}
		return false, nil
	}
	return true, nil
}

