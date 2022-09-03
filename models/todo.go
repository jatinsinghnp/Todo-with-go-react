package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ID    int    `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
