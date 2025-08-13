package parser

import (
	"fmt"
	"io"
	parserlib "kidstales/internal/parser-lib"
	"regexp"
	"strconv"
	"strings"
)

type BookPageParser struct{}

func (p *BookPageParser) Parse(r io.ReadCloser) (map[string]any, error) {
	root, err := parserlib.ParseHTML(r)
	if err != nil {
		return nil, fmt.Errorf("parse html failed: %w", err)
	}

	iframe, err := root.Query("//iframe")
	if err != nil {
		return nil, fmt.Errorf("iframe error: %w", err)
	}

	firstIframe := iframe.First()
	if firstIframe == nil {
		return nil, fmt.Errorf("no iframe found")
	}

	iframeSrc, found := firstIframe.Attr("src")
	if !found {
		return nil, fmt.Errorf("iframe src not found")
	}

	imageBase := "/proxy" + trimHost(strings.Replace(iframeSrc, "mobile.html", "files/mobile", -1))

	pageCount := 0

	readLink, err := root.Query("//a[@target][@title]")
	if err != nil {
		return nil, fmt.Errorf("read link find error: %w", err)
	}

	if readLink.First() != nil {
		textNode := readLink.First().NextSibling()

		if textNode != nil {
			rx := regexp.MustCompile(`^\((\d+) стр\.\)$`)

			values := rx.FindStringSubmatch(strings.TrimSpace(textNode.Value()))

			if len(values) == 2 {
				pageCount, err = strconv.Atoi(values[1])
				if err != nil {
					return nil, fmt.Errorf("page count value %s parse failed: %w", textNode.Value(), err)
				}

				return map[string]any{
					"ImageBase": imageBase,
					"PageCount": pageCount,
				}, nil
			}
		}
	}

	return nil, fmt.Errorf("page count not found")
}
