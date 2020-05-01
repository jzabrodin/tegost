package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func Benchmark(b *testing.B) {
	var value uint64
	var wg sync.WaitGroup

	wg.Add(b.N)

	for i := 0; i < b.N; i++ {
		go func() {
			atomic.AddUint64(&value, 1)
			wg.Done()
		}()
	}

	wg.Wait()

	if value != uint64(b.N) {
		b.Errorf("expected %d, got %d", b.N, value)
	}
}
