package models

type CreateChatRequest struct {
	RoomID      string
	UserID      string
	Name        string
	Description string
	CreatedBy   string
}
