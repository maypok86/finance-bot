package service

import (
	"context"

	"gitlab.com/LazyBearCT/finance-bot/internal/repository"
)

//go:generate mockgen -source=budget.go -destination=mocks/mock_budget.go

// Budget service.
type Budget interface {
	GetDailyLimitByName(name string) int
	GetBaseDailyLimit() int
}

type budgetService struct {
	ctx  context.Context
	repo repository.Budget
}

// NewBudget creates a new Budget instance.
func NewBudget(ctx context.Context, repo repository.Budget) Budget {
	return &budgetService{
		ctx:  ctx,
		repo: repo,
	}
}

// GetDailyLimitByName returns a daily limit by name.
func (bs *budgetService) GetDailyLimitByName(name string) int {
	budget, err := bs.repo.GetBudgetByCodename(bs.ctx, name)
	if err != nil {
		return 0
	}
	return budget.DailyLimit
}

// GetBaseDailyLimit returns a base daily limit.
func (bs *budgetService) GetBaseDailyLimit() int {
	budget, err := bs.repo.GetBaseBudget(bs.ctx)
	if err != nil {
		return 0
	}
	return budget.DailyLimit
}
