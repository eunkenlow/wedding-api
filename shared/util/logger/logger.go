package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New configures global zap logger
func New() *zap.Logger {
	// Get config
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := config.Build()
	// Replace global zap logger
	zap.ReplaceGlobals(logger)
	return logger
}
