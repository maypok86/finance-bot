package telegram

import (
	"github.com/LazyBearCT/finance-bot/internal/config"
	"github.com/LazyBearCT/finance-bot/internal/logger"
	"github.com/LazyBearCT/finance-bot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Bot telegram
type Bot struct {
	bot     *tgbotapi.BotAPI
	config  *config.Bot
	manager *service.Manager
}

// New telegram bot
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

// Start telegram bot
func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	if err := b.handleUpdates(updates); err != nil {
		return err
	}

	return nil
}

func (b *Bot) send(msg tgbotapi.MessageConfig) {
	if msg.Text == "" {
		return
	}
	_, err := b.bot.Send(msg)
	if err != nil {
		logger.Error(err)
	}
}
