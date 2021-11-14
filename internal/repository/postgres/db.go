package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/maypok86/finance-bot/internal/config"
	"github.com/maypok86/finance-bot/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB Postgres.
type DB struct {
	ctx context.Context
	dsn string
	db  *gorm.DB
}

// New creates a new DB instance.
func New(ctx context.Context, config *config.DB) (*DB, error) {
	db := &DB{
		ctx: ctx,
		dsn: fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.Host, config.Port, config.User, config.Password, config.Name,
		),
	}
	if err := db.Connect(); err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("postgres://%s:%s/%s?sslmode=disable&user=%s&password=%s",
		config.Host, config.Port, config.Name, config.User, config.Password)
	if err := runMigrations(config.MigrationsPath, dsn); err != nil {
		return nil, err
	}
	return db, nil
}

// Connect to DB.
func (db *DB) Connect() (err error) {
	db.db, err = gorm.Open(postgres.Open(db.dsn), new(gorm.Config))
	return
}

const keepAlivePollPeriod = 10

// KeepAlive ...
func (db *DB) KeepAlive() {
	for {
		// Check if PostgreSQL is alive every 10 seconds
		time.Sleep(time.Second * keepAlivePollPeriod)
		database, err := db.db.DB()
		if err == nil && database.PingContext(db.ctx) == nil {
			continue
		}
		logger.Info("Lost PostgreSQL connection. Restoring...")
		if err := db.Connect(); err != nil {
			logger.Error(err)
			continue
		}
		logger.Info("PostgreSQL reconnected")
	}
}
