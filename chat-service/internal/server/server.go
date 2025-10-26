package server

import (
	"context"
	"log/slog"
	"net"

	"github.com/Lemper29/ChatRoom/chat-service/internal/service"
	"github.com/Lemper29/ChatRoom/chat-service/internal/storage"
	pb "github.com/Lemper29/ChatRoom/gen/go/v1/chatroom"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedChatServiceServer
	service *service.Service
	logger  *slog.Logger
	addr    string
}

func NewGrpcServer(addr string, storage storage.Storage, logger *slog.Logger) *Server {
	return &Server{
		addr:    addr,
		logger:  logger,
		service: service.NewService(storage, *logger),
	}
}

func (s *Server) Server() error {
	ctx := context.Background()

	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		s.logger.ErrorContext(ctx, "Failed to listen", "error", err)
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, s)

	s.logger.InfoContext(ctx, "Server start", "port", s.addr)

	if err := grpcServer.Serve(lis); err != nil {
		s.logger.ErrorContext(ctx, "Failed to serve", "error", err)
		return err
	}

	return nil
}

func (s *Server) Connect(stream pb.ChatService_ConnectServer) error {
	return s.service.Connect(stream)
}

func (s *Server) CreateChat(ctx context.Context, req *pb.CreateChatRequest) (*pb.CreateChatResponse, error) {
	return s.service.CreateChat(ctx, req)
}
