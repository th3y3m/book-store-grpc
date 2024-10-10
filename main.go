package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	bookpb "th3y3m/book-store-grpc/book_service/pb"
	userpb "th3y3m/book-store-grpc/user_service/pb"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	// Register BookService
	err := bookpb.RegisterBookServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatalf("Failed to register BookService: %v", err)
	}

	// Register UserService
	err = userpb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost:50052", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatalf("Failed to register UserService: %v", err)
	}

	log.Println("API Gateway is running on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
