package controller

import (
	"ToDoOnGo/db_app"
	"ToDoOnGo/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

var key = "id"

func FindAllToDo(ctx *gin.Context) {
	var DBRead = db_app.DB
	var todo model.ToDo

	err := DBRead.Select(&todo, "SELECT * FROM `to_dos`")
	if err == nil {
		ctx.JSON(http.StatusOK, DBRead)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "problem in the query",
		})
	}

}

func CreateToDo(ctx *gin.Context) {
	var newTodo model.ToDo
	if err := ctx.BindJSON(&newTodo); err != nil {
		return
	}

	var DBCreate = db_app.DB
	DBCreate.Create(&newTodo)
	ctx.JSON(http.StatusCreated, DBCreate)
}

func DeleteToDo(_ *gin.Context) {
	var db gorm.DB
	DeleteToDos(&db)
}

func DeleteToDos(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var todo model.ToDo
		err := db.Where("id = ?", id).First(&todo).Error

		if err != nil {
			ctx.JSON(404, gin.H{"error": "todo not found"})
			return
		}

		db.Delete(&todo)
		ctx.JSON(200, gin.H{"data": "todo deleted successfully"})
	}
}
