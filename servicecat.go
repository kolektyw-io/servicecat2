package main

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"servicecat/models"
	"servicecat/postgresdsn"
	_ "servicecat/routers"
)

var db *gorm.DB
var errdb error

func main() {
	fmt.Println("ServiceCat 0.0.3")
	fmt.Println("================")
	fmt.Println("Booting service...")
	dsn := postgresdsn.CreateDSN("localhost", "servicecat", "servicecat", "servicecat", "5432", "disable", "Europe/Warsaw")
	db, errdb = gorm.Open(postgres.Open(dsn.ReturnDSNAsString()), &gorm.Config{})
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

	// Read


	// Delete - delete product
	//db.Delete(&product, 1)

	db = nil

	fmt.Println("All tests passed.")

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "sessions"
	beego.Run()
}
