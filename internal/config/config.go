package config

// Config app
type Config struct {
	Bot    *Bot    `yaml:"bot"`
	Logger *Logger `yaml:"logger"`
	DB     *DB     `yaml:"db"`
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

type DB struct {
	Type           string `yaml:"type" envconfig:"DB_TYPE"`
	Host           string `yaml:"host" envconfig:"DB_HOST"`
	Port           string `yaml:"port" envconfig:"DB_PORT"`
	User           string `yaml:"user" envconfig:"DB_USER"`
	Name           string `yaml:"name" envconfig:"DB_NAME"`
	Password       string `envconfig:"DB_PASSWORD" required:"true"`
	MigrationsPath string `yaml:"migrations_path" envconfig:"DB_MIGRATIONS_PATH"`
}
