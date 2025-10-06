package models

import (
	"fmt"
	"time"

	pb "github.com/Lemper29/ChatRoom/gen/go/v1"
)

type ChatMessage struct {
	ID        string
	Type      pb.MessageType
	RoomID    string
	UserID    string
	Username  string
	Content   string
	Timestamp int64
}

// Создание нового сообщения
func NewChatMessage(roomID, userID, username, content string, msgType pb.MessageType) *ChatMessage {
	return &ChatMessage{
		ID:        generateID(),
		Type:      msgType,
		RoomID:    roomID,
		UserID:    userID,
		Username:  username,
		Content:   content,
		Timestamp: time.Now().Unix(),
	}
}

// Генерация ID (упрощенная версия)
func generateID() string {
	return time.Now().Format("20060102150405") + "-" + fmt.Sprintf("%d", time.Now().UnixNano())
}
