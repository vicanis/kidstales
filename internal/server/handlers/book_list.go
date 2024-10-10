package handlers

import (
	"fmt"
	"kidstales/internal/client"
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

	if pageNumber > 1 {
		response["PreviousPage"] = pageNumber - 1
	}

	response["NextPage"] = pageNumber + 1

	return response, nil
}
