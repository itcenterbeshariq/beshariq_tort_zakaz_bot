package hendlers

import "gopkg.in/telebot.v4"

func Menu(c telebot.Context) error {
	menu := &telebot.ReplyMarkup{}

	btn100 := menu.Data("100", "btn100")
	btn120 := menu.Data("120", "btn120")
	btn150 := menu.Data("150", "btn150")
	btn200 := menu.Data("200", "btn200")
	btn250 := menu.Data("250", "btn250")
	btn300 := menu.Data("300", "btn300")

	menu.Inline(
		menu.Row(btn100),
		menu.Row(btn120),
		menu.Row(btn150),
		menu.Row(btn200),
		menu.Row(btn250),
		menu.Row(btn300),
	)
	text := c.Text()

	switch text {
	case "🍰 Menyu":
		return c.Send("qancha puligiz borligiga qarab narhlarni ustiga bosing:", menu)
	default:
		return c.Send("🍰 Menyu  dan tanlang.")
	}
}
