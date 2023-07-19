package telegram

import (
	tele "gopkg.in/telebot.v3"
)

func (b *Bot) registerHandlers() {
	b.bot.Handle(tele.OnText, func(c tele.Context) error {
		return nil
	})
}
