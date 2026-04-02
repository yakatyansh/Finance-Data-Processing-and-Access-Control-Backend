package models

import "time"

type Role string

const (
	Admin   Role = "ADMIN"
	Viewer  Role = "VIEWER"
	Analyst Role = "ANALYST"
)

type Status string

const (
	Active   Status = "ACTIVE"
	Inactive Status = "INACTIVE"
)

type User struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	Role      Role
	Status    Status
	CreatedAt time.Time
}
