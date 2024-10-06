package cache

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache_Flow(t *testing.T) {
	cacheDir := os.TempDir()

	testCache := &Cache{
		cacheDir: cacheDir,
		keyEncoder: func(v any) string {
			return v.(string)
		},
	}

	testData := []byte{0x01, 0x02, 0x03}

	testKey := "test"
	testFileName := path.Join(cacheDir, "cache_"+testKey)

	_, found := testCache.Get(testKey)
	require.False(t, found)

	testCache.Put(testKey, testData)

	t.Logf("test file name %s", testFileName)
	require.FileExists(t, testFileName)

	t.Cleanup(func() {
		os.Remove(testFileName)
	})

	data, found := testCache.Get(testKey)
	require.True(t, found)
	require.Equal(t, testData, data)

}
