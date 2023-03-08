package queue

import "testing"

func Benchmark_Queue_Push(b *testing.B) {
	b.ReportAllocs()

	b.Run("single threaded - 1 item", func(b *testing.B) {
		q := New[int]()

		for i := 0; i < b.N; i++ {
			q.Push(i)
		}
	})

	b.Run("single threaded - 100 item", func(b *testing.B) {
		q := New[int]()

		for i := 0; i < b.N; i++ {
			for j := 0; j < 100; j++ {
				q.Push(i)
			}
		}
	})

	b.Run("single threaded - 1000 item", func(b *testing.B) {
		q := New[int]()

		for i := 0; i < b.N; i++ {
			for j := 0; j < 1000; j++ {
				q.Push(i)
			}
		}
	})

	b.Run("multi threaded - 1 per goroutine", func(b *testing.B) {
		q := New[int]()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				q.Push(1)
			}
		})
	})

	b.Run("multi threaded - 100 per goroutine", func(b *testing.B) {
		q := New[int]()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 0; i < 100; i++ {
					q.Push(1)
				}
			}
		})
	})

	b.Run("multi threaded - 1k per goroutine", func(b *testing.B) {
		q := New[int]()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 0; i < 1000; i++ {
					q.Push(1)
				}
			}
		})
	})
}

func Benchmark_Queue_Pop(b *testing.B) {
	b.ReportAllocs()

	b.Run("single threaded - 1 item", func(b *testing.B) {
		q := New[int]()

		for i := 0; i < b.N; i++ {
			q.Push(i)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			q.Pop()
		}
	})

	b.Run("single threaded - 100 item", func(b *testing.B) {
		q := New[int]()

		for i := 0; i < b.N; i++ {
			q.Push(i)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for j := 0; j < 100; j++ {
				q.Pop()
			}
		}
	})

	b.Run("single threaded - 1000 item", func(b *testing.B) {
		q := New[int]()

		for i := 0; i < b.N; i++ {
			q.Push(i)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for j := 0; j < 1000; j++ {
				q.Pop()
			}
		}
	})

	b.Run("multi threaded - 1 per goroutine", func(b *testing.B) {
		q := New[int]()

		for i := 0; i < b.N; i++ {
			q.Push(i)
		}

		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				q.Pop()
			}
		})
	})

	b.Run("multi threaded - 100 per goroutine", func(b *testing.B) {
		q := New[int]()

		for i := 0; i < b.N; i++ {
			q.Push(i)
		}

		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 0; i < 100; i++ {
					q.Pop()
				}
			}
		})
	})

	b.Run("multi threaded - 1k per goroutine", func(b *testing.B) {
		q := New[int]()

		for i := 0; i < b.N; i++ {
			q.Push(i)
		}

		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 0; i < 1000; i++ {
					q.Pop()
				}
			}
		})
	})
}
