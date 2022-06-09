package models

type Account struct {
	ID        int64  `json:"id"`
	PeopleID  int64  `json:"people_id"`
	IsAdmin   bool   `json:"is_admin"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
