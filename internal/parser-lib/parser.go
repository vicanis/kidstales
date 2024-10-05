package parserlib

import (
	"errors"
	"fmt"
	"io"

	"github.com/antchfx/htmlquery"
)

type Parser interface {
	Parse(r io.ReadCloser) (map[string]any, error)
}

var ErrNotFound = errors.New("not found")

func ParseHTML(r io.Reader) (*Node, error) {
	root, err := htmlquery.Parse(r)
	if err != nil {
		return nil, fmt.Errorf("parse html failed: %w", err)
	}

	return &Node{node: root}, nil
}
