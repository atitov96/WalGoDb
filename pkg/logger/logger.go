package logger

import "go.uber.org/zap"

func NewLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	return logger, err
}
