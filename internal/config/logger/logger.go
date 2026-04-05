package logger

import (
	"log"

	"go.uber.org/zap"
)

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger() *Logger {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}

	return &Logger{
		SugaredLogger: zapLogger.Sugar(),
	}
}