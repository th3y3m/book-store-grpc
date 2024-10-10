package repository

import (
	"errors"
	"th3y3m/book-store-grpc/util"

	"gorm.io/gorm"
)

// Book represents the model for a book in the database.
type Book struct {
	ID     int32  `gorm:"primaryKey;autoIncrement"`
	Title  string `gorm:"size:255"`
	Author string `gorm:"size:255"`
}

// BookRepository defines the methods for interacting with books in the database.
type BookRepository interface {
	CreateBook(book *Book) (*Book, error)
	GetBook(id int32) (*Book, error)
	UpdateBook(book *Book) (*Book, error)
	DeleteBook(id int32) (bool, error)
	ListBooks() ([]*Book, error)
}

// bookRepository implements the BookRepository interface with GORM.
type bookRepository struct {
	db *gorm.DB
}

// NewBookRepository creates a new book repository with a connected database instance.
func NewBookRepository() BookRepository {
	db, err := util.ConnectToPostgreSQL()
	if err != nil {
		panic(err)
	}
	return &bookRepository{
		db: db,
	}
}

// CreateBook adds a new book to the database.
func (r *bookRepository) CreateBook(book *Book) (*Book, error) {
	if err := r.db.Create(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

// GetBook retrieves a book from the database by its ID.
func (r *bookRepository) GetBook(id int32) (*Book, error) {
	var book Book
	if err := r.db.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}
	return &book, nil
}

// UpdateBook updates an existing book in the database.
func (r *bookRepository) UpdateBook(book *Book) (*Book, error) {
	if err := r.db.Save(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

// DeleteBook removes a book from the database by its ID.
func (r *bookRepository) DeleteBook(id int32) (bool, error) {
	result := r.db.Delete(&Book{}, id)
	if result.RowsAffected == 0 {
		return false, errors.New("book not found")
	}
	return true, nil
}

// ListBooks retrieves all books from the database.
func (r *bookRepository) ListBooks() ([]*Book, error) {
	var books []*Book
	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
