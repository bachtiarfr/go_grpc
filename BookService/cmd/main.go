package main

import (
	"BookService/config"
	"BookService/internal/domain"
	"BookService/internal/repository"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Build PostgreSQL connection string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	// Connect to PostgreSQL
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	// Verify connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Error verifying database connection: %v", err)
	}

	// Initialize repository
	bookRepo := repository.NewBookRepositoryPG(db)

	// Example: Create a book
	book := &domain.Book{
		ID:         "uuid-v4-generated-id",
		Title:      "Go Programming",
		AuthorID:   "author-id",
		CategoryID: "category-id",
		Stock:      10,
	}
	err = bookRepo.Create(book)
	if err != nil {
		log.Fatalf("Error creating book: %v", err)
	}
	log.Println("Book created successfully!")
}
