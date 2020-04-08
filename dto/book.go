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
	Volume  []int
	Owner   string
	AddTime time.Time
}

type BookStatus int

var bookStatusString = []string{"Not owned", "Wanted", "Owned", "Read", "Lost"}

func (b BookStatus) Int() int {
	return int(b)
}

func (b BookStatus) String() string {
	return bookStatusString[b]
}

const (
	NotOwned BookStatus = iota
	Wanted
	Owned
	Read
	Lost
)

type BookType int

var bookTypeString = []string{"Hard cover", "Soft cover", "PDF", "EBook Provider"}

const (
	HardCover BookType = iota
	SoftCover
	PDF
	EBookProvider
)

func (b BookType) Int() int {
	return int(b)
}

func (b BookType) String() string {
	return bookTypeString[b]
}
