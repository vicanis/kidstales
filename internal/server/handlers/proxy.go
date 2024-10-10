package handlers

import (
	"io"
	"kidstales/internal/client"
	"kidstales/internal/server/render"
	"net/http"
	"strings"
)

func Proxy(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/proxy")

	client := client.New()

	dataReader, err := client.GetWithCache(path)
	if err != nil {
		render.RenderError(err, w)
		return
	}

	defer dataReader.Close()

	data, err := io.ReadAll(dataReader)
	if err != nil {
		render.RenderError(err, w)
		return
	}

	w.Write(data)
}
