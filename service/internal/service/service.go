package service

import (
	"io"
	"log"

	"github.com/Lemper29/ChatRoom/chat-service/pkg/models"
	pb "github.com/Lemper29/ChatRoom/gen/go/v1"
)

type service struct {
	pb.UnimplementedChatServiceServer
}

func NewService() {
	return
}

func (s *service) Connect(stream pb.ChatService_ConnectServer) error {
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
