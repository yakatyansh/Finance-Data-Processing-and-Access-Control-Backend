package controllers

import (
	"finance-backend/services"

	"github.com/gin-gonic/gin"
)

func GetDashboard(c *gin.Context) {
	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	data, err := services.GetSummary(userID.(string))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}
