package services

import (
	"finance-backend/config"
	"finance-backend/models"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user models.User) (models.User, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hashed)
	user.Role = models.Viewer
	user.Status = models.Active

	result := config.DB.Create(&user)
	return user, result.Error
}
