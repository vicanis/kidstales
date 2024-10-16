package handlers

import (
	"errors"
	"fmt"
	"kidstales/internal/client"
	"kidstales/internal/db"
	"kidstales/internal/model"
	"kidstales/internal/parser"
	"net/http"
	"strconv"
)

func BookList(r *http.Request) (map[string]any, error) {
	pageNumber := 1

	if pageNumberStr := r.URL.Query().Get("page"); pageNumberStr != "" {
		d, err := strconv.Atoi(pageNumberStr)
		if err != nil {
			return nil, fmt.Errorf("page number %s parse failed: %w", pageNumberStr, err)
		}

		if d < 1 {
			return nil, fmt.Errorf("page number should be >= 1")
		}

		pageNumber = d
	}

	client := client.New()

	path := "/"
	if pageNumber > 1 {
		path += fmt.Sprintf("page/%d", pageNumber)
	}

	pageReader, err := client.GetWithCache(path)
	if err != nil {
		return nil, err
	}

	response, err := new(parser.BooksListPageParser).Parse(pageReader)
	if err != nil {
		return nil, err
	}

	db := db.GetDefaultDB()

	books := response["Books"].([]*model.Book)
	for _, book := range books {
		_, err = db.GetBook(book.Name)
		if err != nil {
			if !errors.Is(err, model.ErrNotFound) {
				return nil, err
			}

			err = db.Add(book)
			if err != nil {
				return nil, err
			}
		}
	}

	if pageNumber > 1 {
		response["PreviousPage"] = pageNumber - 1
	}

	if response["HasNextPage"].(bool) {
		response["NextPage"] = pageNumber + 1
	}

	return response, nil
}
