package server

import (
	"log"
	"net"

	pb "github.com/Lemper29/ChatRoom/gen/go/v1"
	"google.golang.org/grpc"
)

type server struct {
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
		log.Fatalf("Server err:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}
