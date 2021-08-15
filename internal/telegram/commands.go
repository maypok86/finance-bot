package telegram

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/LazyBearCT/finance-bot/internal/times"
	"github.com/pkg/errors"

	"github.com/LazyBearCT/finance-bot/internal/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	todayError = errors.New("Сегодня ещё нет расходов")
	lastError  = errors.New("Расходы ещё не заведены")
	monthError = errors.New("В этом месяце ещё нет расходов")
)

const (
	commandStart      = "start"
	commandDelete     = "del"
	commandCategories = "categories"
	commandToday      = "today"
	commandMonth      = "month"
	commandLast       = "last"
	commandLimit = "limit"
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
	limit, err := b.manager.Budget.GetDailyLimitByName("base")
	if err != nil {
		b.handleError(id, err)
		return
	}
	msg := tgbotapi.NewMessage(id, fmt.Sprintf("Базовый дневной бюджет: %d", limit))
	b.send(msg)
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
		lastExpenses = append(lastExpenses, info+del)
	}
	msg := tgbotapi.NewMessage(id, "Последние сохранённые траты:\n\n* "+strings.Join(lastExpenses, "\n\n* "))
	b.send(msg)
}

func (b *Bot) handleTodayCommand(message *tgbotapi.Message) {
	id := message.Chat.ID
	msg := tgbotapi.NewMessage(id, b.getStatisticsByPeriod(id, times.Day))
	b.send(msg)
}

func (b *Bot) getStatisticsByPeriod(id int64, period times.Period) string {
	allExpenses, err := b.manager.Expense.GetAllByPeriod(period)
	var periodError error
	switch period {
	case times.Day:
		periodError = todayError
	case times.Month:
		periodError = monthError
	default:
		panic("unknown period")
	}
	if err != nil {
		b.handleError(id, periodError)
		return ""
	}
	baseExpenses, err := b.manager.Expense.GetBaseByPeriod(period)
	if err != nil {
		b.handleError(id, periodError)
		return ""
	}
	dailyLimit, err := b.manager.Budget.GetDailyLimitByName("base")
	if err != nil {
		b.handleError(id, periodError)
		return ""
	}
	var text string
	all := fmt.Sprintf("всего — %d руб.\n", allExpenses)
	switch period {
	case times.Day:
		text = "Расходы сегодня:\n"
		text += all + fmt.Sprintf("базовые — %d руб. из %d руб.\n\n", baseExpenses, dailyLimit)
		text += "За текущий месяц: /month"
	case times.Month:
		text = "Расходы в текущем месяце:\n"
		text += all + fmt.Sprintf("базовые — %d руб. из %d руб.", baseExpenses, time.Now().Day()*dailyLimit)
	}
	return text
}

func (b *Bot) handleCategoriesCommand(message *tgbotapi.Message) {
	categories := b.manager.Category.GetAll()
	categoriesInfo := make([]string, 0, len(categories))
	for _, c := range categories {
		categoriesInfo = append(categoriesInfo, c.Name+" ("+strings.Join(c.Aliases, ", ")+")")
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, "Категории трат:\n\n* "+strings.Join(categoriesInfo, "\n* "))
	b.send(msg)
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

	msg := tgbotapi.NewMessage(id, "Удалил")
	b.send(msg)
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) {
	start := "Бот для учёта финансов\n\n"
	start += "Добавить расход: 250 такси\nСегодняшняя статистика: /today\n"
	start += "За текущий месяц: /month\nПоследние внесённые расходы: /last\nКатегории трат: /categories\n"
	start += "Базовый дневной бюджет: /limit"
	msg := tgbotapi.NewMessage(message.Chat.ID, start)
	b.send(msg)
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды :(")
	b.send(msg)
}
