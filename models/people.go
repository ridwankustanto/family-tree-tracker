package models

type People struct {
	ID             string `json:"id"`
	FirstName      string `json:"firstname"`
	LastName       string `json:"lastname"`
	Nickname       string `json:"nickname"`
	PlaceOfBirth   string `json:"place_of_birth"`
	DateOfBirth    string `json:"date_of_birth"`
	Gender         string `json:"gender"`
	Religion       string `json:"religion"`
	FamilyID       string `json:"family_id"`
	OriginAddress  string `json:"origin_address"`
	CurrentAddress string `json:"current_address"`
	Country        string `json:"country"`
	Province       string `json:"province"`
	City           string `json:"city"`
	District       string `json:"disctrict"`
	Subdistrict    string `json:"subdisctrict"`
	MobilePhone    string `json:"mobile_phone"`
	Phone          string `json:"phone"`
	Photo          string `json:"photo"`
	IsOrigin       string `json:"is_origin"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
