package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("userRole", "ADMIN")
		c.Set("userID", "123")

		c.Next()
	}
}
