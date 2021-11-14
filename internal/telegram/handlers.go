package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/maypok86/finance-bot/internal/logger"
	"github.com/maypok86/finance-bot/pkg/times"
	"github.com/pkg/errors"
)

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		message := update.Message
		if message == nil { // ignore any non-Message Updates
			continue
		}
		if message.From.ID != b.config.AccessID {
			return errors.New("wrong id")
		}

		if message.IsCommand() {
			b.handleCommand(message)
			continue
		}
		b.handleMessage(message)
	}
	return nil
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	logger.Infof("[%s] %s", message.From.UserName, message.Text)
	id := message.Chat.ID

	expense, err := b.manager.Expense.AddExpense(message.Text)
	if err != nil {
		b.handleError(id, err)
		return
	}

	amounts := fmt.Sprintf("Добавлены траты %d руб на %s.\n\n", expense.Amount, expense.CategoryCodename)
	statistics, err := b.getStatisticsByPeriod(times.Day)
	text := amounts
	if err == nil {
		text += statistics
	}
	b.send(id, text)
}

func (b *Bot) handleError(id int64, err error) {
	logger.Error(err)
	b.send(id, err.Error())
}
