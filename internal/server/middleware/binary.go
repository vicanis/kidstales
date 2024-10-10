package middleware

import (
	"kidstales/internal/server/render"
	"net/http"
)

type binaryHandlerFunc func(r *http.Request) ([]byte, error)

func WithBinaryResponse(contentType string, handler binaryHandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := handler(r)
		if err != nil {
			render.RenderError(err, w)
			return
		}

		w.Header().Set("Content-Type", contentType)
		w.Write(data)
	})
}
