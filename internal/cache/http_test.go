package cache

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHttpRequestCacheKeyEncoder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		req := &http.Request{
			RequestURI: "/proxy/knigi/rusalochka/files/mobile/1.jpg",
			URL:        &url.URL{},
		}

		require.Equal(t, "c2204760a1576633d75b7f56e27845cd", keyFromHttpRequest(req))
	})

	t.Run("query", func(t *testing.T) {
		req1 := &http.Request{
			RequestURI: "/proxy/knigi/detskie",
			URL: &url.URL{
				RawQuery: "",
			},
		}
		key1 := keyFromHttpRequest(req1)
		require.NotEmpty(t, key1)

		req2 := &http.Request{
			RequestURI: "/proxy/knigi/detskie",
			URL: &url.URL{
				RawQuery: "page=1",
			},
		}
		key2 := keyFromHttpRequest(req2)
		require.NotEmpty(t, key2)

		require.NotEqual(t, key1, key2)
	})
}
