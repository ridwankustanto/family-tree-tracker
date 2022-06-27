package account

import (
	"context"
	"database/sql"
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
	err := r.db.QueryRow("SELECT id, username, password FROM accounts WHERE username=$1", a.Username).Scan(&a.ID, &a.Username, &a.Password)
	// log.Println(a, "repository")
	if(err != nil){
		return a, err
	}
	return a, nil
}

func (r postgresRepository) Location(ctx context.Context) (string, error){
	
	return "", nil

}


