package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"servicecat/controllers"
)

func init() {
    beego.Router("/backend", &controllers.MainController{})
    beego.Router("/", &controllers.ClientController{})
}
