package hendlers

import (
	"beshariq_tort_zakaz_bot/services/telegram/helpers"
	"gopkg.in/telebot.v4"
)

func Start(c telebot.Context) error {

	menu := &telebot.ReplyMarkup{ResizeKeyboard: true}
	btnM := menu.Text(helpers.BtnMenu)
	btnO := menu.Text(helpers.BtnOrder)
	btnC := menu.Text(helpers.BtnContact)

	menu.Reply(
		menu.Row(btnM, btnO),
		menu.Row(btnC),
	)

	msg := `
Assalomu alaykum! 👋

Siz Beshariq To‘rt Zakaz Botiga xush kelibsiz.
📌 Bu bot orqali siz:
	1️⃣ To‘rt buyurtma qilishingiz mumkin
	2️⃣ Narxlar va menyu bilan tanishishingiz mumkin
	3️⃣ Buyurtmangizni tez va oson rasmiylashtirishingiz mumkin

👉 Quyidagi tugmalardan foydalanib davom eting:
	🍰 Menyu
	🛒 Buyurtma berish
	📞 Aloqa
`
	return c.Send(msg, menu)
}
