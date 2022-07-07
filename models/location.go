package models

type LocationInput struct {
	ID			string `json:"id"`
	ParentID 	string `json:"parent_id,omitempty"`
	Name 		string `json:"name"`
	Code 		string `json:"code"`
	Type	string `json:"type"`
	CreatedAt 	string `json:"created_at,omitempty"`
	UpdatedAt 	string `json:"updated_at"`
}

type Country struct {
	ID			string `json:"id"`
	Name 		string `json:"name"`
	Code 		string `json:"code"`
}

type Province struct {
	ID			string `json:"id"`
	CountryID	string `json:"country_id"`
	Name 		string `json:"name"`
	Code 		string `json:"code"`
}

type City struct {
	ID			string `json:"id"`
	ProvinceID	string `json:"province_id"`
	Name 		string `json:"name"`
	Code 		string `json:"code"`
}

type District struct {
	ID			string `json:"id"`
	CityID	string `json:"city_id"`
	Name 		string `json:"name"`
	Code 		string `json:"code"`	
}

type Subdistrict struct {
	ID			string `json:"id"`
	DistrictID	string `json:"district_id"`
	Name 		string `json:"name"`
	Code 		string `json:"code"`
}

type CountryReturn struct {
	ID			string `json:"id"`
	Name		string `'json:"name"`
	Code 		string `json:"code"`
	Provinces 	[]Province `json:"provinces"`
}

type ProvinceReturn struct {
	ID			string `json:"id"`
	CountryID	string `json:"province_id"`	
	Name		string `'json:"name"`
	Code 		string `json:"code"`
	City 		[]City `json:"city"`
}

type CityReturn struct {
	ID			string `json:"id"`
	ProvinceID	string `json:"province_id"`	
	Name		string `'json:"name"`
	Code 		string `json:"code"`
	District 	[]District `json:"districts"`
}

type DistrictReturn struct {
	ID			string `json:"id"`
	CityID		string `json:"city_id"`
	Name		string `'json:"name"`
	Code 		string `json:"code"`
	Subdistrict []Subdistrict `json:"Subdistricts"`
}

