package models

import "time"

type RecordType string

const (
	Income  RecordType = "INCOME"
	Expense RecordType = "EXPENSE"
)

type Record struct {
	ID        string `gorm:"primaryKey"`
	UserID    string
	Amount    float64
	Type      RecordType
	Category  string
	Date      time.Time
	Note      string
	CreatedAt time.Time
}
