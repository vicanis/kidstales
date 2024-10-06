package cache

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func keyFromHttpRequest(v any) string {
	r := v.(*http.Request)

	key := struct {
		RequestURI string `json:"request_uri"`
		Query      string `json:"query"`
	}{
		RequestURI: r.RequestURI,
		Query:      r.URL.RawQuery,
	}

	jsonString, err := json.Marshal(key)
	if err != nil {
		log.Fatalf("failed to marshal cache key request %#v: %v", key, err)
	}

	h := md5.New()
	h.Write(jsonString)

	return fmt.Sprintf("%x", h.Sum(nil))
}
