package postgres

import (
	"context"
	"github.com/LazyBearCT/finance-bot/internal/model"
	"gorm.io/gorm"
)

type ExpensePostgres struct {
	db *gorm.DB
}

func NewExpenseRepository(db *DB) *ExpensePostgres {
	return &ExpensePostgres{
		db: db.db,
	}
}

func (ep *ExpensePostgres) CreateExpense(ctx context.Context, expense *model.Expense) (*model.Expense, error) {
	if err := ep.db.Create(expense).Error; err != nil {
		return nil, err
	}
	return expense, nil
}

func (ep *ExpensePostgres) GetAllTodayExpenses(ctx context.Context) (int, error) {
	var allExpenses int
	if err := ep.db.Model(new(model.Expense)).Select("SUM(amount)").Where(
		"created_date = CURRENT_DATE",
	).Find(&allExpenses).Error; err != nil {
		return 0, err
	}
	return allExpenses, nil
}

func (ep *ExpensePostgres) GetBaseTodayExpenses(ctx context.Context) (int, error) {
	var baseExpenses int
	if err := ep.db.Model(new(model.Expense)).Select("SUM(amount)").Where(
		`created_date = CURRENT_DATE AND 
			category_codename IN (SELECT codename FROM categories WHERE is_base_expense=true)`,
	).Scan(&baseExpenses).Error; err != nil {
		return 0, err
	}
	return baseExpenses, nil
}
