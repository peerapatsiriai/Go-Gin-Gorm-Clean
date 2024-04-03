package domain

import (
	entities "go/api/internal/entities"
)

type BookUseCase interface {
	CreateBook(book *entities.Book) error
	UpdateBook(id uint, book *entities.Book) error
	DeleteBook(id uint) error
	GetBookByID(id uint) (*entities.Book, error)
	GetAllBooks() ([]*entities.Book, error)
}

type BookRepository interface {
	Create(book *entities.Book) error
	Update(book *entities.Book) error
	Delete(id uint) error
	FindByID(id uint) (*entities.Book, error)
	FindAll() ([]*entities.Book, error)
}
