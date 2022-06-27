package models

type LocationInput struct {
	ID			string `json:"id"`
	ParentID 	string `json:"parent_id"`
	Name 		string `json:"name"`
	Code 		string `json:"code"`
}