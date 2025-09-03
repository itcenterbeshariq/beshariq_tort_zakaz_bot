package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {

	admin := c.GetSession("admin")
	if admin == nil {
		c.Redirect("/login", 302)
		return
	}
}
