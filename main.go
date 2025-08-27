package main

import (
	_ "beshariq_tort_zakaz_bot/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

