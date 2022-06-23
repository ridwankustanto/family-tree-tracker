package models

type Account struct {
	ID        string `json:"id"`
	PeopleID  string `json:"people_id"`
	Role      string `json:"role"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AccountLogin struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
}