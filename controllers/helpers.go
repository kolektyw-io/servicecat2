package controllers

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"servicecat/models"
)

var user *models.User

func Login(db *gorm.DB, username string, password string) bool {
	result := db.First(&user, "name = ?", username)
	if result.Error != nil {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	return err == nil
}

func Register(db *gorm.DB, username string, password string) bool {
	return false
}
