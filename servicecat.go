package main

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"

	"servicecat/models"
	_ "servicecat/routers"
)

var db *gorm.DB
var errdb error

func main() {
	fmt.Println("ServiceCat 0.0.2")
	fmt.Println("================")
	fmt.Println("Booting service...")
	dsn := CreateDSN(os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_SSLMODE"), os.Getenv("POSTGRES_TZ"))
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
	db.Create(&models.Product{Code: "D42", Price: 100})

	// Read
	var product models.Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(&models.Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)

	db = nil

	fmt.Println("All tests passed.")

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "sessions"
	beego.Run()

}
