package application

import (
	"time"

	"github.com/arvancloud/uptime-webhook/configs"
	"github.com/evalphobia/logrus_sentry"
	log "github.com/sirupsen/logrus"
)

func SetupLogger(config *configs.Config) error {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat:  time.RFC3339,
		DisableTimestamp: false,
		DataKey:          "",
		FieldMap:         nil,
		CallerPrettyfier: nil,
		PrettyPrint:      config.App.Log.PrettyPrint,
	})

	log.SetLevel(parseLogLevel(config.App.Log.LogLevel))
	sentryLogLevels := make([]log.Level, len(config.App.Log.SentryLogLevels))
	i := 0
	for _, level := range config.App.Log.SentryLogLevels {
		sentryLogLevels[i] = parseLogLevel(level)
		i++
	}

	sentryHook, err := logrus_sentry.NewAsyncSentryHook(config.App.Log.SentryDSN, sentryLogLevels)
	if err != nil {
		return err
	}

	sentryHook.Timeout = time.Second * 10
	sentryHook.SetEnvironment(string(config.App.Env))
	sentryHook.StacktraceConfiguration.Enable = true
	sentryHook.SetRelease(Version)
	log.AddHook(sentryHook)
	return nil
}

func parseLogLevel(level string) log.Level {
	if level == "" {
		return log.InfoLevel
	}
	levels := []string{"panic", "fatal", "error", "warn", "info", "debug", "trace"}
	for index, lvl := range levels {
		if lvl == level {
			return log.Level(index)
		}
	}
	return log.InfoLevel
}
