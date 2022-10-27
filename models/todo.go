package models

type Todo struct {
	//gorm.Model
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}
