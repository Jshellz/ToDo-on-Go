package routes

import (
	"ToDo/app/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllToDo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.Todos)
}

func AddToDo(c *gin.Context) {
	var newToDo model.Todo

	if err := c.BindJSON(&newToDo); err != nil {
		return
	}

	model.Todos = append(model.Todos, newToDo)
	c.IndentedJSON(http.StatusCreated, newToDo)

}
