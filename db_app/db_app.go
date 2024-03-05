package db_app

import (
	"ToDoOnGo/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	panicConnect = "failed to connect database"
	dsn          = "user:pass@tcp/db_app"
)

var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func ConnectToDB() {
	if err != nil {
		panic(panicConnect)
	}

	if err := db.AutoMigrate(&model.ToDo{}); err != nil {
		fmt.Printf("failed to migrate %v", err)
	}
}

func AddNewTodo() {
	db.Create(&model.ToDo{Title: "test", Description: "test", Complete: false})
}
