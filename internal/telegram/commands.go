package telegram

import (
	"fmt"
	"strings"

	"github.com/LazyBearCT/finance-bot/internal/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandStart      = "start"
	commandDelete     = "del"
	commandCategories = "categories"
	commandToday      = "today"
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
	case commandToday:
		b.handleTodayCommand(message)
	case "limit":
		limit, err := b.manager.Budget.GetDailyLimitByName("base")
		if err != nil {
			b.handleError(message.Chat.ID, err)
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Базовый дневной бюджет: %d", limit))
		b.send(msg)
	default:
		b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleTodayCommand(message *tgbotapi.Message) {
	id := message.Chat.ID
	allExpenses, err := b.manager.Expense.GetAllToday()
	if err != nil {
		b.handleError(id, err)
		return
	}
	baseExpenses, err := b.manager.Expense.GetBaseToday()
	if err != nil {
		b.handleError(id, err)
		return
	}
	dailyLimit, err := b.manager.Budget.GetDailyLimitByName("base")
	if err != nil {
		b.handleError(id, err)
		return
	}
	msg := tgbotapi.NewMessage(id, "")
	msg.Text += "Расходы сегодня:\n"
	msg.Text += fmt.Sprintf("всего — %d руб.\n", allExpenses)
	msg.Text += fmt.Sprintf("базовые — %d руб. из %d руб.\n\n", baseExpenses, dailyLimit)
	msg.Text += "За текущий месяц: /month"
	b.send(msg)
}

func (b *Bot) handleCategoriesCommand(message *tgbotapi.Message) {
	categories, err := b.manager.Category.GetAll()
	if err != nil {
		b.handleError(message.Chat.ID, err)
		return
	}
	categoriesInfo := make([]string, 0, len(categories))
	for _, c := range categories {
		categoriesInfo = append(categoriesInfo, c.Name+" ("+strings.Join(c.Aliases, ", ")+")")
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, b.config.Message.Response.Categories+strings.Join(categoriesInfo, "\n* "))
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
