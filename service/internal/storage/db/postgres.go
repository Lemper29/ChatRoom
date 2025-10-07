package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Lemper29/ChatRoom/chat-service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	db *gorm.DB
}

func NewPostgresStorage(config postgres.Config) (*gorm.DB, error) {
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

	return db, nil
}

func (p *PostgresDB) CreateChat(ctx context.Context, req *models.CreateChatRequest) (*models.CreateChatResponse, error)
