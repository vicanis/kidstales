package middleware

import (
	"kidstales/internal/server/render"
	"net/http"
)

type textResponseHandlerFunc func(r *http.Request) (map[string]any, error)

func WithHtmlResponse(templateName string, handler textResponseHandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := handler(r)
		if err != nil {
			render.RenderError(err, w)
		} else {
			render.RenderTemplate(templateName, data, w)
		}
	})
}
