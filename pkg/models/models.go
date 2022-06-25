package models

type User struct {
	ID      int    `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Date    string `json:"date"`
	Country string `json:"country"`
}