package bench

import (
	"crypto/rand"
	"hash/maphash"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/cespare/xxhash"
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

func BenchmarkHashing(b *testing.B) {
	b.ReportAllocs()

	b.Run("native", func(b *testing.B) {
		set := map[string]int{
			"4B":   4,
			"16B":  16,
			"100B": 100,
			"4KB":  4096,
			"10MB": 10 * 1e6,
		}

		for k, val := range set {
			b.Run(k, func(b *testing.B) {
				b.SetBytes(int64(val))

				byts := make([]byte, val)
				if _, err := rand.Read(byts); err != nil {
					b.Fatal(err)
				}

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					var hasher maphash.Hash
					hasher.Write(byts)
					_ = hasher.Sum64()
				}
			})
		}
	})

	b.Run("xxhash", func(b *testing.B) {
		set := map[string]int{
			"4B":   4,
			"16B":  16,
			"100B": 100,
			"4KB":  4096,
			"10MB": 10 * 1e6,
		}

		for k, val := range set {
			b.Run(k, func(b *testing.B) {
				b.SetBytes(int64(val))

				byts := make([]byte, val)
				if _, err := rand.Read(byts); err != nil {
					b.Fatal(err)
				}

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					hasher := xxhash.New()
					hasher.Write(byts)
					_ = hasher.Sum64()
				}
			})
		}
	})
}
