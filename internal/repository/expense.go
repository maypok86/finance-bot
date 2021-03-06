package repository

import (
	"context"

	"github.com/maypok86/finance-bot/internal/model"
	"github.com/maypok86/finance-bot/pkg/times"
)

//go:generate mockgen -source=expense.go -destination=mocks/mock_expense.go

// Expense repository.
type Expense interface {
	GetAllExpensesByPeriod(ctx context.Context, period times.Period) (int, error)
	GetBaseExpensesByPeriod(ctx context.Context, period times.Period) (int, error)
	GetLastExpenses(ctx context.Context) ([]*model.Expense, error)
	CreateExpense(ctx context.Context, expense *model.Expense) (*model.Expense, error)
	DeleteExpenseByID(ctx context.Context, id int) error
}
