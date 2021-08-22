package routers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"servicecat/controllers"
)

func init() {
	dsn := "host=ec2-54-220-53-223.eu-west-1.compute.amazonaws.com user=rgzkkffjhwswsf password=0b8c58b6d586348f45c6ee6dd6a14030805df81035109842b0b12d4ad8a1059a dbname=d55nlr4gefst4q port=5432 sslmode=require TimeZone=Asia/Shanghai"
	db, errdb := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errdb != nil {
		fmt.Println(errdb)
	}
	beego.Router("/backend", &controllers.MainController{Db: db})
	beego.Router("/api", &controllers.MainController{Db: db})
	beego.Router("/user", &controllers.UserLoginController{Db: db})
	beego.Router("/", &controllers.HomeController{Db: db})
}
