package dto

import "time"

type Book struct {
	ID      string
	Title   string
	Edition string
	Author  []string
	Tags    []string
	Type    BookType
	Status  BookStatus
	Volume  []string
	AddTime time.Time
}
