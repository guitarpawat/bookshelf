package db

import (
	"errors"
	"github.com/guitarpawat/bookshelf/db/mongodb"
	"github.com/guitarpawat/bookshelf/db/repo"
)

// Factory provides an interface to provide dependency inversion of repository
// for using with other package in the program.
type Factory interface {
	GetBooksRepo() repo.BooksRepo
}

// GetRepoInstance returns database vendor repository that specified as dbName in arguments of the function.
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
