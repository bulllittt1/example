package stack_test

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/strelnykov/examples/stack"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	a := assert.New(t)

	var (
		value1 = 777
		value2 = "value2"
	)

	s := stack.NewStack()

	s.Push(value1)
	s.Push(value2)

	res1 := s.Pop()
	a.Equal(value2, res1)

	res2 := s.Pop()
	a.Equal(value1, res2)

	res3 := s.Pop()
	a.Nil(res3)
}

// go test -race
func TestStackConcurrent(t *testing.T) {
	tests := 10
	s := stack.NewStack()

	var wg sync.WaitGroup
	for i := 0; i < tests; i++ {
		value := i
		wg.Add(1)
		go func() {
			s.Push(value)
			wg.Done()
		}()
	}

	wg.Wait()

	for i := 0; i < tests; i++ {
		wg.Add(1)
		go func() {
			_ = s.Pop()
			wg.Done()
		}()
	}
	wg.Wait()
}

var result interface{}

func BenchmarkStack(b *testing.B) {
	s := stack.NewStack()
	tests := make([]int64, 1000)
	for i := range tests {
		tests[i] = rand.Int63()
	}

	b.ResetTimer()
	var r interface{}
	for i := 0; i < b.N; i++ {

		for _, value := range tests {
			s.Push(value)
		}

		for range tests {
			r = s.Pop()
		}
	}
	result = r

}
