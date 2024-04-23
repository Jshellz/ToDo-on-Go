package main

import (
	"ToDo/app/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", routes.GetAllToDo)
	r.POST("/createToDo", routes.AddToDo)

	if run := r.Run(":8080"); run != nil {
		fmt.Printf("Failed to start server: %v\n", run.Error())
	}
}
