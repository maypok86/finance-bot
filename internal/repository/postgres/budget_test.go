package postgres

import (
	"context"
	"github.com/LazyBearCT/finance-bot/internal/model"
	"github.com/LazyBearCT/finance-bot/pkg/random"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func createRandomBudget() *model.Budget {
	return &model.Budget{
		Codename: random.Name(),
		DailyLimit: int(random.Int(0, math.MaxInt32)),
	}
}

func TestGetBaseBudget(t *testing.T) {
	budget := &model.Budget{
		Codename: "base",
		DailyLimit: 500,
	}
	repo := NewBudgetRepository(db)
	b, err := repo.GetBaseBudget(context.Background())
	require.NoError(t, err)
	require.Equal(t, budget.Codename, b.Codename)
	require.Equal(t, budget.DailyLimit, b.DailyLimit)
}

func TestGetBudgetByCodename(t *testing.T) {
	budget := createRandomBudget()
	err := db.db.Create(budget).Error
	require.NoError(t, err)
	repo := NewBudgetRepository(db)
	b, err := repo.GetBudgetByCodename(context.Background(), budget.Codename)
	require.NoError(t, err)
	require.Equal(t, budget.Codename, b.Codename)
	require.Equal(t, budget.DailyLimit, b.DailyLimit)
}
