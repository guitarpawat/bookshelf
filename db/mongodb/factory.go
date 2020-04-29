package mongodb

import (
	"github.com/guitarpawat/bookshelf/db"
	"github.com/guitarpawat/bookshelf/db/repo"
)

type Factory struct {
}

func (f *Factory) GetBooksRepo() repo.BooksRepo {
	if booksRepo == nil {
		booksRepo = newBooksRepo(DefaultBooksCollectionName)
	}
	return booksRepo
}

var instance *Factory

func getInstance() db.Factory {
	if instance == nil {
		instance = &Factory{}
	}
	return instance
}
