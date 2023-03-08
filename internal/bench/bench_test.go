package bench

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkInts(b *testing.B) {
	b.Run("int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = i
		}
	})

	b.Run("int atomic", func(b *testing.B) {
		i := atomic.Int64{}

		for {
			li := i.Load()
			if li >= int64(b.N) {
				break
			}
			i.Add(1)
		}
	})

	b.Run("int with locks", func(b *testing.B) {
		mu := sync.Mutex{}
		i := 0

		for {
			mu.Lock()
			if i >= b.N {
				break
			}
			i++
			mu.Unlock()
		}
	})
}
