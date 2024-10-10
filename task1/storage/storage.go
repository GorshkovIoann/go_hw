package storage

import (
	"task1/book"
)

type Storage interface {
	Add(book *book.Book)
	GetById(id int) (*book.Book, error)
	SearchForBookByTitle(title string) (*book.Book, error)
}
