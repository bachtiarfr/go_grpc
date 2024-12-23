package usecase

import (
	"BookService/internal/domain"
	"BookService/internal/repository"
	"errors"
)

type bookUsecase struct {
	bookRepo repository.BookRepositoryPG
}

func NewBookUsecase(bookRepo *repository.BookRepositoryPG) BookUsecase {
	return &bookUsecase{
		bookRepo: bookRepo,
	}
}

func (u *bookUsecase) GetBookByID(id string) (*domain.Book, error) {
	book, err := u.bookRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("book not found")
	}
	return book, nil
}

func (u *bookUsecase) CreateBook(book *domain.Book) error {
	if book.Title == "" {
		return errors.New("book title cannot be empty")
	}
	return u.bookRepo.Create(book)
}

func (u *bookUsecase) UpdateBook(book *domain.Book) error {
	existingBook, err := u.bookRepo.GetByID(book.ID)
	if err != nil {
		return errors.New("book not found")
	}
	existingBook.Title = book.Title
	existingBook.Stock = book.Stock
	return u.bookRepo.Update(existingBook)
}

func (u *bookUsecase) DeleteBook(id string) error {
	_, err := u.bookRepo.GetByID(id)
	if err != nil {
		return errors.New("book not found")
	}
	return u.bookRepo.Delete(id)
}
