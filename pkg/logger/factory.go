package logger

import (
	"skeleton/pkg/config"

	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
}

func NewLogger(env config.Environment, cfg *config.LoggingProperties) *Logger {
	if cfg == nil {
		panic("cfg cannot be nil")
	}

	var loggerConfig zap.Config
	if env == config.EnvironmentProduction {
		loggerConfig = zap.NewProductionConfig()
	} else {
		loggerConfig = zap.NewDevelopmentConfig()
	}

	level := zap.InfoLevel // Fallback level
	if lvl, ok := LoggerLevelMap[cfg.Level]; ok {
		level = lvl
	}
	loggerConfig.Level = zap.NewAtomicLevelAt(level)

	logger, err := loggerConfig.Build()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	return &Logger{
		logger,
	}
}
