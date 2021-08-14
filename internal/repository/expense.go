package repository

import (
	"context"

	"github.com/LazyBearCT/finance-bot/internal/model"
)

//go:generate mockgen -source=expense.go -destination=mocks/mock_expense.go

type Expense interface {
	GetAllTodayExpenses(ctx context.Context) (int, error)
	GetBaseTodayExpenses(ctx context.Context) (int, error)
	GetLastExpenses(ctx context.Context) ([]*model.Expense, error)
	CreateExpense(ctx context.Context, expense *model.Expense) (*model.Expense, error)
	DeleteExpenseByID(ctx context.Context, id int) error
}
