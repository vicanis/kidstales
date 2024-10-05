package parserlib

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetLargestSrc(t *testing.T) {
	const (
		srcBase = "https://skazkiwsem.fun/wp-content/uploads/2023/11/"

		src650 = srcBase + "0-0001-650x650.jpg"
		src144 = srcBase + "0-0001-144x144.jpg"
		src45  = srcBase + "0-0001-45x45.jpg"

		size650 = src650 + " 650w"
		size144 = src144 + " 144w"
		size45  = src45 + " 45w"
	)

	testCases := []struct {
		Name          string
		SrcSet        string
		ExpectedSrc   string
		ExpectedError bool
	}{
		{
			"empty",
			"",
			"",
			true,
		},
		{
			"single",
			size650,
			src650,
			false,
		},
		{
			"descending-2",
			joinSrcSet(size144, size45),
			src144,
			false,
		},
		{
			"ascending-2",
			joinSrcSet(size45, size144),
			src144,
			false,
		},
		{
			"descending-3",
			joinSrcSet(size650, size144, size45),
			src650,
			false,
		},
		{
			"ascending-3",
			joinSrcSet(size45, size144, size650),
			src650,
			false,
		},
		{
			"mixed-3",
			joinSrcSet(size45, size650, size144),
			src650,
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.Name, func(t *testing.T) {
			src, err := GetLargestSrc(tc.SrcSet)
			if tc.ExpectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.ExpectedSrc, src)
			}
		})
	}
}

func joinSrcSet(srcset ...string) string {
	return strings.Join(srcset, ", ")
}
