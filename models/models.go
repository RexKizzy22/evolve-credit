package models

import (
	"database/sql"
	// "time"
)

type Models struct {
	DB DBModel
}

func NewModel(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

type User struct {
	ID      int    `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Date    string `json:"date"`
	Country string `json:"country"`
}
