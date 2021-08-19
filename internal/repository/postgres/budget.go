package postgres

import (
	"context"

	"github.com/LazyBearCT/finance-bot/internal/model"
	"gorm.io/gorm"
)

// BudgetPostgres repository.
type BudgetPostgres struct {
	db *gorm.DB
}

// NewBudgetRepository creates a new BudgetPostgres instance.
func NewBudgetRepository(db *DB) *BudgetPostgres {
	return &BudgetPostgres{
		db: db.db,
	}
}

// GetBudgetByCodename returns model.Budget by codename.
func (bp *BudgetPostgres) GetBudgetByCodename(ctx context.Context, name string) (*model.Budget, error) {
	var budget model.Budget
	if err := bp.db.Where("codename = ?", name).Take(&budget).Error; err != nil {
		return nil, err
	}
	return &budget, nil
}

// GetBaseBudget returns base model.Budget.
func (bp *BudgetPostgres) GetBaseBudget(ctx context.Context) (*model.Budget, error) {
	return bp.GetBudgetByCodename(ctx, "base")
}
