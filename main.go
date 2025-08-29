package main

import (
	_ "beshariq_tort_zakaz_bot/routers"
	"beshariq_tort_zakaz_bot/services/telegram"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {

	go telegram.Bot()

	beego.Run()
}
