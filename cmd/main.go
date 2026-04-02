package main

import (
	"finance-backend/config"
	"finance-backend/models"
	"finance-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8080")
	config.ConnectDB()

	config.DB.AutoMigrate(&models.User{})
}
