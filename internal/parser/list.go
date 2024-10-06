package parser

import (
	"fmt"
	"io"
	"kidstales/internal/config"
	"kidstales/internal/model"
	parserlib "kidstales/internal/parser-lib"
	"regexp"
	"strings"
)

type BooksListPageParser struct{}

func (p *BooksListPageParser) Parse(r io.ReadCloser) (map[string]any, error) {
	defer r.Close()

	root, err := parserlib.ParseHTML(r)
	if err != nil {
		return nil, fmt.Errorf("parse html failed: %w", err)
	}

	linkNodes, err := root.Query(
		"//figure[@class]",
		func(node *parserlib.Node) (*parserlib.Node, bool) {
			value, found := node.Attr("class")
			return node, found && value == "post-thumbnail"
		},
		func(node *parserlib.Node) (*parserlib.Node, bool) {
			return node.FirstChild(), true
		},
	)
	if err != nil {
		return nil, fmt.Errorf("link node query failed: %w", err)
	}

	books := make([]*model.Book, 0, len(linkNodes.Nodes()))

	for _, node := range linkNodes.Nodes() {
		book := &model.Book{}

		if pageURL, found := node.Attr("href"); found {
			book.PageURL = "/book" + trimHost(pageURL)
		} else {
			continue
		}

		pictureNodes, err := node.Query("img")
		if err != nil {
			return nil, fmt.Errorf("book picture query failed: %w", err)
		}

		if pictureNode := pictureNodes.First(); pictureNode != nil {
			if srcset, found := pictureNode.Attr("srcset"); found {
				bookPictureURL, err := parserlib.GetLargestSrc(srcset)
				if err != nil {
					return nil, err
				}

				book.PictureURL = "/proxy" + strings.TrimPrefix(bookPictureURL, config.Host)
			}

			if bookNameAuthor, found := pictureNode.Attr("alt"); found {
				book.Name, book.Author, err = getBookNameAuthor(bookNameAuthor)
				if err != nil {
					return nil, fmt.Errorf("book name & author (%s) parse failed: %w", bookNameAuthor, err)
				}
			}
		}

		books = append(books, book)
	}

	return map[string]any{
		"Books": books,
	}, nil
}

func getBookNameAuthor(line string) (name, author string, err error) {
	rx := regexp.MustCompile(`[«"](.*)[»"] (.*)$`)

	values := rx.FindStringSubmatch(line)

	if len(values) < 3 {
		return "", "", model.ErrNotFound
	}

	return values[1], values[2], nil
}

func trimHost(s string) string {
	return strings.TrimPrefix(s, config.Host)
}
