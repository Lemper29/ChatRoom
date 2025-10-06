package main

import (
	"log"

	"github.com/Lemper29/ChatRoom/chat-service/internal/server"
)

func main() {
	server := server.NewGrpcServer("localhost:8080")

	if err := server.Server(); err != nil {
		log.Fatalf("Server err: %v", err)
	}
}
