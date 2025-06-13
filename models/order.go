package models

type Order struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Amount int    `json:"amount"`
}