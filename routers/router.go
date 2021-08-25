package routers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"servicecat/controllers"
	"servicecat/postgresdsn"
)

func init() {
	dsn := postgresdsn.CreateDSN("localhost", "servicecat", "servicecat", "servicecat", "5432", "disable", "Europe/Warsaw")
	db, errdb := gorm.Open(postgres.Open(dsn.ReturnDSNAsString()), &gorm.Config{})
	if errdb != nil {
		fmt.Println(errdb)
	}

	beego.Router("/backend", &controllers.MainController{Db: db})
	beego.Router("/api", &controllers.MainController{Db: db})
	beego.Router("/user/register", &controllers.UserRegisterController{Db: db})
	beego.Router("/user/recover", &controllers.UserRecoverController{Db: db})
	beego.Router("/user", &controllers.UserLoginController{Db: db})
	beego.Router("/", &controllers.HomeController{Db: db})
}
