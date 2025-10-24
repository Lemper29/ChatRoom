package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/Lemper29/ChatRoom/chat-gateway/internal/utils"
	"github.com/Lemper29/ChatRoom/chat-gateway/pkg/models"
	pb "github.com/Lemper29/ChatRoom/gen/go/v1/chatroom"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Handler struct {
	chatRoom pb.ChatServiceClient
}

func NewHandler() *Handler {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return &Handler{
		chatRoom: pb.NewChatServiceClient(conn),
	}
}

func (h *Handler) CreateChat(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var payload models.CreateChatRequest
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
	}

	req := pb.CreateChatRequest{
		RoomId:      payload.RoomID,
		UserId:      payload.UserID,
		Name:        payload.Name,
		Description: payload.Description,
	}

	resp, err := h.chatRoom.CreateChat(ctx, &req)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
	}

	utils.WriteJSON(w, http.StatusOK, resp)
}
