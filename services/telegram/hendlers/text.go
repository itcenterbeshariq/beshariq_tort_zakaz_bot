package hendlers

import (
	"beshariq_tort_zakaz_bot/services/telegram/helpers"
	"gopkg.in/telebot.v4"
)

func Text(c telebot.Context) error {
	text := c.Text()

	switch text {
	case helpers.BtnContact:
		return Contact(c)
	case helpers.BtnMenu:
		return Menu(c)
	case helpers.BtnOrder:
		return Order(c)
	case helpers.CommandStart:
		return Start(c)
	}

	return nil
}
