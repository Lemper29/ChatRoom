package db

import (
	"context"
	"log/slog"
	"time"

	"github.com/Lemper29/ChatRoom/chat-service/internal/storage"
	"github.com/Lemper29/ChatRoom/chat-service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewPostgresStorage(config postgres.Config, logger *slog.Logger) (storage.Storage, error) {
	db, err := gorm.Open(postgres.New(config), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	logger.InfoContext(context.Background(), "Database connection successfully")
	return &PostgresDB{db: db, logger: logger}, nil
}

func (p *PostgresDB) CreateChat(ctx context.Context, req *models.CreateChatRequest) (*models.CreateChatResponse, error) {
	// Создаем объект комнаты для БД
	chatRoom := &models.ChatRoom{
		ID:   req.RoomID,
		Name: req.Name,
	}

	// Создаем запись в БД
	if err := p.db.WithContext(ctx).Create(chatRoom).Error; err != nil {
		p.logger.ErrorContext(ctx, "Failed to create chat room", "error", err)
		return nil, err
	}

	// Получаем созданную запись (опционально, но полезно для получения полных данных)
	var createdRoom models.ChatRoom
	if err := p.db.WithContext(ctx).Where("id = ?", req.RoomID).First(&createdRoom).Error; err != nil {
		p.logger.ErrorContext(ctx, "Failed to fetch created chat room", "error", err)
		return nil, err
	}

	// Формируем ответ
	response := &models.CreateChatResponse{
		RoomID: createdRoom.ID,
		Name:   chatRoom.Name,
	}

	p.logger.InfoContext(ctx, "Create chat", "room", response.RoomID, "name", response.Name)
	return response, nil
}
