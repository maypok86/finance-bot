package repository

import (
	"context"

	"gitlab.com/LazyBearCT/finance-bot/internal/model"
)

//go:generate mockgen -source=budget.go -destination=mocks/mock_budget.go

// Budget repository.
type Budget interface {
	GetBudgetByCodename(ctx context.Context, name string) (*model.Budget, error)
	GetBaseBudget(ctx context.Context) (*model.Budget, error)
}
