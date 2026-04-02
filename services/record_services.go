package services

import (
	"finance-backend/config"
	"finance-backend/models"
)

func CreateRecord(record models.Record) (models.Record, error) {
	result := config.DB.Create(&record)
	return record, result.Error
}

func GetRecords() ([]models.Record, error) {
	var records []models.Record
	result := config.DB.Find(&records)
	return records, result.Error
}

func GetFilteredRecords(typeFilter string, category string) ([]models.Record, error) {
	var records []models.Record

	query := config.DB

	if typeFilter != "" {
		query = query.Where("type = ?", typeFilter)
	}

	if category != "" {
		query = query.Where("category = ?", category)
	}

	query.Find(&records)

	return records, nil
}
