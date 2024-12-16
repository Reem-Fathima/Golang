package main

import (
	"task-management-system/internal/auth"
	"task-management-system/internal/handlers"
	"task-management-system/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	r := gin.Default()
	r.POST("/register", handlers.Register)

	authGroup := r.Group("/tasks", auth.AuthMiddleware())
	authGroup.POST("/create", handlers.CreateTask)
	authGroup.GET("/", handlers.GetTasks)

	r.Run(":8080")
}
