package telegram

import (
	"fmt"

	"github.com/LazyBearCT/finance-bot/internal/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pkg/errors"
)

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		if update.Message.From.ID != b.config.AccessID {
			return errors.New("wrong id")
		}

		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}
		b.handleMessage(update.Message)
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
	msg := tgbotapi.NewMessage(id, amounts+b.getTodayStatistics(id))
	b.send(msg)
}

func (b *Bot) handleError(id int64, err error) {
	logger.Error(err)
	msg := tgbotapi.NewMessage(id, err.Error())
	b.send(msg)
}
