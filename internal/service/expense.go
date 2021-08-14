package service

import (
	"context"
	"github.com/LazyBearCT/finance-bot/internal/times"

	"github.com/LazyBearCT/finance-bot/internal/logger"
	"github.com/LazyBearCT/finance-bot/internal/model"
	"github.com/oriser/regroup"
	"github.com/pkg/errors"

	"github.com/LazyBearCT/finance-bot/internal/repository"
)

//go:generate mockgen -source=expense.go -destination=mocks/mock_expense.go

type Expense interface {
	GetAllByPeriod(period times.Period) (int, error)
	GetBaseByPeriod(period times.Period) (int, error)
	GetLastExpenses() ([]*model.Expense, error)
	AddExpense(rawMessage string) (*model.Expense, error)
	DeleteByID(id int) error
}

type expenseService struct {
	ctx      context.Context
	repo     repository.Expense
	category Category
}

func NewExpense(ctx context.Context, expenseRepo repository.Expense, category Category) Expense {
	return &expenseService{
		ctx:      ctx,
		repo:     expenseRepo,
		category: category,
	}
}

func (es *expenseService) GetAllByPeriod(period times.Period) (int, error) {
	allExpenses, err := es.repo.GetAllExpensesByPeriod(es.ctx, period)
	if err != nil {
		return 0, err
	}
	return allExpenses, nil
}

func (es *expenseService) GetBaseByPeriod(period times.Period) (int, error) {
	baseExpenses, err := es.repo.GetBaseExpensesByPeriod(es.ctx, period)
	if err != nil {
		return 0, err
	}
	return baseExpenses, nil
}

func (es *expenseService) GetLastExpenses() ([]*model.Expense, error) {
	expenses, err := es.repo.GetLastExpenses(es.ctx)
	if err != nil {
		return nil, err
	}
	return expenses, nil
}

func (es *expenseService) DeleteByID(id int) error {
	return es.repo.DeleteExpenseByID(es.ctx, id)
}

type Message struct {
	Amount       int    `regroup:"amount"`
	CategoryText string `regroup:"text"`
}

var re = regroup.MustCompile("(?P<amount>[\\d ]+) (?P<text>.*)")

func (es *expenseService) AddExpense(rawMessage string) (*model.Expense, error) {
	parsedMessage, err := parseMessage(rawMessage)
	if err != nil {
		return nil, err
	}

	category := es.category.GetByName(parsedMessage.CategoryText)
	if category == nil {
		return nil, errors.New("category not found")
	}

	expense, err := es.repo.CreateExpense(es.ctx, &model.Expense{
		Amount:           parsedMessage.Amount,
		CategoryCodename: category.Codename,
		RawText:          rawMessage,
	})
	if err != nil {
		return nil, err
	}

	return &model.Expense{
		Amount:           expense.Amount,
		CategoryCodename: category.Name,
	}, nil
}

func parseMessage(rawMessage string) (Message, error) {
	var message Message
	if err := re.MatchToTarget(rawMessage, &message); err != nil {
		logger.Error(message)
		return Message{}, errors.New("Не могу понять сообщение. Напишите сообщение в формате, например:\n1500 метро")
	}
	logger.Info(message)
	return message, nil
}
