package controller

import (
	"ToDoOnGo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindAllToDo(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, model.Value)
}

func CreateToDo(ctx *gin.Context) {
	var newTodo model.ToDo

	if err := ctx.BindJSON(&newTodo); err != nil {
		return
	}

	valueTodo := append(model.Value, newTodo)
	ctx.IndentedJSON(http.StatusCreated, valueTodo)
}
