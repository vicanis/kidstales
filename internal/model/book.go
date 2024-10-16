package model

type Book struct {
	Name       string
	Author     string
	PageURL    string
	PictureURL string
}

type BookList struct {
	Books   []*Book
	HasNext bool
}
