package bot

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/maypok86/finance-bot/internal/config"
	"github.com/maypok86/finance-bot/internal/repository"
	"github.com/maypok86/finance-bot/internal/service"
	"github.com/maypok86/finance-bot/internal/telegram"
	"github.com/pkg/errors"
)

// App telegram bot.
type App struct {
	bot *telegram.Bot
}

// New create new telegram bot app.
func New(ctx context.Context, c *config.Config) (*App, error) {
	repo, err := repository.New(ctx, c.DB)
	if err != nil {
		return nil, err
	}

	manager, err := service.NewManager(ctx, repo)
	if err != nil {
		return nil, err
	}

	bot, err := telegram.New(c.Bot, manager)
	if err != nil {
		return nil, err
	}

	return &App{
		bot: bot,
	}, nil
}

// Start telegram bot app.
func (a *App) Start() error {
	eChan := make(chan error)
	quit := make(chan os.Signal, 1)

	go func() {
		if err := a.bot.Start(); err != nil {
			eChan <- err
		}
	}()

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case err := <-eChan:
		return errors.Wrap(err, "bot start failed")
	case <-quit:
	}

	return nil
}
