package controllers

import beego "github.com/beego/beego/v2/server/web"

type ClientController struct {
	beego.Controller
}

func (c *ClientController) Get()  {
	c.Data["Website"] = "michal.dev"
	c.Data["Email"] = "michal@michal.dev"
	c.TplName = "index.html"
}
