package model

type User struct {
	UserID string `json:"userid,omitempty"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
}
