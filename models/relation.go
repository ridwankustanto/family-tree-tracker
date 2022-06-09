package models

type Relation struct {
	ID           string `json:"id"`
	PeopleID     string `json:"people_id"`
	WithPeopleID string `json:"with_people_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
