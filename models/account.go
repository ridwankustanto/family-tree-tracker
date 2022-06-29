package models

type Account struct {
	ID        string `json:"id"`
	Role      string `json:"role"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AccountLogin struct {
	ID        string `json:"id"`
	Role      string `json:"role"`
	Username  string `json:"username"`
	Password  string `json:"password,omitempty"`
}

