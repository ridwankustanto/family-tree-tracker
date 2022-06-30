package models

type LocationInput struct {
	ID			string `json:"id"`
	ParentID 	string `json:"parent_id"`
	Name 		string `json:"name"`
	Code 		string `json:"code"`
	RequestType	string `json:"type"`
	CreatedAt 	string `json:"created_at"`
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

type CountryReturn struct {
	Name		string `'json:"name"`
	Code 		string `json:"code"`
	Provinces 	[]Province `json:"provinces"`
}



