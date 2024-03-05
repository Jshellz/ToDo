package model

import "gorm.io/gorm"

type ToDo struct {
	gorm.Model
	Title       string
	Description string
	Complete    bool
}

// Value test
var Value = []ToDo{
	{Title: "test", Description: "test", Complete: false},
}
