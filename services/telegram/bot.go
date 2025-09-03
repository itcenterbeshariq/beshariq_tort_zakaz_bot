package telegram

import (
	"beshariq_tort_zakaz_bot/services/telegram/hendlers"
	beego "github.com/beego/beego/v2/server/web"
	"gopkg.in/telebot.v4"
	"log"
	"time"
)

func Bot() {

	token := beego.AppConfig.DefaultString("telegram::token", "")

	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(telebot.OnText, hendlers.Text)

	b.Start()
}
