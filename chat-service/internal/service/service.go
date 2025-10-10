package service

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/Lemper29/ChatRoom/chat-service/internal/storage"
	"github.com/Lemper29/ChatRoom/chat-service/pkg/models"
	pb "github.com/Lemper29/ChatRoom/gen/go/v1"
)

type Service struct {
	pb.UnimplementedChatServiceServer
	repo storage.Storage
}

func NewService(repo storage.Storage) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Connect(stream pb.ChatService_ConnectServer) error {
	for {
		firstMsg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Printf("Received from %s: %s in RoomId %s", firstMsg.GetUsername(), firstMsg.GetContent(), firstMsg.GetRoomId())

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
			log.Printf("Error sending message: %v", err)
			return err
		}
	}
}

func (s *Service) CreateChat(ctx context.Context, req *pb.CreateChatRequest) (*pb.CreateChatResponse, error) {
	reqCreateChat := models.CreateChatRequest{
		RoomID:      req.RoomId,
		UserID:      req.UserId,
		Name:        req.Name,
		Description: req.Description,
	}

	createChat, err := s.repo.CreateChat(ctx, &reqCreateChat)
	if err != nil {
		return nil, fmt.Errorf("failed to create chat room: %w", err)
	}

	resCreateChat := pb.CreateChatResponse{
		RoomId: createChat.RoomID,
		Name:   createChat.Name,
	}

	return &resCreateChat, nil
}
