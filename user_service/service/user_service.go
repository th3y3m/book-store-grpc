package service

import (
	"context"
	"log"
	pb1 "th3y3m/book-store-grpc/book_service/pb" // Update with the correct path to the generated gRPC code
	pb "th3y3m/book-store-grpc/user_service/pb"  // Update with the correct path to the generated gRPC code
	repository "th3y3m/book-store-grpc/user_service/repo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserService struct {
	pb.UnimplementedUserServiceServer // Embed this struct to implement the interface
	repo                              repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) ListUsers(ctx context.Context, req *pb.EmptyUser) (*pb.ListUsersResponse, error) {

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to BookService: %v", err)
	}

	bookClient := pb1.NewBookServiceClient(conn)

	bookResp, err := bookClient.ListBooks(ctx, &pb1.Empty{})
	if err != nil {
		log.Printf("Failed to list books: %v", err)
	} else {
		log.Printf("Books: %v", bookResp.Books)
	}

	users, err := s.repo.ListUsers()
	if err != nil {
		return nil, err
	}

	var userProtos []*pb.User
	for _, user := range users {
		userProtos = append(userProtos, &pb.User{
			Id:   user.ID,
			Name: user.Name,
		})
	}
	return &pb.ListUsersResponse{Users: userProtos}, nil
}

// func (s *UserService) ListUsers(ctx context.Context, req *pb.EmptyUser) (*pb.ListUsersResponse, error) {

// 	conn, err := grpc.DialContext(ctx, "localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

// 	if err != nil {
// 		log.Fatalf("Failed to connect to BookService: %v", err)
// 	}
// 	defer conn.Close()

// 	bookClient := pb1.NewBookServiceClient(conn)

// 	bookResp, err := bookClient.ListBooks(ctx, &pb1.Empty{})
// 	if err != nil {
// 		log.Printf("Failed to list books: %v", err)
// 	} else {
// 		log.Printf("Books: %v", bookResp.Books)
// 	}

// 	users, err := s.repo.ListUsers()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var userProtos []*pb.User
// 	for _, user := range users {
// 		userProtos = append(userProtos, &pb.User{
// 			Id:   user.ID,
// 			Name: user.Name,
// 		})
// 	}
// 	return &pb.ListUsersResponse{Users: userProtos}, nil
// }
