package db

import (
	"github.com/guitarpawat/bookshelf/db/repo"
)

// Factory provides an interface to provide dependency inversion of repository
// for using with other package in the program.
type Factory interface {
	GetBooksRepo() repo.BooksRepo
}
