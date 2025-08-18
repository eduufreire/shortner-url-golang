package logger

import (
	"log/slog"
	"os"
	"sync"
)

var (
	logger *slog.Logger
	sync1  sync.Once
)

func InitLogger() {
	sync1.Do(func() {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	})
}

func Info(msg string) {
	logger.Info(msg)
}
