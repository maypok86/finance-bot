package main

import (
	"context"
	"flag"
	"log"

	"github.com/maypok86/finance-bot/internal/app/bot"
	"github.com/maypok86/finance-bot/internal/config"
	"github.com/maypok86/finance-bot/internal/logger"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "configs/config.yml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	ctx := context.Background()

	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	c, err := config.Parse(configPath)
	if err != nil {
		return err
	}

	logger.Configure(c.Logger)
	logger.Info(c.Logger)

	b, err := bot.New(ctx, c)
	if err != nil {
		return err
	}

	return b.Start()
}
