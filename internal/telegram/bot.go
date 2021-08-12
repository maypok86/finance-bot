package telegram

import (
	"github.com/LazyBearCT/finance-bot/internal/config"
	"github.com/LazyBearCT/finance-bot/internal/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pkg/errors"
)

// Bot telegram
type Bot struct {
	bot    *tgbotapi.BotAPI
	userID int
}

// New telegram bot
func New(c *config.Bot) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(c.BotToken)
	if err != nil {
		return nil, err
	}

	logger.Infof("Authorized on account %s", bot.Self.UserName)

	return &Bot{
		bot:    bot,
		userID: c.AccessID,
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

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		if update.Message.From.ID != b.userID {
			return errors.New("wrong id")
		}

		logger.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		_, err := b.bot.Send(msg)
		if err != nil {
			logger.Error(err)
		}
	}
	return nil
}
