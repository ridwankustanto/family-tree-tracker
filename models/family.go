package models

type Family struct {
	ID           int64  `json:"id"`
	PeopleID     int64  `json:"people_id"`
	WithPeopleID int64  `json:"with_people_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
