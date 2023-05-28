package configs

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Env Environment `yaml:"env"`
		Log struct {
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

	Notifier struct {
		Squadcast struct {
			IsEnabled bool `yaml:"enable"`
			Teams     map[string]string
		} `yaml:"squadcast"`
	}
}

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
				log.WithError(err).Error("error closing file")
			}
		}
	}()
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	log.Info(cfg.Notifier.Squadcast)

	return &cfg, err
}
