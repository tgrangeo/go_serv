package models

import "gorm.io/gorm"

type Todolist struct {
	gorm.Model
	Name string `json:"name" gorm:"text;not null;default:null`
	Description string `json:"description" gorm:"text;not null;default:null`
}