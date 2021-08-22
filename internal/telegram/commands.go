package telegram

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pkg/errors"
	"gitlab.com/LazyBearCT/finance-bot/internal/logger"
	"gitlab.com/LazyBearCT/finance-bot/pkg/times"
)

var (
	errLast  = errors.New("Расходы ещё не заведены")
	errToday = errors.New("Сегодня ещё нет расходов")
	errMonth = errors.New("В этом месяце ещё нет расходов")
)

const (
	commandStart      = "start"
	commandDelete     = "del"
	commandCategories = "categories"
	commandToday      = "today"
	commandMonth      = "month"
	commandLast       = "last"
	commandLimit      = "limit"
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
	case commandLimit:
		b.handleLimitCommand(message)
	default:
		b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleLimitCommand(message *tgbotapi.Message) {
	id := message.Chat.ID
	limit := b.manager.Budget.GetBaseDailyLimit()
	b.send(id, fmt.Sprintf("Базовый дневной бюджет: %d", limit))
}

func (b *Bot) handleTodayCommand(message *tgbotapi.Message) {
	b.handleCommandByPeriod(message, times.Day)
}

func (b *Bot) handleMonthCommand(message *tgbotapi.Message) {
	b.handleCommandByPeriod(message, times.Month)
}

func (b *Bot) handleCommandByPeriod(message *tgbotapi.Message, period times.Period) {
	id := message.Chat.ID
	text, err := b.getStatisticsByPeriod(period)
	if err != nil {
		b.handleError(id, err)
	}
	b.send(id, text)
}

func (b *Bot) handleLastCommand(message *tgbotapi.Message) {
	id := message.Chat.ID

	expenses, err := b.manager.Expense.GetLastExpenses()
	if err != nil {
		b.handleError(id, errLast)
		return
	}
	lastExpenses := make([]string, 0, len(expenses))
	for _, expense := range expenses {
		info := fmt.Sprintf("%d руб. на %s — нажми ", expense.Amount, expense.CategoryCodename)
		del := fmt.Sprintf("/%s%d для удаления", commandDelete, expense.ID)
		lastExpenses = append(lastExpenses, info+del)
	}
	b.send(id, "Последние сохранённые траты:\n\n* "+strings.Join(lastExpenses, "\n\n* "))
}

func (b *Bot) getStatisticsByPeriod(period times.Period) (string, error) {
	allExpenses := b.manager.Expense.GetAllByPeriod(period)
	var periodError error
	switch period {
	case times.Day:
		periodError = errToday
	case times.Month:
		periodError = errMonth
	default:
		panic("unknown period")
	}
	if allExpenses == 0 {
		return "", periodError
	}
	baseExpenses := b.manager.Expense.GetBaseByPeriod(period)
	dailyLimit := b.manager.Budget.GetBaseDailyLimit()
	var text string
	all := fmt.Sprintf("всего — %d руб.\n", allExpenses)
	switch period {
	case times.Day:
		text = "Расходы сегодня:\n"
		text += all + fmt.Sprintf("базовые — %d руб. из %d руб.\n\n", baseExpenses, dailyLimit)
		text += fmt.Sprintf("За текущий месяц: /%s", commandMonth)
	case times.Month:
		text = "Расходы в текущем месяце:\n"
		text += all + fmt.Sprintf("базовые — %d руб. из %d руб.", baseExpenses, time.Now().Day()*dailyLimit)
	}
	return text, nil
}

func (b *Bot) handleCategoriesCommand(message *tgbotapi.Message) {
	categories := b.manager.Category.GetAll()
	categoriesInfo := make([]string, 0, len(categories))
	for _, c := range categories {
		categoriesInfo = append(categoriesInfo, c.Name+" ("+strings.Join(c.Aliases, ", ")+")")
	}
	b.send(message.Chat.ID, "Категории трат:\n\n* "+strings.Join(categoriesInfo, "\n* "))
}

func (b *Bot) handleDeleteCommand(message *tgbotapi.Message) {
	id := message.Chat.ID

	rowID, err := strconv.Atoi(message.Text[len(commandDelete)+1:])
	if err != nil {
		b.handleError(id, err)
		return
	}

	if err := b.manager.Expense.DeleteByID(rowID); err != nil {
		b.handleError(id, err)
		return
	}

	b.send(id, "Удалил")
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) {
	start := "Бот для учёта финансов\n\n"
	start += fmt.Sprintf("Добавить расход: 250 такси\nСегодняшняя статистика: /%s\n", commandToday)
	start += fmt.Sprintf("За текущий месяц: /%s\nПоследние внесённые расходы: /%s\n", commandMonth, commandLast)
	start += fmt.Sprintf("Категории трат: /%s\nБазовый дневной бюджет: /%s", commandCategories, commandLimit)
	b.send(message.Chat.ID, start)
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) {
	b.send(message.Chat.ID, "Я не знаю такой команды :(")
}
