package telegram

import (
	tele "gopkg.in/telebot.v3"
)

// registerCommands register handlers for bot commands
func (b *Bot) registerCommands() {
	b.bot.Handle("/start", func(c tele.Context) error {
		return c.Send(b.config.Messages.Start)
	})

}
