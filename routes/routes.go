package routes

import (
	"finance-backend/controllers"
	"finance-backend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
	}
	record := r.Group("/records")
	record.Use(middleware.AuthMiddleware())
	{
		record.POST("/", middleware.Authorize("ADMIN"), controllers.CreateRecord)
		record.GET("/", middleware.Authorize("ADMIN", "ANALYST", "VIEWER"), controllers.GetRecords)
	}
	dashboard := r.Group("/dashboard")
	dashboard.Use(middleware.AuthMiddleware())
	{
		dashboard.GET("/summary", controllers.GetDashboard)
	}
}
