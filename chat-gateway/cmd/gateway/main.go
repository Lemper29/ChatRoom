package main

import (
	"context"
	"log"
	"net/http"

	pb "github.com/Lemper29/ChatRoom/gen/go/v1/chatroom"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := pb.RegisterChatServiceHandlerFromEndpoint(ctx, mux, "localhost:8080", opts)
	if err != nil {
		panic(err)
	}

	log.Printf("server listening at 8081")

	if err := http.ListenAndServe(":8081", mux); err != nil {
		panic(err)
	}
}
