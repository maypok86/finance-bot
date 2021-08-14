package service

import (
	"context"

	"github.com/LazyBearCT/finance-bot/internal/repository"
)

//go:generate mockgen -source=expense.go -destination=mocks/mock_expense.go

type Expense interface {
	GetAllToday() (int, error)
	GetBaseToday() (int, error)
}

type expenseService struct {
	ctx  context.Context
	repo repository.Expense
}

func NewExpense(ctx context.Context, repo repository.Expense) Expense {
	return &expenseService{
		ctx:  ctx,
		repo: repo,
	}
}

func (es *expenseService) GetAllToday() (int, error) {
	allExpenses, err := es.repo.GetAllTodayExpenses(es.ctx)
	if err != nil {
		return 0, err
	}
	return allExpenses, err
}

func (es *expenseService) GetBaseToday() (int, error) {
	baseExpenses, err := es.repo.GetBaseTodayExpenses(es.ctx)
	if err != nil {
		return 0, err
	}
	return baseExpenses, err
}
