package handlers

import (
	"kidstales/internal/client"
	"kidstales/internal/parser"
	"net/http"
)

func BookList(r *http.Request) (map[string]any, error) {
	client := client.New()

	pageReader, err := client.GetWithCache("/")
	if err != nil {
		return nil, err
	}

	return new(parser.BooksListPageParser).Parse(pageReader)
}
