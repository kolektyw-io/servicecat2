package controllers

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"html/template"
	"servicecat/models"
)

type UserLoginController struct {
	UserAwareController
	Db *gorm.DB
}

func (c *UserLoginController) Get() {
	c.StartSession()
	c.TplName = "login.html"
	c.Data["xsrf"] = template.HTML(c.XSRFFormHTML())
}

func (c *UserLoginController) Post() {
	c.StartSession()
	var muser *models.User
	user := c.GetString("user")
	fmt.Println(user)

	if Login(c.Db, c.GetString("user"), c.GetString("password")) {
		// Generate session and enable it in memory
		// Save session cookie
		token := uuid.New().String()
		c.SetSession("token", token)
		result := c.Db.First(&muser, "name = ?", user)
		if result.Error != nil {
			panic(result.Error)
		}
		ses := &models.Session{Token: token}
		c.Db.Create(ses)
		muser.Sessions = append(muser.Sessions, *ses)
		c.Db.Updates(&muser)

		c.Redirect("/", 302)
	} else {
		c.Redirect("/zle", 302)
	}
}
