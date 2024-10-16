//go:build integration

package db_test

import (
	"fmt"
	"kidstales/internal/db"
	"kidstales/internal/model"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetDB(t *testing.T) {
	db := db.GetTestDB(t)
	require.NotNil(t, db)
}

func TestEmptyDB(t *testing.T) {
	db := db.GetTestDB(t)

	bookList, err := db.GetBookList(0, 1)
	require.NoError(t, err)
	require.Zero(t, len(bookList.Books))
	require.False(t, bookList.HasNext)
}

func TestCreateAndGetOne(t *testing.T) {
	db := db.GetTestDB(t)

	const bookName = "test-book-name-1"

	testBook := &model.Book{
		Name:   bookName,
		Author: "test-book-author-1",
	}

	actualBook, err := db.GetBook(bookName)
	require.Error(t, err)
	require.ErrorIs(t, err, model.ErrNotFound)

	err = db.Add(testBook)
	require.NoError(t, err)

	actualBook, err = db.GetBook(bookName)
	require.NoError(t, err)
	require.Equal(t, &testBook, actualBook)
}

func TestCreateAndGetListBooks(t *testing.T) {
	db := db.GetTestDB(t)

	const limit = 10

	for i := 0; i < limit*2; i++ {
		testBook := &model.Book{
			Name:   fmt.Sprintf("test-name-%d", i),
			Author: fmt.Sprintf("author-%d", i/2),
		}

		err := db.Add(testBook)
		require.NoError(t, err)
	}

	for page := 0; page < 3; page++ {
		bookList, err := db.GetBookList(page, limit)
		require.NoError(t, err)

		switch page {
		case 0, 1:
			require.Equal(t, limit, len(bookList.Books))

		case 2:
			require.Zero(t, len(bookList.Books))
		}

		switch page {
		case 0:
			require.True(t, bookList.HasNext)

		case 1, 2:
			require.False(t, bookList.HasNext)
		}
	}
}
