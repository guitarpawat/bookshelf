package repo

import (
	"github.com/guitarpawat/bookshelf/dto"
)

// BooksRepo is an interface to retrieve and save the book data to the database.
type BooksRepo interface {
	// GetById gets the book data from database and transform it to DTO.
	GetById(id string) (dto.Book, error)

	// GetPaginationSortByTimeDesc gets the books sorted by latest to oldest with paging.
	// It accepts limit for paging limit and lastId as the oldest item from the last page.
	// It returns books, id of the oldest item in this page and error if has.
	GetPaginationSortByTimeDesc(limit int, lastId string) ([]dto.Book, string, error)

	// Save saves book data to the database the ID and AddTime is ignored and generate by repository.
	Save(book dto.Book) error
}
