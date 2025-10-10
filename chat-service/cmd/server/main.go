package main

import (
	"log"

	"github.com/Lemper29/ChatRoom/chat-service/internal/server"
	"github.com/Lemper29/ChatRoom/chat-service/internal/storage/db"
	"gorm.io/driver/postgres"
)

func main() {
	cfg := postgres.Config{
		DSN:                  "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	}

	storage, err := db.NewPostgresStorage(cfg)
	if err != nil {
		log.Fatalf("DB error: %v", err)
	}

	server := server.NewGrpcServer(":8080", storage)

	if err := server.Server(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
