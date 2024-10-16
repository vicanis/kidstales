package db

import (
	"kidstales/internal/config"
	"kidstales/internal/model"
)

type DB interface {
	GetBookList(page, limit int) (*model.BookList, error)
	GetBook(name string) (*model.Book, error)
	Add(book model.Book) error
}

func GetDefaultDB() DB {
	return NewSqliteDB(config.DatabasePath)
}
