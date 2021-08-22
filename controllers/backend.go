package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/gorm"
)

type MainController struct {
	beego.Controller
	Db *gorm.DB
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}
