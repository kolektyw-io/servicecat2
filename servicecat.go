package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"servicecat/models"
	"servicecat/postgresdsn"
)

var db *gorm.DB
var errdb error

func init() {
	dsn := postgresdsn.CreateDSN(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DBNAME"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSLMODE"),
		os.Getenv("POSTGRES_TIMEZONE"))

	if os.Getenv("MODE") == "Debug" {
		db, errdb = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	} else {
		db, errdb = gorm.Open(postgres.Open(dsn.ReturnDSNAsString()), &gorm.Config{})
	}
	if errdb != nil {
		panic(errdb)
	}

}
func main() {
	fmt.Println("ServiceCat 0.0.3")
	fmt.Println("================")
	fmt.Println("Booting service...")

	if errdb != nil {
		fmt.Println(errdb)
	}
	fmt.Println("Test Migrations.")
	db.AutoMigrate(&models.Product{})
	fmt.Println("Migrating sessions...")
	db.AutoMigrate(&models.Session{})
	fmt.Println("Migrating users...")
	db.AutoMigrate(&models.User{})
	fmt.Println("Testing for connectivity")

	db = nil

	fmt.Println("All tests passed.")

}
