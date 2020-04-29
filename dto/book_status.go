package dto

import "errors"

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

func GetBookStatuses() []BookStatus {
	return []BookStatus{NotOwned, Wanted, Owned, Read, Lost}
}

func ToBookStatus(status string) (BookStatus, error) {
	for k, v := range bookStatusString {
		if v == status {
			return GetBookStatuses()[k], nil
		}
	}

	return 0, errors.New("cannot find matching book status")
}
