// Package minmaxer will keep track of a collection of (unsigned) integers, and return the minimum, maximum, and total
// sum of all collected values on demand. The collection is thread safe.
package minmaxer

import (
	"sync"
)

type Collection struct {
	values           []uint
	minimum, maximum uint
	sum              uint64
	sync.Mutex
}

func New() *Collection {
	c := &Collection{
		values:  make([]uint, 0),
		minimum: ^uint(0),
	}

	return c
}

func (c *Collection) Add(v uint) {
	c.Lock()
	defer c.Unlock()

	c.values = append(c.values, v)

	if v < c.minimum {
		c.minimum = v
	}

	if v > c.maximum {
		c.maximum = v
	}

	c.sum += uint64(v)
}

func (c *Collection) Minimum() uint {
	c.Lock()
	defer c.Unlock()

	return c.minimum
}

func (c *Collection) Maximum() uint {
	c.Lock()
	defer c.Unlock()

	return c.maximum
}

func (c *Collection) Count() int {
	c.Lock()
	defer c.Unlock()

	return len(c.values)
}

func (c *Collection) Sum() uint64 {
	c.Lock()
	defer c.Unlock()

	return c.sum
}
