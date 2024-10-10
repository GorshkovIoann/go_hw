package storage

import (
	"fmt"
	"task1/book"
)

type MapStorage struct {
	books map[int]*book.Book
}

func NewMapStorage() *MapStorage {
	return &MapStorage{books: make(map[int]*book.Book)}
}

func (s *MapStorage) Add(book *book.Book) {
	s.books[book.ID] = book
}

func (s *MapStorage) GetById(id int) (*book.Book, error) {
	book, ok := s.books[id]
	if !ok {
		return nil, fmt.Errorf("book not found")
	}

	return book, nil
}

func (s *MapStorage) SearchForBookByTitle(title string) (*book.Book, error) {
	for _, book := range s.books {
		if book.Title == title {
			return book, nil
		}
	}
	return nil, fmt.Errorf("book not found")
}
