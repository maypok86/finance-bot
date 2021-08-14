package postgres

import (
	"context"

	"github.com/LazyBearCT/finance-bot/internal/model"
	"gorm.io/gorm"
)

type BudgetPostgres struct {
	db *gorm.DB
}

func NewBudgetRepository(db *DB) *BudgetPostgres {
	return &BudgetPostgres{
		db: db.db,
	}
}

func (bp *BudgetPostgres) GetBudgetByCodename(ctx context.Context, name string) (*model.Budget, error) {
	var budget model.Budget
	if err := bp.db.Where("codename = ?", name).Take(&budget).Error; err != nil {
		return nil, err
	}
	return &budget, nil
}
