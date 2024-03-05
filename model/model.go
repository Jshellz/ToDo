package model

import "gorm.io/gorm"

type ToDo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
}
