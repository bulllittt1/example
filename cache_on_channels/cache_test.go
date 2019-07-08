package cache

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	const (
		key = iota
		keyF
	)
	a := assert.New(t)
	c, cM := newCache(), newCacheM()

	var val = rand.Int63()
	c.set(key, val)
	a.Equal(val, c.read(key))
	a.Nil(c.read(keyF))

	cM.set(key, val)
	a.Equal(val, cM.read(key))
	a.Nil(cM.read(keyF))
}

// Benchmark
// go test -bench=.
//
var load int

func BenchmarkLight(b *testing.B) {
	load = 10
	b.Run("Channels", func(b *testing.B) {
		testCache(newCache(), b)
	})
	b.Run("Mutex", func(b *testing.B) {
		testCache(newCacheM(), b)
	})
}

func BenchmarkHeavy(b *testing.B) {
	load = 1 << 20
	b.Run("Channels", func(b *testing.B) {
		testCache(newCache(), b)
	})
	b.Run("Mutex", func(b *testing.B) {
		testCache(newCacheM(), b)
	})
}

func testCache(c cacher, b *testing.B) {
	a := assert.New(b)

	keys := make([]int, load)
	vals := make([]int64, load)

	for i := 0; i < load; i++ {
		key := rand.Int()
		keys[i] = key

		val := rand.Int63()
		vals[i] = val
	}

	for n := 0; n < b.N; n++ {
		var wg sync.WaitGroup
		for i := 0; i < load; i++ {
			wg.Add(1)
			j := i
			go func() {
				c.set(keys[j], vals[j])
				wg.Done()
			}()
		}
		wg.Wait()

		for i := 0; i < load; i++ {
			wg.Add(1)
			j := i
			go func() {
				r := c.read(keys[j])
				a.Equal(vals[j], r)
				wg.Done()
			}()
		}
		wg.Wait()

		for i := 0; i < load/2; i++ {
			wg.Add(2)
			j := i
			go func() {
				c.set(keys[j], vals[j])
				wg.Done()
			}()
			go func() {
				j = load - j - 1 // inverse
				r := c.read(keys[j])
				a.Equal(vals[j], r)
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
