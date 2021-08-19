package service

import (
	"context"

	"github.com/LazyBearCT/finance-bot/internal/logger"
	"github.com/LazyBearCT/finance-bot/internal/model"
	"github.com/LazyBearCT/finance-bot/internal/repository"
	"github.com/LazyBearCT/finance-bot/pkg/times"
	"github.com/oriser/regroup"
	"github.com/pkg/errors"
)

//go:generate mockgen -source=expense.go -destination=mocks/mock_expense.go

// Expense service
type Expense interface {
	GetAllByPeriod(period times.Period) int
	GetBaseByPeriod(period times.Period) int
	GetLastExpenses() ([]*model.Expense, error)
	AddExpense(rawMessage string) (*model.Expense, error)
	DeleteByID(id int) error
}

type expenseService struct {
	ctx      context.Context
	repo     repository.Expense
	category Category
}

// NewExpense creates a new instance of Expense
func NewExpense(ctx context.Context, expenseRepo repository.Expense, category Category) Expense {
	return &expenseService{
		ctx:      ctx,
		repo:     expenseRepo,
		category: category,
	}
}

// GetAllByPeriod returns all model.Expense instances by period
func (es *expenseService) GetAllByPeriod(period times.Period) int {
	allExpenses, err := es.repo.GetAllExpensesByPeriod(es.ctx, period)
	if err != nil {
		return 0
	}
	return allExpenses
}

// GetBaseByPeriod returns base model.Expense instances by period
func (es *expenseService) GetBaseByPeriod(period times.Period) int {
	baseExpenses, err := es.repo.GetBaseExpensesByPeriod(es.ctx, period)
	if err != nil {
		return 0
	}
	return baseExpenses
}

// GetLastExpenses returns last model.Expense instances
func (es *expenseService) GetLastExpenses() ([]*model.Expense, error) {
	return es.repo.GetLastExpenses(es.ctx)
}

// DeleteByID deletes model.Expense instance by id
func (es *expenseService) DeleteByID(id int) error {
	return es.repo.DeleteExpenseByID(es.ctx, id)
}

// Message ...
type Message struct {
	Amount       int    `regroup:"amount"`
	CategoryText string `regroup:"text"`
}

var re = regroup.MustCompile("(?P<amount>[\\d ]+) (?P<text>.*)")

// AddExpense creates a new model.Expense instance
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
