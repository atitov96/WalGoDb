package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(level string, output string) (*zap.Logger, error) {
	config := zap.NewProductionConfig()

	logLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		return nil, err
	}
	config.Level = zap.NewAtomicLevelAt(logLevel)

	if output != "stdout" {
		config.OutputPaths = []string{output}
	}

	logger, err := config.Build()
	return logger, err
}
