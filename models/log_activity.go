package models

type LogActivity struct {
	ID        string `json:"id"`
	PeopleID  string `json:"people_id"`
	Action    string `json:"action"`
	IPAddress string `json:"ip_address"`
	Client    string `json:"client"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
