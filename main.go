package main

import (
	"ToDoOnGo/controller"
	"ToDoOnGo/db_app"
	"fmt"
	"github.com/gin-gonic/gin"
)

var host = "localhost:8000"

func init() {
	db_app.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.GET("/", controller.FindAllToDo)
	r.POST("/createToDo", controller.CreateToDo)
	r.DELETE("/deleteToDo/:id", controller.DeleteToDo)

	if err := r.Run(host); err != nil {
		fmt.Println("Ok")
	}
}
