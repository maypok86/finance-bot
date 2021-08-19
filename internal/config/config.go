package config

// Config app.
type Config struct {
	Bot    *Bot    `yaml:"bot"`
	Logger *Logger `yaml:"logger"`
	DB     *DB     `yaml:"db"`
}

// Logger config.
type Logger struct {
	Level string `yaml:"level" envconfig:"LOGGER_LEVEL"`
	File  string `yaml:"file" envconfig:"LOGGER_FILE"`
}

// Bot config.
type Bot struct {
	BotToken string `envconfig:"BOT_TOKEN"`
	AccessID int    `envconfig:"ACCESS_ID"`
}

// DB config.
type DB struct {
	Type           string `yaml:"type" envconfig:"DB_TYPE"`
	Host           string `yaml:"host" envconfig:"DB_HOST"`
	Port           string `yaml:"port" envconfig:"DB_PORT"`
	User           string `yaml:"user" envconfig:"DB_USER"`
	Name           string `yaml:"name" envconfig:"DB_NAME"`
	Password       string `yaml:"password" envconfig:"DB_PASSWORD"`
	MigrationsPath string `yaml:"migrations_path" envconfig:"DB_MIGRATIONS_PATH"`
}
