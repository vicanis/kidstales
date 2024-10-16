//go:build integration

package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func GetTestDB(t *testing.T) DB {
	tempFile, err := os.CreateTemp("", "test_db_*.sqlite")
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, tempFile.Close())
		require.NoError(t, os.Remove(tempFile.Name()))
	})

	return NewSqliteDB(tempFile.Name())
}
