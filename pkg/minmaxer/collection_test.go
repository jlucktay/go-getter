package minmaxer_test

import (
	"sync"
	"testing"

	"github.com/jlucktay/go-getter/pkg/minmaxer"
	"github.com/matryer/is"
)

type testCase struct {
	inputs         []uint
	expMin, expMax uint
	expSum         uint64
	expCount       int
}

type testCaseMap map[string]testCase

func testCases() testCaseMap {
	return testCaseMap{
		"Hello world": {
			inputs:   []uint{1, 2, 3, 4, 5},
			expMin:   1,
			expMax:   5,
			expSum:   15,
			expCount: 5,
		},
		"dlrow olleH": {
			inputs:   []uint{5, 4, 3, 2, 1},
			expMin:   1,
			expMax:   5,
			expSum:   15,
			expCount: 5,
		},
		"Single input": {
			inputs:   []uint{42},
			expMin:   42,
			expMax:   42,
			expSum:   42,
			expCount: 1,
		},
		"Several identical inputs": {
			inputs:   []uint{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			expMin:   1,
			expMax:   1,
			expSum:   30,
			expCount: 30,
		},
		"No input at all": {
			inputs:   []uint{},
			expMin:   ^uint(0),
			expMax:   0,
			expSum:   0,
			expCount: 0,
		},
	}
}

func TestCollection(t *testing.T) {
	t.Parallel()
	for name, tC := range testCases() {
		tC := tC
		t.Run(name, func(t *testing.T) {
			c := minmaxer.New()
			for _, v := range tC.inputs {
				c.Add(v)
			}
			i := is.New(t)
			i.Equal(tC.expMin, c.Minimum())
			i.Equal(tC.expMax, c.Maximum())
			i.Equal(tC.expSum, c.Sum())
			i.Equal(tC.expCount, c.Count())
		})
	}
}

func TestCollectionParallel(t *testing.T) {
	t.Parallel()
	for name, tC := range testCases() {
		tC := tC
		t.Run(name, func(t *testing.T) {
			c := minmaxer.New()
			wg := &sync.WaitGroup{}
			for _, v := range tC.inputs {
				wg.Add(1)
				go func(v uint) {
					defer wg.Done()
					c.Add(v)
				}(v)
			}
			wg.Wait()
			i := is.New(t)
			i.Equal(tC.expMin, c.Minimum())
			i.Equal(tC.expMax, c.Maximum())
			i.Equal(tC.expSum, c.Sum())
			i.Equal(tC.expCount, c.Count())
		})
	}
}
