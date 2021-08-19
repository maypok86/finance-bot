package service

import (
	"context"
	"testing"

	"github.com/LazyBearCT/finance-bot/internal/model"
	mock_repository "github.com/LazyBearCT/finance-bot/internal/repository/mocks"
	"github.com/LazyBearCT/finance-bot/pkg/random"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestGetBaseDailyLimit(t *testing.T) {
	input := &model.Budget{
		Codename:   "base",
		DailyLimit: 500,
	}
	tests := []struct {
		name         string
		expectations func(budgetRepo *mock_repository.MockBudget)
		input        *model.Budget
	}{
		{
			name: "valid and found",
			expectations: func(budgetRepo *mock_repository.MockBudget) {
				budgetRepo.EXPECT().GetBaseBudget(context.Background()).Return(input, nil)
			},
			input: input,
		},
		{
			name: "not found",
			expectations: func(budgetRepo *mock_repository.MockBudget) {
				budgetRepo.EXPECT().GetBaseBudget(context.Background()).Return(nil, errors.New("error"))
			},
			input: new(model.Budget),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_repository.NewMockBudget(c)
			test.expectations(repo)
			service := NewBudget(context.Background(), repo)
			limit := service.GetBaseDailyLimit()
			require.Equal(t, test.input.DailyLimit, limit)
		})
	}
}

func TestGetDailyLimitByName(t *testing.T) {
	input := &model.Budget{
		Codename:   random.Name(),
		DailyLimit: random.Int(),
	}
	tests := []struct {
		name         string
		expectations func(budgetRepo *mock_repository.MockBudget, budget *model.Budget)
		input        *model.Budget
	}{
		{
			name: "valid and found",
			expectations: func(budgetRepo *mock_repository.MockBudget, budget *model.Budget) {
				budgetRepo.EXPECT().GetBudgetByCodename(context.Background(), budget.Codename).Return(budget, nil)
			},
			input: input,
		},
		{
			name: "not found",
			expectations: func(budgetRepo *mock_repository.MockBudget, budget *model.Budget) {
				budgetRepo.EXPECT().GetBudgetByCodename(context.Background(), budget.Codename).Return(nil, errors.New("error"))
			},
			input: new(model.Budget),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_repository.NewMockBudget(c)
			test.expectations(repo, test.input)
			service := NewBudget(context.Background(), repo)
			limit := service.GetDailyLimitByName(test.input.Codename)
			require.Equal(t, test.input.DailyLimit, limit)
		})
	}
}
