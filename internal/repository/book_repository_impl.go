package repository

import (
	"go/api/internal/domain"
	entities "go/api/internal/entities"

	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) domain.BookRepository {
	return &BookRepositoryImpl{
		db: db,
	}
}

func (r *BookRepositoryImpl) Create(book *entities.Book) error {
	return r.db.Create(book).Error
}

func (r *BookRepositoryImpl) Update(book *entities.Book) error {
	return r.db.Save(book).Error
}

func (r *BookRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entities.Book{}, id).Error
}

func (r *BookRepositoryImpl) FindByID(id uint) (*entities.Book, error) {
	var book entities.Book
	if err := r.db.First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepositoryImpl) FindAll() ([]*entities.Book, error) {
	var books []*entities.Book
	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
