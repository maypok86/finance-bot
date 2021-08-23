package postgres

import (
	"context"
	"log"
	"os"
	"testing"

	"gitlab.com/LazyBearCT/finance-bot/internal/config"
	"gitlab.com/LazyBearCT/finance-bot/internal/model"
)

var (
	db         *DB
	categories []*model.DBCategory
)

func TestMain(m *testing.M) {
	c := new(config.Config)
	if err := c.ParseFile("../../../configs/test.yml"); err != nil {
		log.Fatal("cannot load config:", err)
	}

	var err error
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
