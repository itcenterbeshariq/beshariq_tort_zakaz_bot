package telegram

import (
	beego "github.com/beego/beego/v2/server/web"
	"gopkg.in/telebot.v4"
	"log"
	"time"
)

func Bot() {

	// sss

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

	b.Handle("/start", func(c telebot.Context) error {
		return c.Send("Assalomu alaykum! 👋\nSiz Beshariq To‘rt Zakaz Botiga xush kelibsiz.\n\n📌 Bu bot orqali siz:\n1️⃣ To‘rt buyurtma qilishingiz mumkin\n2️⃣ Narxlar va menyu bilan tanishishingiz mumkin\n3️⃣ Buyurtmangizni tez va oson rasmiylashtirishingiz mumkin\n\n" +
			"👉 Quyidagi tugmalardan foydalanib davom eting:\n\n🍰 Menyu\n\n🛒 Buyurtma berish\n\n📞 Aloqa")
	})

	b.Start()
}
