package logger

import (
	"log/slog"
	"os"
)

// SetupLogger - инициализирует логгер из пакета slog с выводом в формате JSON и устанавливает
// его логгером по умолчанию. Это позволяет не передавать этот кастомный логгер другим объектам.
func SetupLogger() {
	log := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
	slog.SetDefault(log)
}

// Err - обертка для ошибки, представляет ее как атрибут логгера.
func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
