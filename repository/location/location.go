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
	GetAllCountry(ctx context.Context) ([]models.Country, error)
	GetProvince(ctx context.Context, id string) (models.ProvinceReturn, error)
	GetCity(ctx context.Context, id string) (models.CityReturn, error)
	GetDistrict(ctx context.Context, id string) (models.DistrictReturn, error)
	GetSubdistrict(ctx context.Context, id string) (models.Subdistrict, error)
	UpdateLocation(ctx context.Context, input models.LocationInput) (sql.Result, error)
	DeleteLocation(ctx context.Context, input models.LocationInput)(sql.Result, error)
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
	switch a.Type {
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

func (r postgresRepository) GetAllCountry(ctx context.Context) ([]models.Country, error){
	var country []models.Country

	result, err:= r.db.Query("SELECT id, name, code FROM country") 
	if err != nil {
		return country, err
	}

	for result.Next(){
		var arr models.Country
		if err := result.Scan(&arr.ID, &arr.Name, &arr.Code); err != nil {
			return country, err
		}
		country = append(country, arr)
	}

	if result.Err() != nil {
		return country, err
	}

	return country, nil
}

func (r postgresRepository) GetProvince(ctx context.Context, id string) (models.ProvinceReturn, error){
	province:= new(models.ProvinceReturn)

	err := r.db.QueryRow("SELECT id, name, code, country_id FROM provinces WHERE id=$1", id).Scan(&province.ID, &province.CountryID, &province.Name, &province.Code)
	if err != nil{
		return *province, err
	}
	cities, err := r.db.Query("SELECT id, name, code, province_id FROM city WHERE province_id=$1", id)
	if err != nil{
		return *province, err
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

func (r postgresRepository) GetCity(ctx context.Context, id string) (models.CityReturn, error){
	city:= new(models.CityReturn)

	err := r.db.QueryRow("SELECT id, name, code, province_id FROM city WHERE id=$1", id).Scan(&city.ID, &city.ProvinceID, &city.Name, &city.Code)
	if err != nil{
		return *city, err
	}
	districts, err := r.db.Query("SELECT id, name, code, city_id FROM district WHERE city_id=$1", id)
	if err != nil{
		return *city, err
	}
	var district []models.District
	for districts.Next(){
		var arr models.District
		if err := districts.Scan(&arr.ID, &arr.CityID, &arr.Name, &arr.Code); err != nil {
			return *city, err
		}
		district = append(district, arr)
	}
	
	city.District = district
	if err = districts.Err(); err != nil {
		return *city, err
	}

	return *city, nil
}

func (r postgresRepository) GetDistrict(ctx context.Context, id string) (models.DistrictReturn, error){
	district := new(models.DistrictReturn)

	err := r.db.QueryRow("SELECT id, name, code, city_id FROM districts WHERE id=$1", id).Scan(&district.ID, &district.CityID, &district.Name, &district.Code)
	if err != nil {
		return *district, err
	}

	subdistricts, err := r.db.Query("SELECT id, name, code, district_id FROM district WHERE district_id=$1", id)
	if err != nil {
		return *district, err
	}
	var subdistrict []models.Subdistrict
	for subdistricts.Next(){
		var arr models.Subdistrict
		if err := subdistricts.Scan(&arr.ID, &arr.DistrictID, &arr.Name, &arr.Code); err != nil {
			return *district, err
		}
		subdistrict = append(subdistrict, arr)
	}

	district.Subdistrict = subdistrict
	if err = subdistricts.Err(); err != nil {
		return *district, err
	}
	return *district, nil
}

func (r postgresRepository) GetSubdistrict(ctx context.Context, id string) (models.Subdistrict, error){
	subdistrict := new(models.Subdistrict)

	err:= r.db.QueryRow("SELECT id, name, code, district_id FROM subdistricts WHERE id=$1", id).Scan(&subdistrict.ID)
	if err != nil{
		return *subdistrict, err
	}
	return *subdistrict, nil
}

func (r postgresRepository) UpdateLocation(ctx context.Context, input models.LocationInput) (sql.Result, error) {
	switch input.Type {
	case "country":
		result, err := r.db.ExecContext(ctx, "UPDATE country SET name=$2, code=$3, updated_at=$4 WHERE id=$1", input.ID, input.Name, input.Code, input.UpdatedAt)
		log.Println(result)
		return result, err
	case "provinces":
		result, err := r.db.ExecContext(ctx, "UPDATE provinces SET name=$2, code=$3, country_id=$4, updated_at=$5 WHERE id=$1", input.ID, input.Name, input.Code, input.ParentID, input.UpdatedAt)
		log.Println(result)
		return result, err
	case "city":
		result, err := r.db.ExecContext(ctx, "UPDATE city SET name=$2, code=$3, province_id=$4, updated_at=$5 WHERE id=$1", input.ID, input.Name, input.Code, input.ParentID, input.UpdatedAt)
		log.Println(result)
		return result, err
	case "districts":
		result, err := r.db.ExecContext(ctx, "UPDATE districts SET name=$2, code=$3, city_id=$4, updated_at=$5 WHERE id=$1", input.ID, input.Name, input.Code, input.ParentID, input.UpdatedAt)
		log.Println(result)
		return result, err
	case "subdistricts":
		result, err := r.db.ExecContext(ctx, "UPDATE subdistricts SET name=$2, code=$3, district_id=$4, updated_at=$5 WHERE id=$1", input.ID, input.Name, input.Code, input.ParentID, input.UpdatedAt)
		log.Println(result)
		return result, err
	default:
		log.Println("Enter Location Type!")
		return nil, errors.New("Enter Location Type!")
	}
}

func (r postgresRepository) DeleteLocation(ctx context.Context, input models.LocationInput)(sql.Result, error){
	switch input.Type{
	case "country":
		result, err := r.db.ExecContext(ctx, "DELETE FROM country WHERE id=$1", input.ID)
		log.Println(result)
		return result, err
		
	case "provinces":
		result, err := r.db.ExecContext(ctx, "DELETE FROM country WHERE id=$1", input.ID)
		log.Println(result)
		return result, err
	
	case "city":
		result, err := r.db.ExecContext(ctx, "DELETE FROM country WHERE id=$1", input.ID)
		log.Println(result)
		return result, err
	
	case "districts":
		result, err := r.db.ExecContext(ctx, "DELETE FROM country WHERE id=$1", input.ID)
		log.Println(result)
		return result, err
	
	case "subdistricts":
		result, err := r.db.ExecContext(ctx, "DELETE FROM country WHERE id=$1", input.ID)
		log.Println(result)
		return result, err 
	default:
		log.Println("Enter Location Type!")
		return nil, errors.New("Enter Location Type!")
	}
}
