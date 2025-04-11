package utils

import (
	"go.uber.org/zap"
)

func CreateLogger() *zap.Logger {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	return logger
}
