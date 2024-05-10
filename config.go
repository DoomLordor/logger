package logger

type Config struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"INFO"`
	LogJson  bool   `env:"LOG_JSON" envDefault:"true"`
	Telegram TelegramConfig
}

type TelegramConfig struct {
	Token     string `env:"LOG_TELEGRAM_TOKEN"`
	ChatID    int64  `env:"LOG_TELEGRAM_CHAT_ID"`
	Namespace string `env:"LOG_TELEGRAM_NAMESPACE" envDefault:"test"`
	Subsystem string `env:"LOG_TELEGRAM_SUBSYSTEM" envDefault:"test"`
}

func (tc TelegramConfig) Check() bool {
	if tc.Token == "" {
		return false
	}
	if tc.ChatID == 0 {
		return false
	}
	return true
}
