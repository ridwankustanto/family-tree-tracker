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
	GetCountry(ctx context.Context, id string)(models.CountryReturn, error)
	GetProvince(ctx context.Context, id string) (models.ProvinceReturn, error)
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
	log.Println(a)
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

func (r postgresRepository) GetCountry(ctx context.Context, id string)(models.CountryReturn, error){
	country := new(models.CountryReturn)
	err := r.db.QueryRow("SELECT id, name, code FROM country WHERE id=$1", id).Scan(&country.ID,&country.Name, &country.Code)
	if(err != nil){
		return *country, err
	}
	provinces, err := r.db.Query("SELECT id, name, code, country_id FROM provinces WHERE country_id=$1", id)
	if(err != nil){
		return *country, err
	}
	var province []models.Province
	
	for provinces.Next(){
		var arr models.Province
		if err := provinces.Scan(&arr.ID, &arr.CountryID, &arr.Name, &arr.Code); err != nil {
			return *country, err
		}
		province = append(province, arr)
	}
	
	country.Provinces = province
	if err = provinces.Err(); err != nil {
		return *country, err
	}
	
	return *country, nil
}

func (r postgresRepository) GetProvince(ctx context.Context, id string) (models.ProvinceReturn, error){
	province:= new(models.ProvinceReturn)

	err := r.db.QueryRow("SELECT id, name, code, country_id FROM provinces WHERE country_id=$1", id).Scan(&province.ID, &province.CountryID, &province.Name, &province.Code)
	if err != nil{
		return *province,nil
	}
	cities, err := r.db.Query("SELECT id, name, code, province_id FROM city WHERE province_id=$1", id)
	if err != nil{
		return *province,nil
	}
	var city []models.City
	for cities.Next(){
		var arr models.City
		if err := cities.Scan(&arr.ID, &arr.ProvinceID, &arr.Name, &arr.Code); err != nil {
			return *province, err
		}
		city = append(city, arr)
	}
	
	province.City = city
	if err = cities.Err(); err != nil {
		return *province, err
	}

	return *province, nil
}

