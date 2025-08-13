package cache

import (
	"kidstales/internal/cache/sqlite"
	"log"
)

type Cache struct {
	keyEncoder func(v any) string
}

func NewHttpRequestCache() *Cache {
	return &Cache{
		keyEncoder: keyFromHttpRequest,
	}
}

func (c *Cache) Get(key any) (data []byte, found bool) {
	var err error
	data, err = sqlite.NewDBCache().Get(c.keyEncoder(key))
	if err == nil {
		found = true
	} else {
		log.Print(err)
	}

	return
}

func (c *Cache) Put(key any, data []byte) {
	err := sqlite.NewDBCache().Set(c.keyEncoder(key), data)
	if err != nil {
		log.Print(err)
	}
}
