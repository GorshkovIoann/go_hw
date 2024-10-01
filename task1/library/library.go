package library

import (
	"task1/book"
	"task1/storage"
)

type library interface {
	findBookById(id int) (bool, book.Book)
	findBookByAuthor(author string) (bool, book.Book)
	putBook(title string, author string)
	libraryRelocation()
	setIdFunction()
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

func (l *Library) FindBookByName(name string) (*book.Book, bool) {
	return l.Storage.SearchForBookByTitle(name)
}
