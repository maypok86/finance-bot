package service

import (
	"context"
	"github.com/LazyBearCT/finance-bot/internal/repository"
)

//go:generate mockgen -source=budget.go -destination=mocks/mock_budget.go

type Budget interface {
	GetDailyLimitByName(name string) (int, error)
}

type budgetService struct {
	ctx        context.Context
	repo       repository.Budget
}

func NewBudget(ctx context.Context, repo repository.Budget) Budget {
	return &budgetService{
		ctx:  ctx,
		repo: repo,
	}
}

func (bs *budgetService) GetDailyLimitByName(name string) (int, error) {
	budget, err := bs.repo.GetBudgetByCodename(bs.ctx, name)
	if err != nil {
		return 0, nil
	}
	return budget.DailyLimit, nil
}