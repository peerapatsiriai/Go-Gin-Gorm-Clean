package domain

import (
	entities "go/api/internal/entities"
)

type BookRepository interface {
	Create(book *entities.Book) error
	Update(book *entities.Book) error
	Delete(id uint) error
	FindByID(id uint) (*entities.Book, error)
	FindAll() ([]*entities.Book, error)
}
