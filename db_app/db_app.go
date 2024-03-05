package db_app

import (
	"ToDoOnGo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dsn = "user:pass@tcp/db_app"
)

var DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func ConnectToDB() {
	if err != nil {
		return
	}

	err := DB.AutoMigrate(&model.ToDo{})
	if err != nil {
		return
	}
}
