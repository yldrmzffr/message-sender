package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"message-sender/config"
	"sync"
)

var (
	Log  *zap.Logger
	once sync.Once
)

func InitLogger(cfg *config.LogConfig) {
	once.Do(func() {
		zapConfig := zap.NewProductionConfig()

		level, err := zapcore.ParseLevel(cfg.Level)
		if err != nil {
			level = zapcore.DebugLevel
		}

		zapConfig.Level.SetLevel(level)

		zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

		logger, err := zapConfig.Build()
		if err != nil {
			panic(err)
		}

		Log = logger
	})
}

func GetLogger() *zap.Logger {
	if Log == nil {
		InitLogger(&config.LogConfig{})
	}

	return Log
}

func ToZapField(field interface{}) zap.Field {
	switch f := field.(type) {
	case error:
		return zap.Error(f)
	case zap.Field:
		return f
	default:
		return zap.Any("field", f)
	}
}

func convertFields(fields ...interface{}) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for i, field := range fields {
		zapFields[i] = ToZapField(field)
	}

	return zapFields
}

func Info(message string, fields ...interface{}) {
	GetLogger().Info(message, convertFields(fields...)...)
}

func Debug(message string, fields ...interface{}) {
	GetLogger().Debug(message, convertFields(fields...)...)
}

func Warn(message string, fields ...interface{}) {
	GetLogger().Warn(message, convertFields(fields...)...)
}

func Error(message string, fields ...interface{}) {
	GetLogger().Error(message, convertFields(fields...)...)
}
