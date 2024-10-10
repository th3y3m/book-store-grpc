package main

import (
	"log"
	"net"
	"th3y3m/book-store-grpc/user_service/pb"
	repository "th3y3m/book-store-grpc/user_service/repo"
	"th3y3m/book-store-grpc/user_service/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	repo := repository.NewUserRepository()
	bookService := service.NewUserService(repo)

	pb.RegisterUserServiceServer(grpcServer, bookService)

	// Enable reflection for gRPC clients like Postman or grpcurl
	reflection.Register(grpcServer)

	log.Println("gRPC server is running on port 50052...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
