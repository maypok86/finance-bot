package main

import (
	"flag"
	"log"

	"github.com/LazyBearCT/finance-bot/internal/app/bot"
	"github.com/LazyBearCT/finance-bot/internal/config"
	"github.com/LazyBearCT/finance-bot/internal/logger"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "configs/config.yml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	c, err := config.Parse(configPath)
	if err != nil {
		return err
	}

	logger.Configure(c.Logger)
	logger.Info(c.Logger)

	b, err := bot.New(c)
	if err != nil {
		return err
	}

	return b.Start()
}
