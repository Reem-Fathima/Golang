package auth

import (
	"net/http"
	"task-management-system/internal/models"
	"task-management-system/pkg/db"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, hasAuth := c.Request.BasicAuth()
		if !hasAuth {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			return
		}

		var user models.User
		if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
