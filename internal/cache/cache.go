package cache

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

const prefixLength = 2

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
	keyString := c.keyEncoder(key)

	dirName, fileName := c.getPathComponents(keyString)

	exists, err := isDirExists(dirName)
	if err != nil {
		log.Fatalf("isDirExists failed: %v", err)
	}

	if !exists {
		return
	}

	f, err := os.Open(path.Join(dirName, fileName))
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
	keyString := c.keyEncoder(key)

	dir, file := c.getPathComponents(keyString)

	err := createDirIfNeeded(dir)
	if err != nil {
		log.Fatal("create temp dir failed: %w", err)
	}

	err = os.WriteFile(path.Join(dir, file), data, 0644)
	if err != nil {
		log.Fatalf("cache key write failed: %v", err)
	}
}

func (c *Cache) getPathComponents(key string) (dir, file string) {
	return path.Join(c.cacheDir, key[:prefixLength]), key[prefixLength:]
}

func createDirIfNeeded(dir string) error {
	exists, err := isDirExists(dir)
	if err != nil {
		return err
	}

	if !exists {
		log.Printf("create temp dir %s", dir)
		return os.Mkdir(dir, 0755)
	}

	return nil
}

func isDirExists(dir string) (bool, error) {
	info, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	if info != nil && !info.IsDir() {
		return false, fmt.Errorf("%s is not directory", dir)
	}

	return true, nil
}
