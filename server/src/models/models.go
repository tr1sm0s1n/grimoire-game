package models

type Certificate struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Course string `json:"course"`
	Grade  string `json:"grade"`
	Date   string `json:"date"`
}
