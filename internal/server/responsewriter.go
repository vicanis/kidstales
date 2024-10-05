package server

import (
	"net/http"
	"strings"
)

type textResponseHandlerFunc func(r *http.Request) (map[string]any, error)

func WithHtmlResponse(templateName string, handler textResponseHandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/html; charset=utf-8")

		bodyWriter := &strings.Builder{}

		defer func() {
			mustRenderTemplate("base.go.tmpl", w, bodyWriter.String(), true)
		}()

		data, err := handler(r)
		if err != nil {
			templateName = "error.go.tmpl"
			data = map[string]any{
				"ErrorMessage": err.Error(),
			}
		}

		mustRenderTemplate(templateName, bodyWriter, data, false)
	})
}
