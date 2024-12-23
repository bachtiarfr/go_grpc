package usecase

import (
	"BookService/internal/domain"
	"errors"
)

type BookUsecase struct {
	repository domain.BookRepository
}

func NewBookUsecase(repo domain.BookRepository) *BookUsecase {
	return &BookUsecase{repository: repo}
}

func (uc *BookUsecase) BorrowBook(id string) error {
	book, err := uc.repository.GetByID(id)
	if err != nil {
		return err
	}

	if book.Stock <= 0 {
		return errors.New("book out of stock")
	}

	book.Stock -= 1
	return uc.repository.Update(book)
}
