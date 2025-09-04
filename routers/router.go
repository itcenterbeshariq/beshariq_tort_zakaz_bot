package routers

import (
	"beshariq_tort_zakaz_bot/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.HomeController{}, "get:Home")
	beego.Router("/login", &controllers.AuthController{}, "get:LoginForm;post:Login")
	beego.Router("/logout", &controllers.AuthController{}, "get:Logout")

}
