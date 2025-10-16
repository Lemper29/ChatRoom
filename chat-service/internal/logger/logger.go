package logger

import (
	"log/slog"
	"os"
	"path/filepath"
)

func New(service string, level slog.Level) *slog.Logger {
	opts := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				source, _ := a.Value.Any().(*slog.Source)
				if source != nil {
					source.File = filepath.Base(source.File)
				}
			}
			return a
		},
	}

	var handler slog.Handler = slog.NewTextHandler(os.Stdout, opts)
	if "" == "production" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	logger := slog.New(handler).WithGroup(service)
	return logger
}
