package storage

import (
	"fmt"
	"task1/book"
)

type SliceStorage struct {
	books []*book.Book
}

func NewSliceStorage() *SliceStorage {
	return &SliceStorage{books: make([]*book.Book, 0)}
}

func (s *SliceStorage) Add(book *book.Book) {
	s.books = append(s.books, book)
}

func (s *SliceStorage) GetById(id int) (*book.Book, error) {
	for _, book := range s.books {
		if book.ID == id {
			return book, nil
		}
	}
	return nil, fmt.Errorf("book not found")
}

func (s *SliceStorage) SearchForBookByTitle(title string) (*book.Book, error) {
	for _, book := range s.books {
		if book.Title == title {
			return book, nil
		}
	}
	return nil, fmt.Errorf("book not found")
}
