package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/gorm"
	"servicecat/models"
)

type UserAwareController struct {
	beego.Controller
	Db   *gorm.DB
	ses  *models.Session
	user *models.User
}

func (c *UserAwareController) Prepare() {
	c.StartSession()
	token := c.GetSession("token")
	tx := c.Db.First(&c.ses, "token = ?", token)

	if tx.Error != nil {
		c.ses = nil
		c.user = nil
	} else {
		c.Db.First(&c.user, "id = ?", c.ses.UserID)
	}

}
