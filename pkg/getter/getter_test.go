package getter_test

import (
	"testing"

	"github.com/jlucktay/go-getter/pkg/getter"
	"github.com/matryer/is"
)

func TestGetter(t *testing.T) {
	t.Parallel()
	testCases := map[string]struct {
		getURL   string
		expected uint
	}{
		"Forty two": {
			getURL:   "http://httpbin.org/bytes/42",
			expected: 42,
		},
		"Zero": {
			getURL:   "http://httpbin.org/bytes/0",
			expected: 0,
		},
		"Timeout": {
			getURL:   "http://httpbin.org/delay/10",
			expected: 0,
		},
	}
	for name, tC := range testCases {
		tC := tC
		t.Run(name, func(t *testing.T) {
			g := getter.New(5)
			actual := g.Get(tC.getURL)
			i := is.New(t)
			i.Equal(tC.expected, actual)
		})
	}
}
