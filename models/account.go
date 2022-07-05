package models

type Account struct {
	ID        string `json:"id"`
	Role      string `json:"role"`
	Username  string `json:"username"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type AccountLogin struct {
	ID        string `json:"id"`
	Role      string `json:"role"`
	Username  string `json:"username"`
	Password  string `json:"password,omitempty"`
}

