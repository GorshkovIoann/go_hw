package storage

import (
	"task1/book"
)

type Storage interface {
	Add(book *book.Book)
	GetById(id int) (*book.Book, bool)
	SearchForBookByTitle(title string) (*book.Book, bool)
}

type SliceStorage struct {
	books []*book.Book
}

func NewSliceStorage() *SliceStorage {
	return &SliceStorage{books: make([]*book.Book, 0)}
}

func (s *SliceStorage) Add(book *book.Book) {
	s.books = append(s.books, book)
}

func (s *SliceStorage) GetById(id int) (*book.Book, bool) {
	for _, book := range s.books {
		if book.ID == id {
			return book, true
		}
	}
	return nil, false
}

func (s *SliceStorage) SearchForBookByTitle(title string) (*book.Book, bool) {
	for _, book := range s.books {
		if book.Title == title {
			return book, true
		}
	}
	return nil, false
}

type MapStorage struct {
	books map[int]*book.Book
}

func NewMapStorage() *MapStorage {
	return &MapStorage{books: make(map[int]*book.Book)}
}

func (s *MapStorage) Add(book *book.Book) {
	s.books[book.ID] = book
}

func (s *MapStorage) GetById(id int) (*book.Book, bool) {
	book, ok := s.books[id]
	if !ok {
		return nil, ok
	}

	return book, ok
}

func (s *MapStorage) SearchForBookByTitle(title string) (*book.Book, bool) {
	for _, book := range s.books {
		if book.Title == title {
			return book, true
		}
	}
	return nil, false
}
