package logging

import (
	"log/slog"
	"os"
)

var DEBUG = true

type Logger struct {
	History []string
}

var (
	instance *Logger = nil
)

func GetInstance() *Logger {
	if instance != nil {
		return instance
	} else {
		return nil
	}
}

func InitLogger() error {
	var handler slog.Handler

	if DEBUG == true {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)

	return nil
}

func SaveCrashReport() {

}
