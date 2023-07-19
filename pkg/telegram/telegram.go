package telegram

import (
	"time"

	"github.com/shalfbea/openSupportTgBot/pkg/config"
	"github.com/shalfbea/openSupportTgBot/pkg/logger"
	tele "gopkg.in/telebot.v3"
)

type Bot struct {
	bot    *tele.Bot
	config *config.Config
	logger logger.Logger
}

func NewBot(config *config.Config, logger logger.Logger) (bot *Bot, err error) {
	prefs := tele.Settings{
		Token:  config.TelegramToken,
		Poller: &tele.LongPoller{Timeout: time.Duration(config.PollingTimeout) * time.Second},
	}
	telebot, err := tele.NewBot(prefs)
	if err != nil {
		logger.Errorf("token : %s, Problem with newBot: %v", config.TelegramToken, err)
		return nil, err
	}
	bot = &Bot{
		bot:    telebot,
		config: config,
		logger: logger,
	}

	bot.registerHandlers()
	bot.registerCommands()
	return bot, err
}

func RunBot(b *Bot) {
	b.logger.Info("Bot started!")
	b.bot.Start()
}
