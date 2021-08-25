package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html/template"
	"servicecat/models"
)

type UserLoginController struct {
	beego.Controller
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
	status := Login(c.Db, c.GetString("user"), c.GetString("password"))
	fmt.Println("status logowania", status)
	if status == true {
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

type UserRegisterController struct {
	beego.Controller
	Db *gorm.DB
}

func (c *UserRegisterController) Get() {
	c.StartSession()
	c.TplName = "register.html"
	c.Data["xsrf"] = template.HTML(c.XSRFFormHTML())
}

func (c *UserRegisterController) Post() {
	c.StartSession()
	var user *models.User
	username := c.GetString("user")
	password := c.GetString("password")
	_ = c.GetString("password-repeat")

	result := c.Db.First(&user, "name = ?", username)
	if result.Error != nil {
		fmt.Println("Nie ma takiego usera, login zwraca false", username)
		fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
		if err != nil {
			return
		}
		c.Db.Create(&models.User{Name: username, Password: string(fromPassword)})
	}
	c.Redirect("/user/login", 302)

}

type UserRecoverController struct {
	beego.Controller
	Db *gorm.DB
}

func (c *UserRecoverController) Get() {
	c.StartSession()
	c.TplName = "register.html"
	c.Data["xsrf"] = template.HTML(c.XSRFFormHTML())
}

func (c *UserRecoverController) Post() {
	c.StartSession()
	var muser *models.User
	user := c.GetString("user")
	fmt.Println(user)
	status := Login(c.Db, c.GetString("user"), c.GetString("password"))
	fmt.Println("status logowania", status)
	if status == true {
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
