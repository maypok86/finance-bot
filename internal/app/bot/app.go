package bot

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/LazyBearCT/finance-bot/internal/config"
	"github.com/LazyBearCT/finance-bot/internal/telegram"
	"github.com/pkg/errors"
)

type App struct {
	bot *telegram.Bot
}

func New(c *config.Config) (*App, error) {
	bot, err := telegram.New(c.Bot)
	if err != nil {
		return nil, err
	}
	return &App{
		bot: bot,
	}, nil
}

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
