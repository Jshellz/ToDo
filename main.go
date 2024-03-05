package main

import (
	"ToDoOnGo/controller"
	"ToDoOnGo/db_app"
	"github.com/gin-gonic/gin"
)

var host = "localhost:8000"
var panicRunHost = "failed to run"

func init() {
	db_app.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.GET("/", controller.FindAllToDo)
	r.POST("/createToDo", controller.CreateToDo)

	if err := r.Run(host); err != nil {
		panic(panicRunHost)
	}
}
