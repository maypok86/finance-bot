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

	Message Message `yaml:"message"`
}

type DB struct {
	Type           string `yaml:"type" env:"DB_TYPE"`
	Host           string `yaml:"host" env:"DB_HOST"`
	Port           string `yaml:"port" env:"DB_PORT"`
	User           string `yaml:"user" env:"DB_USER"`
	Name           string `yaml:"name" env:"DB_NAME"`
	Password       string `env:"DB_PASSWORD" required:"true"`
	MigrationsPath string `yaml:"migrations_path" env:"DB_MIGRATIONS_PATH"`
}

type Message struct {
	Response Response `yaml:"response"`
	Error    Error    `yaml:"error"`
}

type Response struct {
	Start      string `yaml:"start"`
	Delete     string `yaml:"delete"`
	Categories string `yaml:"categories"`
	Unknown    string `yaml:"unknown"`
}

type Error struct {
	Unknown string `yaml:"unknown"`
}
