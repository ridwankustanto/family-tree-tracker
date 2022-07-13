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
	GetCountryByID(ctx context.Context, id string)(models.LocationReturn, error)
	GetProvinceByID(ctx context.Context, id string) (models.LocationReturn, error)
	GetCityByID(ctx context.Context, id string) (models.LocationReturn, error)
	GetDistrictByID(ctx context.Context, id string) (models.LocationReturn, error)
	GetSubdistrictByID(ctx context.Context, id string) (models.LocationReturn, error)
	GetAllCountry(ctx context.Context) ([]models.LocationReturn, error)
	GetAllProvince(ctx context.Context) ([]models.LocationReturn, error)
	GetAllCity(ctx context.Context) ([]models.LocationReturn, error)
	GetAllDistrict(ctx context.Context) ([]models.LocationReturn, error)
	GetAllSubdistrict(ctx context.Context) ([]models.LocationReturn, error)
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
		result, err:= r.db.ExecContext(ctx, "INSERT INTO city(id, province_id, name, code, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6)", 
		a.ID, a.ParentID, a.Name, a.Code, a.CreatedAt, a.UpdatedAt)
		log.Println(result)
		return "City Created!", err
	case "districts":
		result, err:= r.db.ExecContext(ctx, "INSERT INTO districts(id, city_id,  name, code, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6)", 
		a.ID, a.ParentID, a.Name, a.Code, a.CreatedAt, a.UpdatedAt)
		log.Println(result)
		return "District Created", err
	case "subdistricts":
		result, err:= r.db.ExecContext(ctx, "INSERT INTO subdistricts(id, district_id, name, code, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6)", 
		a.ID, a.ParentID, a.Name, a.Code, a.CreatedAt, a.UpdatedAt)
		log.Println(result)
		return "Subdistricts Created", err
	default:
		log.Println("Enter Location Type")
		return "", errors.New("Enter Location Type / 'type'")
	}
}

func StatementMasterByID(input string) (string, error) {
	switch input{
	case "country":
		return "SELECT id, $2, name, code from country where id=$1", nil
	case "provinces":
		return "SELECT id, $2, name, code from provinces where id=$1", nil 
	case "city":
		return "SELECT id, $2, name, code from city where id=$1", nil
	case "districts":
		return "SELECT id, $2, name, code from districts where id=$1", nil
	case "subdistricts":
		return "SELECT id, $2, name, code from subdistricts where id=$1", nil
	default:
		return "", errors.New("invalid Location Type")
	}
}



func (r postgresRepository) GetCountryByID(ctx context.Context, id string)(models.LocationReturn, error){
	country := new(models.LocationReturn)
	country.Type = "Country"
	err := r.db.QueryRow("SELECT id, name, code FROM country WHERE id=$1", id).Scan(&country.ID,&country.Name, &country.Code)
	if(err != nil){
		return *country, err
	}
	provinces, err := r.db.Query("SELECT id, name, code, country_id FROM provinces WHERE country_id=$1", id)
	if(err != nil){
		return *country, err
	}
	var province []models.Child
	
	for provinces.Next(){
		var arr models.Child
		if err := provinces.Scan(&arr.ID, &arr.Name, &arr.Code, &arr.ParentID); err != nil {
			return *country, err
		}
		arr.Type = "Province"
		province = append(province, arr)
	}
	
	country.Child = province
	if err = provinces.Err(); err != nil {
		return *country, err
	}
	
	return *country, nil
}

func (r postgresRepository) GetProvinceByID(ctx context.Context, id string) (models.LocationReturn, error){
	province:= new(models.LocationReturn)
	province.Type = "Province"
	err := r.db.QueryRow("SELECT id, country_id, name, code FROM provinces WHERE id=$1", id).Scan(&province.ID, &province.ParentID, &province.Name, &province.Code)
	if err != nil{
		return *province, err
	}
	cities, err := r.db.Query("SELECT id, province_id, name, code FROM city WHERE province_id=$1", id)
	if err != nil{
		return *province, err
	}
	var city []models.Child
	for cities.Next(){
		var arr models.Child
		if err := cities.Scan(&arr.ID, &arr.Name, &arr.Code, &arr.ParentID); err != nil {
			return *province, err
		}
		arr.Type = "City"
		city = append(city, arr)
	}
	
	province.Child = city
	if err = cities.Err(); err != nil {
		return *province, err
	}

	return *province, nil
}

