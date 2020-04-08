package repo

import "github.com/guitarpawat/bookshelf/dto"

// BooksRepo is an interface to retrieve and save the book data to the database.
type BooksRepo interface {
	// GetById gets the book data from database and transform it to DTO.
	GetById(id string) (dto.Book, error)

	// Save saves book data to the database the ID and AddTime is ignored and generate by repository.
	Save(book dto.Book) error
}
