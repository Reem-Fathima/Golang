package handlers

import (
	"net/http"
	"task-management-system/internal/models"
	"task-management-system/pkg/db"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	user, _ := c.Get("user")
	var req models.Task
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.UserID = user.(models.User).ID
	if err := db.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
}

func GetTasks(c *gin.Context) {
	user, _ := c.Get("user")
	var tasks []models.Task
	if err := db.DB.Where("user_id = ?", user.(models.User).ID).Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
