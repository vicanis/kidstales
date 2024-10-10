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

	testKey := "01234567890123456789012345678912"

	testDirName := path.Join(cacheDir, testKey[:prefixLength])
	testFileName := testKey[prefixLength:]

	t.Logf("test dir %s, file name %s", testDirName, testFileName)

	_, found := testCache.Get(testKey)
	require.False(t, found)

	testCache.Put(testKey, testData)

	require.DirExists(t, testDirName)
	require.FileExists(t, path.Join(testDirName, testFileName))

	t.Cleanup(func() {
		require.NoError(t, os.Remove(path.Join(testDirName, testFileName)))
		require.NoError(t, os.Remove(testDirName))
	})

	data, found := testCache.Get(testKey)
	require.True(t, found)
	require.Equal(t, testData, data)
}
