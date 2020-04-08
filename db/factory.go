package db

import (
	"errors"
	"github.com/guitarpawat/bookshelf/db/mongodb"
	"github.com/guitarpawat/bookshelf/db/repo"
)

type Factory interface {
	GetBooksRepo() repo.BooksRepo
}

func GetRepoInstance(dbName DatabaseName) Factory {
	switch dbName {
	case MongoDB:
		return mongodb.GetInstance()
	default:
		panic(errors.New("database type is not supported"))
	}
}

type DatabaseName string

var MongoDB DatabaseName = "MongoDB"
