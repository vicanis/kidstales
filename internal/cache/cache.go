package cache

import (
	"io"
	"log"
	"os"
	"path"
)

type Cache struct {
	cacheDir   string
	keyEncoder func(v any) string
}

func NewHttpRequestCache(dir string) *Cache {
	return &Cache{
		cacheDir:   dir,
		keyEncoder: keyFromHttpRequest,
	}
}

func (c *Cache) Get(key any) (data []byte, found bool) {
	f, err := os.Open(c.getFileName(c.keyEncoder(key)))
	if err != nil {
		if os.IsNotExist(err) {
			return
		}

		log.Fatalf("cache file open failed: %v", err)
	}

	defer f.Close()

	data, err = io.ReadAll(f)
	if err != nil {
		log.Fatalf("cache file read failed: %v", err)
	}

	found = true

	return
}

func (c *Cache) Put(key any, data []byte) {
	err := os.WriteFile(c.getFileName(c.keyEncoder(key)), data, 0644)
	if err != nil {
		log.Fatalf("cache key write failed: %v", err)
	}
}

func (c *Cache) getFileName(key string) string {
	return path.Join(c.cacheDir, "cache_"+key)
}
