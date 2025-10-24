package models

import (
	"fmt"
	"time"

	pb "github.com/Lemper29/ChatRoom/gen/go/v1/chatroom"
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

type ChatRoom struct {
	ID          string    `gorm:"primaryKey;type:varchar(255)"`
	Name        string    `gorm:"type:text;not null"`
	Description string    `gorm:"type:text"`
	CreatedBy   string    `gorm:"type:varchar(255);not null"` // Убедитесь что это поле заполняется
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type CreateChatRequest struct {
	RoomID      string
	UserID      string
	Name        string
	Description string
}

type CreateChatResponse struct {
	RoomID string
	Name   string
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
