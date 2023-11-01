package queue

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/tangledbytes/godc/internal/util"
)

func Test_Queue_Push(t *testing.T) {
	t.Run("single threaded", func(t *testing.T) {
		type test struct {
			name string
			data []int
			want []int
		}

		tests := []test{
			{
				name: "empty",
				data: []int{},
				want: []int{},
			},
			{
				name: "single element",
				data: []int{1},
				want: []int{1},
			},
			{
				name: "multiple elements",
				data: []int{1, 2, 3, 4, 5},
				want: []int{1, 2, 3, 4, 5},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				q := New[int]()

				for _, data := range tt.data {
					q.Push(data)
				}

				got := make([]int, 0, len(tt.want))
				for data := range iterQueue(q.head.Load().Next.Load()) {
					got = append(got, data)
				}

				if len(got) != len(tt.want) {
					t.Errorf("got %v, want %v", got, tt.want)
				}

				for i := range got {
					if got[i] != tt.want[i] {
						t.Errorf("got %v, want %v", got, tt.want)
					}
				}
			})
		}
	})

	t.Run("multi threaded", func(t *testing.T) {
		type test struct {
			name string
			data []int
			want []int
		}

		tests := []test{
			{
				name: "empty",
				data: []int{},
				want: []int{},
			},
			{
				name: "single element",
				data: []int{1},
				want: []int{1},
			},
			{
				name: "multiple elements",
				data: util.GenerateIntSeries(1, 1000),
				want: util.GenerateIntSeries(1, 1000),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				q := New[int]()

				var wg sync.WaitGroup

				for _, data := range tt.data {
					wg.Add(1)
					go func(data int) {
						q.Push(data)
						wg.Done()
					}(data)
				}

				wg.Wait()

				got := make([]int, 0, len(tt.want))
				for data := range iterQueue(q.head.Load().Next.Load()) {
					got = append(got, data)
				}

				if !util.CompareSliceUnordered(got, tt.want) {
					t.Errorf("got %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func Test_Queue_Pop(t *testing.T) {
	t.Run("single threaded", func(t *testing.T) {
		type testremovedata struct {
			want [][2]interface{}
		}

		type test struct {
			name   string
			data   []int
			remove testremovedata
			final  []int
		}

		tests := []test{
			{
				name: "empty - pop none",
				data: []int{},
				remove: testremovedata{
					want: [][2]interface{}{},
				},
				final: []int{},
			},
			{
				name: "empty - pop values",
				data: []int{},
				remove: testremovedata{
					want: [][2]interface{}{
						{0, false},
						{0, false},
					},
				},
				final: []int{},
			},
			{
				name: "single element",
				data: []int{1},
				remove: testremovedata{
					want: [][2]interface{}{
						{1, true},
						{0, false},
					},
				},
				final: []int{},
			},
			{
				name: "multiple elements",
				data: util.GenerateIntSeries(1, 1000),
				remove: testremovedata{
					want: [][2]interface{}{
						{1, true},
						{2, true},
					},
				},
				final: util.GenerateIntSeries(3, 1000),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				q := New[int]()

				for _, data := range tt.data {
					q.Push(data)
				}

				for _, data := range tt.remove.want {
					if got, ok := q.Pop(); got != data[0] || ok != data[1] {
						t.Errorf("got %v, want %v", [2]interface{}{got, ok}, data)
					}
				}

				got := make([]int, 0, len(tt.final))
				for data := range iterQueue(q.head.Load().Next.Load()) {
					got = append(got, data)
				}

				if len(got) != len(tt.final) {
					t.Errorf("got %v, want %v", got, tt.final)
				}

				for i := range got {
					if got[i] != tt.final[i] {
						t.Errorf("got %v, want %v", got, tt.final)
					}
				}
			})
		}
	})

	t.Run("multi threaded", func(t *testing.T) {
		type test struct {
			name   string
			data   []int
			remove int
			final  []int
		}

		tests := []test{
			{
				name:   "empty - pop none",
				data:   []int{},
				remove: 0,
				final:  []int{},
			},
			{
				name:   "empty - pop values",
				data:   []int{},
				remove: 2,
				final:  []int{},
			},
			{
				name:   "single element",
				data:   []int{1},
				remove: 2,
				final:  []int{},
			},
			{
				name:   "multiple elements",
				data:   util.GenerateIntSeries(1, 1000),
				remove: 200,
				final:  util.GenerateIntSeries(201, 1000),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				q := New[int]()

				for _, data := range tt.data {
					q.Push(data)
				}

				var wg sync.WaitGroup

				for i := 0; i < tt.remove; i++ {
					wg.Add(1)
					go func() {
						q.Pop()
						wg.Done()
					}()
				}

				wg.Wait()

				got := make([]int, 0, len(tt.final))
				for data := range iterQueue(q.head.Load().Next.Load()) {
					got = append(got, data)
				}

				if !util.CompareSliceUnordered(got, tt.final) {
					t.Errorf("got %v, want %v", got, tt.final)
				}
			})
		}
	})
}

func Test_Queue_Len(t *testing.T) {
	t.Run("single threaded", func(t *testing.T) {
		type test struct {
			name string
			data []int
			want int64
		}

		tests := []test{
			{
				name: "empty",
				data: []int{},
				want: 0,
			},
			{
				name: "single element",
				data: []int{1},
				want: 1,
			},
			{
				name: "multiple elements",
				data: util.GenerateIntSeries(1, 1000),
				want: 1000,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				q := New[int]()

				for _, data := range tt.data {
					q.Push(data)
				}

				if got := q.Len(); got != tt.want {
					t.Errorf("got %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func Test_Queue_PushPop(t *testing.T) {
	t.Run("multi threaded", func(t *testing.T) {
		type test struct {
			name   string
			data   []int
			remove int
		}

		tests := []test{
			{
				name:   "empty - pop none",
				data:   []int{},
				remove: 0,
			},
			{
				name:   "empty - pop values",
				data:   []int{},
				remove: 2,
			},
			{
				name:   "single element",
				data:   []int{1},
				remove: 2,
			},
			{
				name:   "multiple elements",
				data:   util.GenerateIntSeries(1, 1000),
				remove: 200,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				q := New[int]()

				var wg sync.WaitGroup

				for _, data := range tt.data {
					wg.Add(1)
					go func(data int) {
						q.Push(data)
						wg.Done()
					}(data)
				}

				var removecount atomic.Int64
				for i := 0; i < tt.remove; i++ {
					wg.Add(1)
					go func() {
						if _, ok := q.Pop(); ok {
							removecount.Add(1)
						}

						wg.Done()
					}()
				}

				wg.Wait()

				var finalcount int64
				for range iterQueue(q.head.Load().Next.Load()) {
					finalcount++
				}

				if (int64(len(tt.data)) - finalcount) != removecount.Load() {
					t.Errorf("got %v, want %v", finalcount, removecount.Load())
				}
			})
		}
	})
}

func Test_Queue_Peek(t *testing.T) {
	t.Run("single threaded", func(t *testing.T) {
		type test struct {
			name string
			data []int
			peek [2]interface{}
		}

		datasets := map[string][]int{
			"empty":    {},
			"single":   {1},
			"multiple": util.GenerateRandomIntSeries(1, 1000),
		}

		tests := []test{
			{
				name: "empty",
				data: datasets["empty"],
				peek: [2]interface{}{0, false},
			},
			{
				name: "single element",
				data: datasets["single"],
				peek: [2]interface{}{datasets["single"][0], true},
			},
			{
				name: "multiple elements",
				data: datasets["multiple"],
				peek: [2]interface{}{datasets["multiple"][0], true},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				q := New[int]()

				for _, data := range tt.data {
					q.Push(data)
				}

				if got, ok := q.Peek(); got != tt.peek[0] || ok != tt.peek[1] {
					t.Errorf("got %v, want %v", got, tt.peek)
				}
			})
		}
	})
}

func iterQueue[T any](node *Node[T]) chan T {
	ch := make(chan T)

	go func() {
		for node != nil {
			ch <- node.Data
			node = node.Next.Load()
		}
		close(ch)
	}()

	return ch
}
