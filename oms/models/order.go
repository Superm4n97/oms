package models

type Order struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Description string `json:"description"`
}
