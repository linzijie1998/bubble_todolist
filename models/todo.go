package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Todo struct {
	gorm.Model
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}
