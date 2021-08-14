package postgres

import (
	"context"

	"github.com/LazyBearCT/finance-bot/internal/logger"
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

func (ep *ExpensePostgres) DeleteExpenseByID(ctx context.Context, id int) error {
	return ep.db.Where("id = ?", id).Delete(model.Expense{}).Error
}

func (ep *ExpensePostgres) GetLastExpenses(ctx context.Context) ([]*model.Expense, error) {
	var expenses []*model.Expense
	rows, err := ep.db.Model(new(model.Expense)).Select("expenses.id, expenses.amount, categories.name").Joins(
		"LEFT JOIN categories ON categories.codename = expenses.category_codename",
	).Order("created_at DESC").Limit(10).Rows()
	defer func() {
		_ = rows.Close()
	}()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id, amount int
		var name string
		if err := rows.Scan(&id, &amount, &name); err != nil {
			logger.Error(err)
			return nil, err
		}
		expenses = append(expenses, &model.Expense{
			ID:               id,
			Amount:           amount,
			CategoryCodename: name,
		})
	}
	return expenses, nil
}

func (ep *ExpensePostgres) GetAllTodayExpenses(ctx context.Context) (int, error) {
	var allExpenses int
	if err := ep.db.Model(new(model.Expense)).Select("SUM(amount)").Where(
		"DATE(created_at) = CURRENT_DATE",
	).Find(&allExpenses).Error; err != nil {
		return 0, err
	}
	return allExpenses, nil
}

func (ep *ExpensePostgres) GetBaseTodayExpenses(ctx context.Context) (int, error) {
	var baseExpenses int
	if err := ep.db.Model(new(model.Expense)).Select("SUM(amount)").Where(
		`DATE(created_at) = CURRENT_DATE AND 
			category_codename IN (SELECT codename FROM categories WHERE is_base_expense=true)`,
	).Scan(&baseExpenses).Error; err != nil {
		return 0, err
	}
	return baseExpenses, nil
}
