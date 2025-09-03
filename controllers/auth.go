package controllers

import (
	"beshariq_tort_zakaz_bot/database"
	models "beshariq_tort_zakaz_bot/static"
	beego "github.com/beego/beego/v2/server/web"
)

type AuthController struct {
	beego.Controller
}

func (c *AuthController) LoginForm() {

	admin := c.GetSession("admin")
	if admin != nil {
		c.Redirect("/", 302)
		return
	}

	c.TplName = "login.html"
}
func (c *AuthController) Login() {
	c.TplName = "login.html"

	role := c.GetString("Role")
	password := c.GetString("Password")

	var user models.User
	if err := database.DB.Where("role = ?", role).First(&user).Error; err != nil {
		c.Data["Error"] = "Foydalanuvchi topilmadi"
		return
	}

	if password != user.Password {
		c.Data["Error"] = "Parol noto'g'ri"
		return
	}

	_ = c.SetSession("admin", user.Role)

	c.Redirect("/", 302)
}
