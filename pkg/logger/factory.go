package logger

import (
	"skeleton/pkg/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

var loggerLevelMap = map[string]zapcore.Level{
	"debug": zap.DebugLevel,
	"info":  zap.InfoLevel,
	"warn":  zap.WarnLevel,
	"error": zap.ErrorLevel,
	"panic": zap.PanicLevel,
	"fatal": zap.FatalLevel,
}

func NewLogger(env config.Environment, cfg config.LoggingProperties) *Logger {
	var loggerConfig zap.Config
	if env == config.EnvironmentProduction {
		loggerConfig = zap.NewProductionConfig()
	} else {
		loggerConfig = zap.NewDevelopmentConfig()
	}

	level := zap.InfoLevel // Fallback level
	if lvl, ok := loggerLevelMap[cfg.Level]; ok {
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
