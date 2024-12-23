package domain

type BookRepository interface {
	GetByID(id string) (*Book, error)
	Create(book *Book) error
	Update(book *Book) error
	Delete(id string) error
	List(filter map[string]interface{}) ([]Book, error)
}
