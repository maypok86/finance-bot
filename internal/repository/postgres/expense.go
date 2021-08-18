package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/LazyBearCT/finance-bot/pkg/times"

	"github.com/LazyBearCT/finance-bot/internal/logger"
	"github.com/LazyBearCT/finance-bot/internal/model"
	"gorm.io/gorm"
)

// ExpensePostgres repository
type ExpensePostgres struct {
	db *gorm.DB
}

// NewExpenseRepository creates a new ExpensePostgres instance
func NewExpenseRepository(db *DB) *ExpensePostgres {
	return &ExpensePostgres{
		db: db.db,
	}
}

// CreateExpense creates a new model.Expense instance
func (ep *ExpensePostgres) CreateExpense(ctx context.Context, expense *model.Expense) (*model.Expense, error) {
	if err := ep.db.Create(expense).Error; err != nil {
		return nil, err
	}
	return expense, nil
}

// DeleteExpenseByID deletes model.Expense instance by id
func (ep *ExpensePostgres) DeleteExpenseByID(ctx context.Context, id int) error {
	return ep.db.Where("id = ?", id).Delete(model.Expense{}).Error
}

// GetLastExpenses returns last model.Expense instances
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

// GetAllExpensesByPeriod returns all model.Expense instances by period
func (ep *ExpensePostgres) GetAllExpensesByPeriod(ctx context.Context, period times.Period) (int, error) {
	var allExpenses int
	if err := ep.db.Model(new(model.Expense)).Select("SUM(amount)").Where(
		"DATE(created_at) " + getConditionByPeriod(period),
	).Find(&allExpenses).Error; err != nil {
		return 0, err
	}
	return allExpenses, nil
}

// GetBaseExpensesByPeriod returns base model.Expense instances by period
func (ep *ExpensePostgres) GetBaseExpensesByPeriod(ctx context.Context, period times.Period) (int, error) {
	var baseExpenses int
	if err := ep.db.Model(new(model.Expense)).Select("SUM(amount)").Where(
		"DATE(created_at) " + getConditionByPeriod(period) + ` AND 
			category_codename IN (SELECT codename FROM categories WHERE is_base_expense=true)`,
	).Scan(&baseExpenses).Error; err != nil {
		return 0, err
	}
	return baseExpenses, nil
}

func getConditionByPeriod(period times.Period) (condition string) {
	switch period {
	case times.Day:
		condition = "= CURRENT_DATE"
	case times.Month:
		t := time.Now()
		first := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
		condition = fmt.Sprintf(">= '%s'", first.Format("2006-01-02"))
	default:
		panic("unknown period")
	}
	return
}
