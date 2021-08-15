package repository

import (
	"context"

	"github.com/LazyBearCT/finance-bot/internal/model"
)

//go:generate mockgen -source=budget.go -destination=mocks/mock_budget.go

type Budget interface {
	GetBudgetByCodename(ctx context.Context, name string) (*model.Budget, error)
}
