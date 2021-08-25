package controllers

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"servicecat/models"
)

var user *models.User

func Login(db *gorm.DB, username string, password string) bool {
	result := db.First(&user, "name = ?", username)
	if result.Error != nil {
		fmt.Println("Nie ma takiego usera, login zwraca false", username)
		return false
	}
	fmt.Println("user to", user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	fmt.Println(err)
	if err == nil {
		fmt.Println("login zwraca true")
		return true
	} else {
		fmt.Println("login zwraca false")
		return false
	}
}

func Register(db *gorm.DB, username string, password string) bool {
	return false
}
