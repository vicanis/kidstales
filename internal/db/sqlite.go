package db

import (
	"io/fs"
	"kidstales/internal/model"
	"log"
	"os"

	pkgerr "github.com/pkg/errors"
	lib "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteDB struct {
	db *gorm.DB
}

func NewSqliteDB(path string) *SqliteDB {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		f, err := os.OpenFile(path, 0644, fs.FileMode(os.O_CREATE))
		if err != nil {
			log.Fatalf("NewSqliteDB: cannot create database file: %v", err)
		}

		_ = f.Close()
	}

	db, err := gorm.Open(lib.Open(path), &gorm.Config{})
	if err != nil {
		log.Fatalf("NewSqliteDB: database connect failed: %v", err)
	}

	db.AutoMigrate(&model.Book{})

	return &SqliteDB{db: db}
}

func (d *SqliteDB) GetBookList(page, limit int) (*model.BookList, error) {
	list := make([]model.Book, 0)

	result := d.db.Offset(page * limit).Limit(limit + 1).Find(&list)
	if result.Error != nil {
		return nil, pkgerr.Wrap(result.Error, "book list select failed")
	}

	if len(list) < limit {
		return &model.BookList{
			Books:   list,
			HasNext: false,
		}, nil
	}

	return &model.BookList{
		Books:   list[:limit],
		HasNext: len(list) > limit,
	}, nil
}

func (d *SqliteDB) GetBook(name string) (*model.Book, error) {
	book := &model.Book{}

	result := d.db.Find(&book, map[string]any{"name": name})
	if result.Error != nil {
		return nil, pkgerr.Wrap(result.Error, "get book failed")
	}

	return book, nil
}

func (d *SqliteDB) Add(book model.Book) error {
	result := d.db.Create(book)
	if result.Error != nil {
		return pkgerr.Wrap(result.Error, "book insert failed")
	}

	return nil
}
