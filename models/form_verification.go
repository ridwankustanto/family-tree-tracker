package models

type FormVerification struct {
	ID          int64  `json:"id"`
	GeneratedBy int64  `json:"generated_by"`
	UsedBy      int64  `json:"Used_by"`
	IsUsed      bool   `json:"is_used"`
	Code        string `json:"code"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
