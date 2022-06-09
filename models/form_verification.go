package models

type FormVerification struct {
	ID          string `json:"id"`
	GeneratedBy string `json:"generated_by"`
	UsedBy      string `json:"Used_by"`
	IsUsed      bool   `json:"is_used"`
	Code        string `json:"code"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
