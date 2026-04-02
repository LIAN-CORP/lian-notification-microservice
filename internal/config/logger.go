package config

import (
	"log"

	"go.uber.org/zap"
)

type Logger struct {
	Log *zap.Logger
	Sugar *zap.SugaredLogger
}

func NewLogger() *Logger {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}

	return &Logger{
		Log: zapLogger,
		Sugar: zapLogger.Sugar(),
	}
}