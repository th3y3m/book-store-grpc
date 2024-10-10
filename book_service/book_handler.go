package main

import (
	"log"
	"net"
	"th3y3m/book-store-grpc/book_service/pb"
	repository "th3y3m/book-store-grpc/book_service/repo"
	"th3y3m/book-store-grpc/book_service/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	repo := repository.NewBookRepository()
	bookService := service.NewBookService(repo)

	pb.RegisterBookServiceServer(grpcServer, bookService)

	// Enable reflection for gRPC clients like Postman or grpcurl
	reflection.Register(grpcServer)

	log.Println("gRPC server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
