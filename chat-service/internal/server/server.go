package server

import (
	"fmt"
	"net"

	"github.com/Lemper29/ChatRoom/chat-service/internal/service"
	"github.com/Lemper29/ChatRoom/chat-service/internal/storage"
	pb "github.com/Lemper29/ChatRoom/gen/go/v1"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedChatServiceServer
	service *service.Service
	addr    string
}

func NewGrpcServer(addr string, storage storage.Storage) *Server {
	return &Server{
		addr:    addr,
		service: service.NewService(storage),
	}
}

func (s *Server) Server() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, s)

	fmt.Printf("Server starting on %s\n", s.addr)

	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
