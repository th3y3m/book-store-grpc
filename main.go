package main

// import (
// 	"log"
// 	"net"
// 	"th3y3m/book-store-grpc/pb"
// 	repository "th3y3m/book-store-grpc/repo"
// 	"th3y3m/book-store-grpc/service"

// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/reflection"
// )

// func main() {
// 	lis, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		log.Fatalf("Failed to listen: %v", err)
// 	}

// 	grpcServer := grpc.NewServer()
// 	repo := repository.NewBookRepository()
// 	bookService := service.NewBookService(repo)

// 	pb.RegisterBookServiceServer(grpcServer, bookService)

// 	// Enable reflection for gRPC clients like Postman or grpcurl
// 	reflection.Register(grpcServer)

// 	log.Println("gRPC server is running on port 50051...")
// 	if err := grpcServer.Serve(lis); err != nil {
// 		log.Fatalf("Failed to serve: %v", err)
// 	}

// 	// dsn := "host=localhost user=postgres password=12345 dbname=BookStoreDb port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"

// 	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	// if err != nil {
// 	// 	log.Fatalf("Failed to connect to the database: %v", err)
// 	// }

// 	// // Auto-migrate models in the correct order to handle foreign key dependencies
// 	// err = db.AutoMigrate(
// 	// 	&repository.Book{},
// 	// 	&repo.User{},
// 	// )
// 	// if err != nil {
// 	// 	log.Fatalf("Failed to migrate the database: %v", err)
// 	// }

// 	// log.Println("Database migration completed successfully!")
// }
