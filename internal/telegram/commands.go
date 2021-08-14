package telegram

import (
	"fmt"
	"github.com/LazyBearCT/finance-bot/internal/times"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"

	"github.com/LazyBearCT/finance-bot/internal/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	todayError = errors.New("Сегодня ещё нет расходов")
	lastError = errors.New("Расходы ещё не заведены")
	monthError = errors.New("В этом месяце ещё нет расходов")
)

const (
	commandStart      = "start"
	commandDelete     = "del"
	commandCategories = "categories"
	commandToday      = "today"
	commandMonth = "month"
	commandLast = "last"
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
	case commandMonth:
		b.handleMonthCommand(message)
	case commandLast:
		b.handleLastCommand(message)
	case "limit":
		limit, err := b.manager.Budget.GetDailyLimitByName("base")
		if err != nil {
			b.handleError(message.Chat.ID, err)
			return
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Базовый дневной бюджет: %d", limit))
		b.send(msg)
	default:
		b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleMonthCommand(message *tgbotapi.Message) {
	id := message.Chat.ID
	msg := tgbotapi.NewMessage(id, b.getStatisticsByPeriod(id, times.Month))
	b.send(msg)
}

func (b *Bot) handleLastCommand(message *tgbotapi.Message) {
	id := message.Chat.ID

	expenses, err := b.manager.Expense.GetLastExpenses()
	if err != nil {
		b.handleError(id, lastError)
		return
	}
	lastExpenses := make([]string, 0, len(expenses))
	for _, expense := range expenses {
		info := fmt.Sprintf("%d руб. на %s — нажми ", expense.Amount, expense.CategoryCodename)
		del := fmt.Sprintf("/del%d для удаления", expense.ID)
		lastExpenses = append(lastExpenses, info + del)
	}
	msg := tgbotapi.NewMessage(id, "Последние сохранённые траты:\n\n* " + strings.Join(lastExpenses, "\n\n* "))
	b.send(msg)
}

func (b *Bot) handleTodayCommand(message *tgbotapi.Message) {
	id := message.Chat.ID
	msg := tgbotapi.NewMessage(id, b.getStatisticsByPeriod(id, times.Day))
	b.send(msg)
}

func (b *Bot) getStatisticsByPeriod(id int64, period times.Period) string {
	allExpenses, err := b.manager.Expense.GetAllByPeriod(period)
	if err != nil {
		b.handleError(id, todayError)
		return ""
	}
	baseExpenses, err := b.manager.Expense.GetBaseByPeriod(period)
	if err != nil {
		b.handleError(id, todayError)
		return ""
	}
	dailyLimit, err := b.manager.Budget.GetDailyLimitByName("base")
	if err != nil {
		b.handleError(id, todayError)
		return ""
	}
	var text string
	all := fmt.Sprintf("всего — %d руб.\n", allExpenses)
	text += fmt.Sprintf("базовые — %d руб. из %d руб.\n\n", baseExpenses, dailyLimit)
	text += "За текущий месяц: /month"
	switch period {
	case times.Day:
		text = "Расходы сегодня:\n"
		text += all + fmt.Sprintf("базовые — %d руб. из %d руб.\n\n", baseExpenses, dailyLimit)
		text += "За текущий месяц: /month"
	case times.Month:
		text = "Расходы в текущем месяце:\n"
		text += all + fmt.Sprintf("базовые — %d руб. из %d руб.", baseExpenses, time.Now().Day() * dailyLimit)
	}
	return text
}

func (b *Bot) handleCategoriesCommand(message *tgbotapi.Message) {
	categories := b.manager.Category.GetAll()
	categoriesInfo := make([]string, 0, len(categories))
	for _, c := range categories {
		categoriesInfo = append(categoriesInfo, c.Name+" ("+strings.Join(c.Aliases, ", ")+")")
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, b.config.Message.Response.Categories+strings.Join(categoriesInfo, "\n* "))
	b.send(msg)
}

func (b *Bot) handleDeleteCommand(message *tgbotapi.Message) {
	id := message.Chat.ID

	rowID, err := strconv.Atoi(message.Text[len(commandDelete) + 1:])
	if err != nil {
		b.handleError(id, err)
		return
	}

	if err := b.manager.Expense.DeleteByID(rowID); err != nil {
		b.handleError(id, err)
		return
	}

	msg := tgbotapi.NewMessage(id, b.config.Message.Response.Delete)
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
