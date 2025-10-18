package service

import (
	"bytes"
	"context"
	"log/slog"
	"testing"

	"github.com/Lemper29/ChatRoom/chat-service/internal/storage"
	"github.com/Lemper29/ChatRoom/chat-service/pkg/models"
	pb "github.com/Lemper29/ChatRoom/gen/go/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateChat(ctx context.Context, req *models.CreateChatRequest) (*models.CreateChatResponse, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.CreateChatResponse), args.Error(1)
}

// Вспомогательная функция для создания сервиса в тестах
func newTestService(repo storage.Storage) *Service {
	logger := slog.New(slog.NewTextHandler(&bytes.Buffer{}, nil))
	return NewService(repo, *logger)
}

func TestCreateChat_Success(t *testing.T) {
	// Создаем mock репозитория
	mockRepo := new(MockRepository)
	s := newTestService(mockRepo)

	// Подготавливаем тестовые данные
	testReq := &pb.CreateChatRequest{
		RoomId:      "room-123",
		UserId:      "user-456",
		Name:        "Test Chat",
		Description: "Test Description",
	}

	expectedRepoReq := &models.CreateChatRequest{
		RoomID:      testReq.RoomId,
		UserID:      testReq.UserId,
		Name:        testReq.Name,
		Description: testReq.Description,
	}

	expectedResponse := &models.CreateChatResponse{
		RoomID: "room-123",
		Name:   "Test Chat",
	}

	// Настраиваем ожидания для mock
	mockRepo.On("CreateChat", mock.Anything, expectedRepoReq).Return(expectedResponse, nil)

	// Вызываем тестируемый метод
	ctx := context.Background()
	res, err := s.CreateChat(ctx, testReq)

	// Проверяем результаты
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "room-123", res.RoomId)
	assert.Equal(t, "Test Chat", res.Name)

	// Проверяем что mock был вызван с правильными параметрами
	mockRepo.AssertCalled(t, "CreateChat", mock.Anything, expectedRepoReq)
	mockRepo.AssertExpectations(t)
}

func TestCreateChat_RepositoryError(t *testing.T) {
	// Создаем mock репозитория
	mockRepo := new(MockRepository)
	s := newTestService(mockRepo)

	testReq := &pb.CreateChatRequest{
		RoomId:      "room-123",
		UserId:      "user-456",
		Name:        "Test Chat",
		Description: "Test Description",
	}

	expectedRepoReq := &models.CreateChatRequest{
		RoomID:      testReq.RoomId,
		UserID:      testReq.UserId,
		Name:        testReq.Name,
		Description: testReq.Description,
	}

	// Настраиваем mock на возврат ошибки
	mockRepo.On("CreateChat", mock.Anything, expectedRepoReq).Return((*models.CreateChatResponse)(nil), assert.AnError)

	// Вызываем тестируемый метод
	ctx := context.Background()
	res, err := s.CreateChat(ctx, testReq)

	// Проверяем что получили ошибку
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, assert.AnError, err) // Проверяем конкретную ошибку

	mockRepo.AssertExpectations(t)
}

func TestNewService(t *testing.T) {
	mockRepo := new(MockRepository)
	service := newTestService(mockRepo)

	assert.NotNil(t, service)
	assert.Equal(t, mockRepo, service.repo)
}
