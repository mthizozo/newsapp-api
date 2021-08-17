package models


type NewsData struct {
	Email           string `json:"email",omitempty`
	Country           string `json:"country" validate:"required"`
}

type Error struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}