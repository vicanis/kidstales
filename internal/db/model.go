package db

import (
	"kidstales/internal/model"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	model.Book
}

func toDBModel(book *model.Book) *Book {
	if book == nil {
		return nil
	}

	return &Book{Book: *book}
}

func toAppModel(book *Book) *model.Book {
	if book == nil {
		return nil
	}

	return &book.Book
}

func toAppModelBooks(books []*Book) []*model.Book {
	list := make([]*model.Book, len(books))
	for _, book := range books {
		list = append(list, toAppModel(book))
	}

	return list
}
