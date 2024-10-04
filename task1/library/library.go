package library

import (
	"task1/book"
	"task1/storage"
)

type library interface {
	AddBook(title, author string)
	FindBookByName(name string) (*book.Book, error)
}

type Library struct {
	storage.Storage
	IdGen func() int
}

func NewLibrary(storage storage.Storage, idGenFunc func() int) *Library {
	return &Library{Storage: storage, IdGen: idGenFunc}
}

func (l *Library) AddBook(title, author string) {
	book := &book.Book{
		ID:     l.IdGen(),
		Title:  title,
		Author: author,
	}
	l.Storage.Add(book)
}

func (l *Library) FindBookByName(name string) (*book.Book, error) {
	return l.Storage.SearchForBookByTitle(name)
}
