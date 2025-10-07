package storage

import (
	"context"

	"github.com/Lemper29/ChatRoom/chat-service/pkg/models"
)

type Storage interface {
	CreateChat(ctx context.Context, req *models.CreateChatRequest) (*models.CreateChatResponse, error)
}
