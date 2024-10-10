package middleware

import (
	"kidstales/internal/cache"
	"kidstales/internal/config"
	"kidstales/internal/server/render"
	"net/http"
)

func WithCache(handler binaryHandlerFunc) http.Handler {
	httpCache := cache.NewHttpRequestCache(config.CacheDir)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if data, found := httpCache.Get(r); found {
			w.Header().Set(cachedHeaderName, "true")
			w.Write(data)
			return
		}

		data, err := handler(r)
		if err != nil {
			render.RenderError(err, w)
			return
		}

		httpCache.Put(r, data)

		w.Write(data)
	})
}
