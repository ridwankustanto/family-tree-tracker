package models

type Event struct {
	ID          string        `json:"id"`
	PeopleID    string        `json:"people_id"`
	Category    EventCategory `json:"category"`
	Description string        `json:"description"`
	Date        string        `json:"date"`
	CreatedAt   string        `json:"created_at"`
	UpdatedAt   string        `json:"updated_at"`
}

type EventCategory string

const (
	Birth     EventCategory = "Birth"
	Death     EventCategory = "Death"
	Marriage  EventCategory = "Marriage"
	Graduated EventCategory = "Graduated"
)
