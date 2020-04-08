package repo

import "github.com/guitarpawat/bookshelf/dto"

type BooksRepo interface {
	// GetById gets the book data from database
	// and transform it to DTO
	GetById(id string) (dto.Book, error)

	// Save saves book data to the database
	// the ID field is ignored and generate by repository
	Save(book dto.Book) error
}
