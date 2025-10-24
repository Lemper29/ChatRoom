package service

import (
	"context"
	"io"
	"log/slog"
	"time"

	"github.com/Lemper29/ChatRoom/chat-service/internal/storage"
	"github.com/Lemper29/ChatRoom/chat-service/pkg/models"
	pb "github.com/Lemper29/ChatRoom/gen/go/v1/chatroom"
)

type Service struct {
	pb.UnimplementedChatServiceServer
	logger *slog.Logger
	repo   storage.Storage
}

func NewService(repo storage.Storage, logger slog.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: &logger,
	}
}

func (s *Service) Connect(stream pb.ChatService_ConnectServer) error {
	ctx := stream.Context()

	for {
		firstMsg, err := stream.Recv()
		if err == io.EOF {
			s.logger.InfoContext(ctx, "Client closed connection")
			return nil
		}
		if err != nil {
			s.logger.ErrorContext(ctx, "Stream receive error", "error", err)
			return err
		}

		modelsChatMessage := models.NewChatMessage(
			firstMsg.RoomId,
			firstMsg.UserId,
			firstMsg.Username,
			firstMsg.Content,
			firstMsg.Type,
		)

		response := pb.ChatMessage{
			Id:        modelsChatMessage.ID,
			Type:      modelsChatMessage.Type,
			RoomId:    modelsChatMessage.RoomID,
			UserId:    modelsChatMessage.UserID,
			Username:  modelsChatMessage.Username,
			Content:   modelsChatMessage.Content,
			Timestamp: modelsChatMessage.Timestamp,
		}

		if err := stream.Send(&response); err != nil {
			s.logger.ErrorContext(ctx, "Error sending message", "error", err)
			return err
		}
	}
}

func (s *Service) CreateChat(ctx context.Context, req *pb.CreateChatRequest) (*pb.CreateChatResponse, error) {
	room := &storage.ChatRoom{
		ID:          req.RoomId,
		Name:        req.Name,
		Description: req.Description,
		CreatedBy:   req.UserId, // user_id становится created_by
		CreatedAt:   time.Now(),
	}

	createChat, err := s.repo.CreateChat(ctx, &reqCreateChat)
	if err != nil {
		s.logger.ErrorContext(ctx, "Failed to create chat room", "error", err)
		return nil, err
	}

	resCreateChat := pb.CreateChatResponse{
		RoomId: createChat.RoomID,
		Name:   createChat.Name,
	}

	return &resCreateChat, nil
}
