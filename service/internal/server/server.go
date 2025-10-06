package server

import (
	"fmt"
	"net"

	pb "github.com/Lemper29/ChatRoom/gen/go/v1"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedChatServiceServer
	addr string
}

func NewGrpcServer(addr string) *server {
	return &server{
		addr: addr,
	}
}

func (s *server) Server() error {
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
