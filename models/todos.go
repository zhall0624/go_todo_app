package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID          uint
	Title       string `json:"title" xml:"title" form:"title" query:"title"`
	Description string `json:"description" xml:"description" form:"description" query:"description"`
	Completed   bool   `json:"Completed" xml:"Completed" form:"Completed" query:"Completed" gorm:"default:false"`
}
