package main

import (
	"golang.org/x/crypto/bcrypt"
	"servicecat/models"
)

var user *models.User

func Login(username string, password string) bool {
	result := db.First(&user, "username = ?", username)
	if result.Error != nil {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	return err == nil
}
