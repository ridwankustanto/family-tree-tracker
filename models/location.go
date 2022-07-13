package models

type LocationInput struct {
	ID			string `json:"id"`
	ParentID 	string `json:"parent_id,omitempty"`
	Name 		string `json:"name"`
	Code 		string `json:"code"`
	Type		string `json:"type"`
	CreatedAt 	string `json:"created_at,omitempty"`
	UpdatedAt 	string `json:"updated_at,omitempty"`
}

type Child struct {
	ID			string `json:"id"`
	Type		string `json:"type"`
	ParentID	string `json:"parent_id"`
	Name 		string `json:"name"`
	Code 		string `json:"code"`
}

type LocationReturn struct {
	ID			string `json:"id"`
	Type		string `json:"type,omitempty"`
	ParentID	string `json:"parent_id,omitempty"`	
	Name		string `'json:"name"`
	Code 		string `json:"code"`
	Child 		[]Child `json:"child,omitempty"`
}

