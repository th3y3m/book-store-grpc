package service

import (
	"context"
	"errors"
	pb "th3y3m/book-store-grpc/book_service/pb" // Update with the correct path to the generated gRPC code
	repository "th3y3m/book-store-grpc/book_service/repo"
)

type BookService struct {
	pb.UnimplementedBookServiceServer // Embed this struct to implement the interface
	repo                              repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	book := &repository.Book{
		Title:  req.Book.Title,
		Author: req.Book.Author,
	}
	createdBook, err := s.repo.CreateBook(book)
	if err != nil {
		return nil, err
	}

	return &pb.CreateBookResponse{Book: &pb.Book{
		Id:     createdBook.ID,
		Title:  createdBook.Title,
		Author: createdBook.Author,
	}}, nil
}

func (s *BookService) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	book, err := s.repo.GetBook(req.Id)
	if err != nil {
		return nil, errors.New("book not found")
	}
	return &pb.GetBookResponse{Book: &pb.Book{
		Id:     book.ID,
		Title:  book.Title,
		Author: book.Author,
	}}, nil
}

func (s *BookService) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {
	book := &repository.Book{
		ID:     req.Book.Id,
		Title:  req.Book.Title,
		Author: req.Book.Author,
	}
	updatedBook, err := s.repo.UpdateBook(book)
	if err != nil {
		return nil, errors.New("book not found")
	}
	return &pb.UpdateBookResponse{Book: &pb.Book{
		Id:     updatedBook.ID,
		Title:  updatedBook.Title,
		Author: updatedBook.Author,
	}}, nil
}

func (s *BookService) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	success, err := s.repo.DeleteBook(req.Id)
	if err != nil {
		return nil, errors.New("book not found")
	}
	return &pb.DeleteBookResponse{Success: success}, nil
}

func (s *BookService) ListBooks(ctx context.Context, req *pb.Empty) (*pb.ListBooksResponse, error) {
	books, err := s.repo.ListBooks()
	if err != nil {
		return nil, err
	}

	var bookProtos []*pb.Book
	for _, book := range books {
		bookProtos = append(bookProtos, &pb.Book{
			Id:     book.ID,
			Title:  book.Title,
			Author: book.Author,
		})
	}
	return &pb.ListBooksResponse{Books: bookProtos}, nil
}
