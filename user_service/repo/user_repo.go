package repository

import (
	"th3y3m/book-store-grpc/util"

	"gorm.io/gorm"
)

// User represents the model for a book in the database.
type User struct {
	ID   int32  `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"size:255"`
}

// UserRepository defines the methods for interacting with books in the database.
type UserRepository interface {
	ListUsers() ([]*User, error)
}

// bookRepository implements the UserRepository interface with GORM.
type bookRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new book repository with a connected database instance.
func NewUserRepository() UserRepository {
	db, err := util.ConnectToPostgreSQL()
	if err != nil {
		panic(err)
	}
	return &bookRepository{
		db: db,
	}
}

func (r *bookRepository) ListUsers() ([]*User, error) {
	var books []*User
	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
