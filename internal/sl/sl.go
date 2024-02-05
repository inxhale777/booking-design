package sl

import (
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

func New() *Logger {
	s := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(s)

	return &Logger{s}
}

func (l *Logger) Fatal(msg string, args ...any) {
	l.Error(msg, args...)
	os.Exit(1)
}

func Err(err error) slog.Attr {
	return slog.Any("error", err)
}
