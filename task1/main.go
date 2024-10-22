package main

import (
	"fmt"
	"idGenerators"

	"task1/book"
	"task1/library"
	"task1/storage"
)

func main() {
	books := []*book.Book{
		&book.Book{Title: "Три мушкетера", Author: "Дюма"},
		&book.Book{Title: "Библия", Author: "Господь"},
		&book.Book{Title: "Вечера на хуторе близ Диканьки", Author: "Николай Гоголь"},
	}

	sliceStorage := storage.NewSliceStorage()
	library := library.NewLibrary(sliceStorage, idgenerators.GenerateRandomID)

	for _, book := range books {
		library.AddBook(book.Title, book.Author)
	}

	foundBook, err := library.FindBookByName("Библия")
	if err != nil {
		fmt.Println("Книга не найдена")
	} else {
		fmt.Printf("Найдена книга: %+v\n", foundBook)
	}

	library.IdGen = idgenerators.GenerateUUID

	library.AddBook("Пособие по выживанию на стипу", "Неизвестный студент")

	foundBook, err = library.FindBookByName("Три мушкетера")
	if err != nil {
		fmt.Println("Книга не найдена")
	} else {
		fmt.Printf("Найден книга: %+v\n", foundBook)
	}

	mapStorage := storage.NewMapStorage()
	library.Storage = mapStorage

	for _, book := range books {
		library.AddBook(book.Title, book.Author)
	}

	for _, title := range []string{"Библия", "Вечера на хуторе близ Диканьки"} {
		foundBook, err = library.FindBookByName(title)
		if err != nil {
			fmt.Println("Не нашел книгу")
		} else {
			fmt.Printf("Найдена книга: %+v\n", foundBook)
		}
	}
}
