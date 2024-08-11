package logger

import (
	"log"
	"log/slog"
	"os"
)


func NewLogger() *slog.Logger {
	opts := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	file, err := os.OpenFile("authErrors.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger := slog.New(slog.NewTextHandler(file, &opts))

	return logger
}
