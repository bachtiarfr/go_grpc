package usecase

import "BookService/internal/domain"

type BookUsecase interface {
	GetBookByID(id string) (*domain.Book, error)
	CreateBook(book *domain.Book) error
	UpdateBook(book *domain.Book) error
	DeleteBook(id string) error
}
