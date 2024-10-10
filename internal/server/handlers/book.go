package handlers

import (
	"fmt"
	"kidstales/internal/client"
	"kidstales/internal/parser"
	"net/http"
	"strconv"
	"strings"
)

func Book(r *http.Request) (map[string]any, error) {
	pageIndex := 0

	pageNumberStr := strings.TrimSpace(r.URL.Query().Get("page"))
	if len(pageNumberStr) > 0 {
		var err error
		pageIndex, err = strconv.Atoi(pageNumberStr)
		if err != nil {
			return nil, fmt.Errorf("page number %s parse failed: %w", pageNumberStr, err)
		}
	}

	bookPath := strings.TrimPrefix(r.URL.Path, "/book")

	client := client.New()

	pageReader, err := client.GetWithCache(bookPath)
	if err != nil {
		return nil, err
	}

	data, err := new(parser.BookPageParser).Parse(pageReader)
	if err != nil {
		return nil, err
	}

	if pageIndex > 0 {
		data["PreviousPageURL"] = fmt.Sprintf("%s?page=%d", r.URL.Path, pageIndex-1)
	}

	pageCount := data["PageCount"].(int)

	if pageCount > 0 && pageIndex < pageCount-1 {
		data["NextPageURL"] = fmt.Sprintf("%s?page=%d", r.URL.Path, pageIndex+1)
	}

	data["CurrentPageNumber"] = pageIndex + 1

	return data, nil
}
