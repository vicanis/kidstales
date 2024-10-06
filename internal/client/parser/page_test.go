package parser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBookPage(t *testing.T) {
	f, err := os.Open("testdata/rusalochka.html")
	require.NoError(t, err)
	defer f.Close()

	parsed, err := new(BookPageParser).Parse(f)
	require.NoError(t, err)

	require.Equal(t, map[string]any{
		"ImageBase": "/proxy/knigi/rusalochka/files/mobile",
	}, parsed)
}
