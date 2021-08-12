package config

// Config app
type Config struct {
	Bot    *Bot
	Logger *Logger `yaml:"logger"`
}

// Logger config
type Logger struct {
	Level string `yaml:"level" env:"LOGGER_LEVEL"`
	File  string `yaml:"file" env:"LOGGER_FILE"`
}

// Bot config
type Bot struct {
	BotToken string `envconfig:"BOT_TOKEN" required:"true"`
	AccessID int    `envconfig:"ACCESS_ID" required:"true"`
}
