package configs

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// Config is the struct that holds all the configuration of the application
type Config struct {
	App struct {
		Env      Environment `yaml:"env"`
		PoolSize int         `yaml:"pool_size"`
		Log      struct {
			LogLevel        string   `yaml:"level"`
			PrettyPrint     bool     `yaml:"pretty_print"`
			SentryDSN       string   `yaml:"sentry_dsn"`
			SentryLogLevels []string `yaml:"sentry_log_levels"`
		} `yaml:"log"`
	} `yaml:"app"`

	API struct {
		ServeAddress string `yaml:"address"`
		AccessToken  string `yaml:"access_token"`
	} `yaml:"api"`

	Monitoring struct {
		IsEnabled bool   `yaml:"enable"`
		Path      string `yaml:"path"`
		Password  string `yaml:"password"`
	} `yaml:"monitoring"`

	Notifier struct {
		Squadcast struct {
			IsEnabled bool `yaml:"enable"`
			Teams     map[string]string
		} `yaml:"squadcast"`
		Telegram struct {
			IsEnabled bool                      `yaml:"enable"`
			Host      string                    `yaml:"host"`
			Token     string                    `yaml:"token"`
			Drop      []string                  `yaml:"drop"`
			Teams     map[string][]TelegramTeam `yaml:"teams"`
		} `yaml:"telegram"`
		Slack struct {
			IsEnabled bool `yaml:"enable"`
			Teams     map[string]string
		} `yaml:"slack"`
		Custom struct {
			IsEnabled bool `yaml:"enable"`
			Teams     map[string]string
		} `yaml:"custom"`
	}

	Version string
}

// TelegramTeam is a struct for Telegram team configuration
type TelegramTeam struct {
	Chat  string `yaml:"chat"`
	Topic string `yaml:"topic"`
}

// Load will loads the configuration from the given path
func Load(configPath string) (*Config, error) {
	var cfg Config
	absPath, err := filepath.Abs(configPath)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}

	defer func() {
		if f != nil {
			if err = f.Close(); err != nil {
				log.WithError(err).Error("[SETUP] Error closing file")
			}
		}
	}()
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	v, ok := os.LookupEnv("APP_VERSION")
	if !ok {
		if cfg.App.Env.IsLocal() {
			v = "local"
		} else {
			v = "unknown"
		}
	}

	cfg.Version = v

	return &cfg, err
}
