package main

import (
	"context"
	"log/slog"

	"github.com/Lemper29/ChatRoom/chat-service/config"
	"github.com/Lemper29/ChatRoom/chat-service/internal/logger"
	"github.com/Lemper29/ChatRoom/chat-service/internal/server"
	"github.com/Lemper29/ChatRoom/chat-service/internal/storage/db"
	"gorm.io/driver/postgres"
)

func main() {
	cfg := config.Envs
	ctx := context.Background()

	logger := logger.New("chat-service", slog.LevelDebug)

	cfgPostgres := postgres.Config{
		DSN:                  cfg.DSN,
		PreferSimpleProtocol: true,
	}

	storage, err := db.NewPostgresStorage(cfgPostgres, logger)
	if err != nil {
		logger.ErrorContext(ctx, "Database error", "error", err)
	}

	server := server.NewGrpcServer(":"+cfg.Port, storage, logger)

	if err := server.Server(); err != nil {
		logger.ErrorContext(ctx, "Server error", "error", err)
	}
}
