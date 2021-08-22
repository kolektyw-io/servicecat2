package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/gorm"
	"html/template"
)

type ClientController struct {
	beego.Controller
	Db *gorm.DB
}

func (c *ClientController) Get() {
	c.StartSession()
	c.Data["Website"] = "michal.dev"
	c.Data["Email"] = "michal@michal.dev"
	c.TplName = "index.html"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["sesja"] = c.GetSession("user")

}

func (c *ClientController) Post() {
	c.Data["Website"] = "michal.dev"
	c.Data["Email"] = "michal@michal.dev"
	c.TplName = "after/index.html"
	c.Data["_xsrf"] = "Data properly sent"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
}
