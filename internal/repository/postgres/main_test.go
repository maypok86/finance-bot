package postgres

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/LazyBearCT/finance-bot/internal/config"
	"github.com/LazyBearCT/finance-bot/internal/model"
)

var (
	db         *DB
	categories []*model.DBCategory
)

func TestMain(m *testing.M) {
	c, err := config.Parse("../../../configs/test.yml")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	db, err = New(context.Background(), c.DB)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	categories, err = NewCategoryRepository(db).GetAllCategories(context.Background())
	if err != nil {
		log.Fatal("cannot get all categories:", err)
	}

	os.Exit(m.Run())
}
