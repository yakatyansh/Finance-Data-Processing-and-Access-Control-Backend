package services

import (
	"finance-backend/config"
	"finance-backend/models"
)

type Summary struct {
	TotalIncome  float64            `json:"totalIncome"`
	TotalExpense float64            `json:"totalExpense"`
	NetBalance   float64            `json:"netBalance"`
	CategoryWise map[string]float64 `json:"categoryWise"`
}

func GetSummary(userID string) (Summary, error) {
	var records []models.Record

	config.DB.Where("user_id = ?", userID).Find(&records)

	summary := Summary{
		CategoryWise: make(map[string]float64),
	}

	for _, r := range records {
		if r.Type == models.Income {
			summary.TotalIncome += r.Amount
		} else {
			summary.TotalExpense += r.Amount
		}

		summary.CategoryWise[r.Category] += r.Amount
	}

	summary.NetBalance = summary.TotalIncome - summary.TotalExpense

	return summary, nil
}
