package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/LazyBearCT/finance-bot/internal/model"
	"github.com/LazyBearCT/finance-bot/pkg/random"
	"github.com/LazyBearCT/finance-bot/pkg/times"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

const (
	countExpenses = 20
)

func getRandomCategoryCodename() string {
	index := random.IntByRange(0, int64(len(categories))-1)
	return categories[index].Codename
}

func containsExpense(t *testing.T, expected []*model.Expense, actual *model.Expense) bool {
	t.Helper()
	for _, expense := range expected {
		if expense.ID == actual.ID {
			require.Equal(t, expense.ID, actual.ID)
			require.Equal(t, expense.Amount, actual.Amount)
			return true
		}
	}
	return false
}

func createRandomExpense(t *testing.T, e *ExpensePostgres) *model.Expense {
	t.Helper()
	expected := &model.Expense{
		ID:               random.Int(),
		Amount:           random.Int(),
		CategoryCodename: getRandomCategoryCodename(),
		CreatedAt:        random.Timestamp(),
		RawText:          random.Name(),
	}
	actual, err := e.CreateExpense(context.Background(), expected)
	require.NoError(t, err)
	require.NotEmpty(t, actual)

	require.Equal(t, expected.Amount, actual.Amount)
	require.Equal(t, expected.CategoryCodename, actual.CategoryCodename)
	require.Equal(t, expected.RawText, actual.RawText)

	require.NotZero(t, actual.ID)
	require.NotZero(t, actual.CreatedAt)

	return actual
}

func truncate(t *testing.T, repo *ExpensePostgres) {
	t.Helper()
	require.NoError(t, repo.db.Exec("TRUNCATE TABLE expenses").Error)
}

func createRandomExpenses(t *testing.T, e *ExpensePostgres, count int) []*model.Expense {
	t.Helper()
	expenses := make([]*model.Expense, 0, count)
	for i := 0; i < count; i++ {
		expenses = append(expenses, createRandomExpense(t, e))
	}
	return expenses
}

func TestCreateExpense(t *testing.T) {
	repo := NewExpenseRepository(db)
	createRandomExpense(t, NewExpenseRepository(db))
	truncate(t, repo)
}

func TestDeleteExpenseByID(t *testing.T) {
	repo := NewExpenseRepository(db)
	expected := createRandomExpense(t, repo)
	err := repo.DeleteExpenseByID(context.Background(), expected.ID)
	require.NoError(t, err)

	var actual *model.Expense
	err = repo.db.Where("id = ?", expected.ID).First(&actual).Error
	require.Error(t, err)
	require.EqualError(t, err, gorm.ErrRecordNotFound.Error())
	require.Empty(t, actual)
	truncate(t, repo)
}

func TestGetLastExpenses(t *testing.T) {
	repo := NewExpenseRepository(db)
	expected := createRandomExpenses(t, repo, countExpenses)
	actual, err := repo.GetLastExpenses(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, actual)

	for i := 0; i < len(actual); i++ {
		require.NotEmpty(t, actual[i])
		require.True(t, containsExpense(t, expected, actual[i]))
	}
	truncate(t, repo)
}

func date(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}

func getByPeriod(date time.Time, t time.Time, period times.Period) bool {
	switch period {
	case times.Day:
		return date == t
	case times.Month:
		return date.After(t)
	}
	return false
}

func getSliceByPeriod(t *testing.T, repo *ExpensePostgres, period times.Period, count int) []*model.Expense {
	t.Helper()
	var condTime time.Time
	x := time.Now()
	switch period {
	case times.Day:
		condTime = time.Date(x.Year(), x.Month(), x.Day(), 0, 0, 0, 0, time.UTC)
	case times.Month:
		condTime = time.Date(x.Year(), x.Month(), 1, 0, 0, 0, 0, time.UTC)
	default:
		panic("unknown period")
	}
	expenses := createRandomExpenses(t, repo, count)
	ans := make([]*model.Expense, 0, len(expenses))
	for _, expense := range expenses {
		if getByPeriod(date(expense.CreatedAt), condTime, period) {
			ans = append(ans, expense)
		}
	}
	return ans
}

func sumAmount(expenses []*model.Expense) int {
	var sum int
	for _, expense := range expenses {
		sum += expense.Amount
	}
	return sum
}

func getAllExpensesByPeriod(t *testing.T, period times.Period) {
	t.Helper()
	repo := NewExpenseRepository(db)
	expenses := getSliceByPeriod(t, repo, period, countExpenses)
	expected := sumAmount(expenses)

	actual, _ := repo.GetAllExpensesByPeriod(context.Background(), period)
	require.Equal(t, expected, actual)
	truncate(t, repo)
}

func TestGetAllExpensesByPeriod(t *testing.T) {
	t.Run("day", func(t *testing.T) {
		getAllExpensesByPeriod(t, times.Day)
	})
	t.Run("month", func(t *testing.T) {
		getAllExpensesByPeriod(t, times.Month)
	})
}

func containsCodename(expense *model.Expense) bool {
	for _, category := range categories {
		if expense.CategoryCodename == category.Codename && category.IsBaseExpense {
			return true
		}
	}
	return false
}

func getBaseExpensesByPeriod(t *testing.T, period times.Period) {
	t.Helper()
	repo := NewExpenseRepository(db)
	sliceExpenses := getSliceByPeriod(t, repo, period, countExpenses)
	expenses := make([]*model.Expense, 0, len(sliceExpenses))
	for _, expense := range sliceExpenses {
		if containsCodename(expense) {
			expenses = append(expenses, expense)
		}
	}
	expected := sumAmount(expenses)

	actual, _ := repo.GetBaseExpensesByPeriod(context.Background(), period)
	require.Equal(t, expected, actual)
	truncate(t, repo)
}

func TestGetBaseExpensesByPeriod(t *testing.T) {
	t.Run("day", func(t *testing.T) {
		getBaseExpensesByPeriod(t, times.Day)
	})
	t.Run("month", func(t *testing.T) {
		getBaseExpensesByPeriod(t, times.Month)
	})
}
