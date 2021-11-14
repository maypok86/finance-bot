package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/maypok86/finance-bot/internal/config"
	"github.com/maypok86/finance-bot/internal/logger"
	"github.com/maypok86/finance-bot/internal/service"
)

// Bot telegram.
type Bot struct {
	bot     *tgbotapi.BotAPI
	config  *config.Bot
	manager *service.Manager
}

// New create a new Bot instance.
func New(c *config.Bot, manager *service.Manager) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(c.BotToken)
	if err != nil {
		return nil, err
	}

	logger.Infof("Authorized on account %s", bot.Self.UserName)

	return &Bot{
		bot:     bot,
		config:  c,
		manager: manager,
	}, nil
}

// Start telegram bot.
func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	return b.handleUpdates(updates)
}

func (b *Bot) send(id int64, text string) {
	if text == "" {
		return
	}
	_, err := b.bot.Send(tgbotapi.NewMessage(id, text))
	if err != nil {
		logger.Error(err)
	}
}
