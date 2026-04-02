package controllers

import (
	"finance-backend/models"
	"finance-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRecord(c *gin.Context) {
	var record models.Record

	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := services.CreateRecord(record)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, created)
}

func GetRecords(c *gin.Context) {
	typeFilter := c.Query("type")
	category := c.Query("category")

	records, _ := services.GetFilteredRecords(typeFilter, category)

	c.JSON(http.StatusOK, records)
}
