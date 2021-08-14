package telegram

import (
	"strings"

	"github.com/LazyBearCT/finance-bot/internal/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandStart      = "start"
	commandDelete     = "del"
	commandCategories = "categories"
	commandHelp       = "help"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) {
	command := message.Command()
	logger.Infof("[%s] %s", message.From.UserName, command)

	if strings.HasPrefix(command, commandDelete) {
		b.handleDeleteCommand(message)
		return
	}

	switch command {
	case commandStart, commandHelp:
		b.handleStartCommand(message)
	case commandCategories:
		b.handleCategoriesCommand(message)
	default:
		b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleCategoriesCommand(message *tgbotapi.Message) {
	categories, err := b.manager.Category.GetAll()
	if err != nil {
		logger.Error(err)
		msg := tgbotapi.NewMessage(message.Chat.ID, err.Error())
		b.send(msg)
		return
	}
	categoriesInfo := make([]string, 0, len(categories))
	for _, c := range categories {
		categoriesInfo = append(categoriesInfo, c.Name + " (" + strings.Join(c.Aliases, ", ") + ")")
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, b.config.Message.Response.Categories + strings.Join(categoriesInfo, "\n* "))
	b.send(msg)
}

func (b *Bot) handleDeleteCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.config.Message.Response.Delete)
	b.send(msg)
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.config.Message.Response.Start)
	b.send(msg)
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.config.Message.Response.Unknown)
	b.send(msg)
}
