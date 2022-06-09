package models

type LogActivity struct {
	ID        int64  `json:"id"`
	PeopleID  int64  `json:"people_id"`
	Action    string `json:"action"`
	IPAddress string `json:"ip_address"`
	Client    string `json:"client"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
