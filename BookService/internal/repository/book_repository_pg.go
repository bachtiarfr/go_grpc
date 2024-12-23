package repository

import (
	"BookService/internal/domain"
	"database/sql"
	"fmt"
)

type BookRepositoryPG struct {
	db *sql.DB
}

func NewBookRepositoryPG(db *sql.DB) *BookRepositoryPG {
	return &BookRepositoryPG{db: db}
}

// GetByID fetches a book by its ID
func (r *BookRepositoryPG) GetByID(id string) (*domain.Book, error) {
	var book domain.Book
	query := "SELECT id, title, author_id, category_id, stock FROM books WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.AuthorID, &book.CategoryID, &book.Stock)
	if err != nil {
		return nil, fmt.Errorf("error fetching book by ID: %w", err)
	}
	return &book, nil
}

// Create adds a new book to the database
func (r *BookRepositoryPG) Create(book *domain.Book) error {
	query := "INSERT INTO books (id, title, author_id, category_id, stock, created_at) VALUES ($1, $2, $3, $4, $5, NOW())"
	_, err := r.db.Exec(query, book.ID, book.Title, book.AuthorID, book.CategoryID, book.Stock)
	if err != nil {
		return fmt.Errorf("error creating book: %w", err)
	}
	return nil
}

// Update modifies an existing book's data
func (r *BookRepositoryPG) Update(book *domain.Book) error {
	query := "UPDATE books SET title = $1, stock = $2, updated_at = NOW() WHERE id = $3"
	_, err := r.db.Exec(query, book.Title, book.Stock, book.ID)
	if err != nil {
		return fmt.Errorf("error updating book: %w", err)
	}
	return nil
}

// Delete removes a book from the database
func (r *BookRepositoryPG) Delete(id string) error {
	query := "DELETE FROM books WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting book: %w", err)
	}
	return nil
}
