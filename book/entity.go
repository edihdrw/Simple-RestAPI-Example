package book

import (
	"time"
)

type Book struct {
	ID          int
	Title       string
	Price       int
	Rating      int
	Discount    int
	Description string
	CreateAt    time.Time
	UpdateAt    time.Time
}
