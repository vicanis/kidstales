package parser

import (
	"fmt"
	"io"
	parserlib "kidstales/internal/parser-lib"
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

	iframeSrc, found := iframe.First().Attr("src")
	if !found {
		return nil, fmt.Errorf("iframe src not found")
	}

	imageBase := "/proxy" + trimHost(strings.Replace(iframeSrc, "mobile.html", "files/mobile", -1))

	return map[string]any{
		"ImageBase": imageBase,
	}, nil
}
