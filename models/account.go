package models

type Account struct {
	ID        string `json:"id"`
	PeopleID  string `json:"people_id"`
	IsAdmin   bool   `json:"is_admin"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AccountLogin struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
}