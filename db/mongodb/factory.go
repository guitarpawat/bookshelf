package mongodb

import (
	"github.com/guitarpawat/bookshelf/db"
	"github.com/guitarpawat/bookshelf/db/repo"
)

type Factory struct {
	booksRepo repo.BooksRepo
}

func (f *Factory) GetBooksRepo() repo.BooksRepo {
	if f.booksRepo == nil {
		f.booksRepo = newBooksRepo(DefaultBooksCollectionName)
	}
	return f.booksRepo
}

var instance *Factory

func getInstance() db.Factory {
	if instance == nil {
		instance = &Factory{}
	}
	return instance
}