func (r postgresRepository) GetCityByID(ctx context.Context, id string) (models.LocationReturn, error){
	city:= new(models.LocationReturn)
	city.Type = "City"

	err := r.db.QueryRow("SELECT id, province_id, name, code FROM city WHERE id=$1", id).Scan(&city.ID, &city.ParentID, &city.Name, &city.Code)
	if err != nil{
		return *city, err
	}
	districts, err := r.db.Query("SELECT id, city_id, name, code FROM district WHERE city_id=$1", id)
	if err != nil{
		return *city, err
	}
	var district []models.Child
	for districts.Next(){
		var arr models.Child
		if err := districts.Scan(&arr.ID, &arr.ParentID, &arr.Name, &arr.Code); err != nil {
			return *city, err
		}
		arr.Type = "District"
		district = append(district, arr)
	}
	
	city.Child = district
	if err = districts.Err(); err != nil {
		return *city, err
	}

	return *city, nil
}

func (r postgresRepository) GetDistrictByID(ctx context.Context, id string) (models.LocationReturn, error){
	district := new(models.LocationReturn)
	district.Type = "District"

	err := r.db.QueryRow("SELECT id, city_id, name, code FROM districts WHERE id=$1", id).Scan(&district.ID, &district.ParentID, &district.Name, &district.Code)
	if err != nil {
		return *district, err
	}

	subdistricts, err := r.db.Query("SELECT id, district_id, name, code, FROM district WHERE district_id=$1", id)
	if err != nil {
		return *district, err
	}
	var subdistrict []models.Child
	for subdistricts.Next(){
		var arr models.Child
		if err := subdistricts.Scan(&arr.ID, &arr.ParentID, &arr.Name, &arr.Code); err != nil {
			return *district, err
		}
		arr.Type = "Sub-District"
		subdistrict = append(subdistrict, arr)
	}

	district.Child = subdistrict
	if err = subdistricts.Err(); err != nil {
		return *district, err
	}
	return *district, nil
}

func (r postgresRepository) GetSubdistrictByID(ctx context.Context, id string) (models.LocationReturn, error){
	subdistrict := new(models.LocationReturn)
	subdistrict.Type = "Sub-District"
	err:= r.db.QueryRow("SELECT id, district_id, name, code FROM subdistricts WHERE id=$1", id).Scan(&subdistrict.ID, &subdistrict.Name, &subdistrict.Code, &subdistrict.ParentID)
	if err != nil{
		return *subdistrict, err
	}
	return *subdistrict, nil
}

func (r postgresRepository) GetAllCountry(ctx context.Context) ([]models.LocationReturn, error){
	var country []models.LocationReturn

	result, err:= r.db.Query("SELECT id, name, code FROM country") 
	if err != nil {
		return country, err
	}

	for result.Next(){
		var arr models.LocationReturn
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

func(r postgresRepository) GetAllProvince(ctx context.Context) ([]models.LocationReturn, error){
	var province []models.LocationReturn

	result, err := r.db.Query("SELECT id, country_id, name, code FROM provinces")
	if err != nil {
		return province, err
	}

	for result.Next(){
		var arr models.LocationReturn
		if err := result.Scan(&arr.ID, &arr.ParentID, &arr.Name, &arr.Code); err != nil {
			return province, err
		}
		arr.Type = "Province"
		province = append(province, arr)
	}
	if result.Err() != nil {
		return province, err
	}
	return province, nil
}

func(r postgresRepository) GetAllCity(ctx context.Context) ([]models.LocationReturn, error){
	var city []models.LocationReturn

	result, err := r.db.Query("SELECT id, province_id, name, code FROM city")
	if err != nil {
		return city, err
	}

	for result.Next(){
		var arr models.LocationReturn
		if err := result.Scan(&arr.ID, &arr.ParentID, &arr.Name, &arr.Code); err != nil {
			return city, err
		}
		arr.Type = "City"
		city = append(city, arr)
	}
	if result.Err() != nil {
		return city, err
	}
	return city, nil
}

func(r postgresRepository) GetAllDistrict(ctx context.Context) ([]models.LocationReturn, error){
	var district []models.LocationReturn

	result, err := r.db.Query("SELECT id, city_id, name, code FROM districts")
	if err != nil {
		return district, err
	}

	for result.Next(){
		var arr models.LocationReturn
		if err := result.Scan(&arr.ID, &arr.ParentID, &arr.Name, &arr.Code); err != nil {
			return district, err
		}
		arr.Type = "District"
		district = append(district, arr)
	}
	if result.Err() != nil {
		return district, err
	}
	return district, nil
}

func(r postgresRepository) GetAllSubdistrict(ctx context.Context) ([]models.LocationReturn, error){
	var subdistrict []models.LocationReturn

	result, err := r.db.Query("SELECT id, district_id, name, code FROM subdistricts")
	if err != nil {
		return subdistrict, err
	}

	for result.Next(){
		var arr models.LocationReturn
		if err := result.Scan(&arr.ID, &arr.ParentID, &arr.Name, &arr.Code); err != nil {
			return subdistrict, err
		}
		arr.Type = "Subdistrict"
		subdistrict = append(subdistrict, arr)
	}
	if result.Err() != nil {
		return subdistrict, err
	}
	return subdistrict, nil
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
