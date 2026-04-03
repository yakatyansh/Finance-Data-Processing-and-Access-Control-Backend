package services

import (
	"errors"
	"finance-backend/config"
	"finance-backend/models"

	"github.com/golang-jwt/jwt/v5"
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

func GenerateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"role": string(user.Role),
	})

	return token.SignedString([]byte("secret"))
}

func LoginUser(email string, password string) (models.User, error) {
	var user models.User

	result := config.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return user, errors.New("user not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("invalid credentials")
	}

	return user, nil
}
