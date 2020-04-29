package dto

import "errors"

type BookType int

var bookTypeString = []string{"Hard cover", "Soft cover", "PDF", "EBook Provider", "Others"}

const (
	HardCover BookType = iota
	SoftCover
	PDF
	EBookProvider
	Others
)

func (b BookType) Int() int {
	return int(b)
}

func (b BookType) String() string {
	return bookTypeString[b]
}

func GetBookTypes() []BookType {
	return []BookType{HardCover, SoftCover, PDF, EBookProvider, Others}
}

func ToBookType(status string) (BookType, error) {
	for k, v := range bookTypeString {
		if v == status {
			return GetBookTypes()[k], nil
		}
	}

	return 0, errors.New("cannot find matching book status")
}
