package postgres

import (
	"context"
	"github.com/LazyBearCT/finance-bot/internal/config"
	"log"
	"os"
	"testing"
)

var (
	db *DB
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

	os.Exit(m.Run())
}
