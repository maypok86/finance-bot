package repository

import (
	"context"
)

//go:generate mockgen -source=expense.go -destination=mocks/mock_expense.go

type Expense interface {
	GetAllTodayExpenses(ctx context.Context) (int, error)
	GetBaseTodayExpenses(ctx context.Context) (int, error)
}
