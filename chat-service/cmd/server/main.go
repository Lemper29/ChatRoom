package main

import (
	"context"
	"log/slog"

	"github.com/Lemper29/ChatRoom/chat-service/internal/logger"
	"github.com/Lemper29/ChatRoom/chat-service/internal/server"
	"github.com/Lemper29/ChatRoom/chat-service/internal/storage/db"
	"gorm.io/driver/postgres"
)

func main() {
	ctx := context.Background()
	logger := logger.New("chat-service", slog.LevelDebug)

	cfg := postgres.Config{
		DSN:                  "host=localhost user=postgres password=5430 dbname=postgres port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	}

	storage, err := db.NewPostgresStorage(cfg, logger)
	if err != nil {
		logger.ErrorContext(ctx, "Database error", "error", err)
	}

	server := server.NewGrpcServer(":8080", storage, logger)

	if err := server.Server(); err != nil {
		logger.ErrorContext(ctx, "Server error", "error", err)
	}
}
