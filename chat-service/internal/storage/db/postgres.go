package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Lemper29/ChatRoom/chat-service/internal/storage"
	"github.com/Lemper29/ChatRoom/chat-service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	db *gorm.DB
}

func NewPostgresStorage(config postgres.Config) (storage.Storage, error) {
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

	fmt.Printf("Postgres Database successful")

	return &PostgresDB{db: db}, nil
}

func (p *PostgresDB) CreateChat(ctx context.Context, req *models.CreateChatRequest) (*models.CreateChatResponse, error) {
	// Создаем объект комнаты для БД
	chatRoom := &models.ChatRoom{
		ID:   req.RoomID,
		Name: req.Name,
	}

	// Создаем запись в БД
	if err := p.db.WithContext(ctx).Create(chatRoom).Error; err != nil {
		return nil, fmt.Errorf("failed to create chat room: %w", err)
	}

	// Получаем созданную запись (опционально, но полезно для получения полных данных)
	var createdRoom models.ChatRoom
	if err := p.db.WithContext(ctx).Where("id = ?", req.RoomID).First(&createdRoom).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch created chat room: %w", err)
	}

	// Формируем ответ
	response := &models.CreateChatResponse{
		RoomID: createdRoom.ID,
		Name:   chatRoom.Name,
	}

	return response, nil
}
