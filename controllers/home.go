package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/gorm"
	"servicecat/models"
)

type HomeController struct {
	beego.Controller
	Db   *gorm.DB
	ses  *models.Session
	user *models.User
}

func (c *HomeController) Prepare() {
	// Start session
	c.StartSession()
	token := c.GetSession("token")
	tx := c.Db.First(&c.ses, "token = ?", token)

	if tx.Error != nil {
		c.Redirect("/user", 302)
	}
}

func (c *HomeController) Get() {
	//
	//c.Data["sesja"] = c.user.Name
	c.TplName = "index.html"

	// If no session data is available move to login page

	// If session is existent check for rights and transfer user to proper place (for staff or for clients)

	// If is_staff or Is admin - to portal for agents
	// If not, - to portal for clients

}
