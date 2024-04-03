package usecase

import (
	"go/api/internal/domain"
	entities "go/api/internal/entities"
)

type BookUseCaseImpl struct {
	bookRepo domain.BookRepository
}

func NewBookUseCase(bookRepo domain.BookRepository) domain.BookUseCase {
	return &BookUseCaseImpl{
		bookRepo: bookRepo,
	}
}

func (uc *BookUseCaseImpl) CreateBook(book *entities.Book) error {
	return uc.bookRepo.Create(book)
}

func (uc *BookUseCaseImpl) UpdateBook(id uint, book *entities.Book) error {
	// Check if the book exists
	existingBook, err := uc.bookRepo.FindByID(id)
	if err != nil {
		return err
	}
	// Update the existing book fields
	existingBook.Title = book.Title
	existingBook.Author = book.Author
	existingBook.Pages = book.Pages
	existingBook.Published = book.Published
	return uc.bookRepo.Update(existingBook)
}

func (uc *BookUseCaseImpl) DeleteBook(id uint) error {
	return uc.bookRepo.Delete(id)
}

func (uc *BookUseCaseImpl) GetBookByID(id uint) (*entities.Book, error) {
	return uc.bookRepo.FindByID(id)
}

func (uc *BookUseCaseImpl) GetAllBooks() ([]*entities.Book, error) {
	return uc.bookRepo.FindAll()
}
