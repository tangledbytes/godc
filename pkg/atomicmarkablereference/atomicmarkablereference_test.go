package atomicmarkablereference_test

import (
	"runtime"
	"sync"
	"testing"

	"github.com/utkarsh-pro/godc/pkg/atomicmarkablereference"
)

func Test_AtomicMarkableReference_New(t *testing.T) {
	type test struct {
		name string
		ref  *int
		mark bool

		wantref  *int
		wantmark bool
	}

	refs := map[string]*int{
		"ref1": new(int),
		"ref2": new(int),
	}

	tests := []test{
		{
			name: "when ref is nil",
			ref:  nil,
			mark: false,

			wantref:  nil,
			wantmark: false,
		},
		{
			name: "when ref is not nil and mark is false",
			ref:  refs["ref1"],
			mark: false,

			wantref:  refs["ref1"],
			wantmark: false,
		},
		{
			name: "when ref is not nil and mark is true",
			ref:  refs["ref2"],
			mark: true,

			wantref:  refs["ref2"],
			wantmark: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mref := atomicmarkablereference.New(tt.ref, tt.mark)

			if mref.GetReference() != tt.wantref {
				t.Errorf("got %v, want %v", mref.GetReference(), tt.wantref)
			}

			if mref.IsMarked() != tt.wantmark {
				t.Errorf("got %v, want %v", mref.IsMarked(), tt.wantmark)
			}
		})
	}
}

func Test_AtomicMarkableReference_CompareAndSet(t *testing.T) {
	type test struct {
		name string
		old  *int
		new  *int
		mark bool

		wantref  *int
		wantmark bool
	}

	refs := map[string]*int{
		"ref1": new(int),
		"ref2": new(int),
	}

	tests := []test{
		{
			name: "when old ref is nil",
			old:  nil,
			new:  refs["ref1"],
			mark: false,

			wantref:  refs["ref1"],
			wantmark: false,
		},
		{
			name: "when old ref is not nil and mark is false",
			old:  refs["ref1"],
			new:  refs["ref2"],
			mark: false,

			wantref:  refs["ref2"],
			wantmark: false,
		},
		{
			name: "when old ref is not nil and mark is true",
			old:  refs["ref2"],
			new:  refs["ref1"],
			mark: true,

			wantref:  refs["ref1"],
			wantmark: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mref := atomicmarkablereference.New(tt.old, tt.mark)
			mref.CompareAndSet(tt.old, tt.new, tt.mark, tt.mark)

			if mref.GetReference() != tt.wantref {
				t.Errorf("got %v, want %v", mref.GetReference(), tt.wantref)
			}

			if mref.IsMarked() != tt.wantmark {
				t.Errorf("got %v, want %v", mref.IsMarked(), tt.wantmark)
			}
		})
	}
}

func Test_AtomicMarkableReference_GetReference(t *testing.T) {
	type test struct {
		name string
		ref  *int
		mark bool

		wantref *int
	}

	refs := map[string]*int{
		"ref1": new(int),
		"ref2": new(int),
	}

	tests := []test{
		{
			name: "when ref is nil",
			ref:  nil,
			mark: false,

			wantref: nil,
		},
		{
			name: "when ref is not nil and mark is false",
			ref:  refs["ref1"],
			mark: false,

			wantref: refs["ref1"],
		},
		{
			name: "when ref is not nil and mark is true",
			ref:  refs["ref2"],
			mark: true,

			wantref: refs["ref2"],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mref := atomicmarkablereference.New(tt.ref, tt.mark)

			if mref.GetReference() != tt.wantref {
				t.Errorf("got %v, want %v", mref.GetReference(), tt.wantref)
			}
		})
	}
}

func Test_AtomicMarkableReference_IsMarked(t *testing.T) {
	type test struct {
		name string
		ref  *int
		mark bool

		wantmark bool
	}

	refs := map[string]*int{
		"ref1": new(int),
		"ref2": new(int),
	}

	tests := []test{
		{
			name: "when ref is nil",
			ref:  nil,
			mark: false,

			wantmark: false,
		},
		{
			name: "when ref is not nil and mark is false",
			ref:  refs["ref1"],
			mark: false,

			wantmark: false,
		},
		{
			name: "when ref is not nil and mark is true",
			ref:  refs["ref2"],
			mark: true,

			wantmark: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mref := atomicmarkablereference.New(tt.ref, tt.mark)

			if mref.IsMarked() != tt.wantmark {
				t.Errorf("got %v, want %v", mref.IsMarked(), tt.wantmark)
			}
		})
	}
}

func Test_AtomicMarkableReference_Set(t *testing.T) {
	type test struct {
		name string
		ref  *int
		mark bool

		wantref  *int
		wantmark bool
	}

	refs := map[string]*int{
		"ref1": new(int),
		"ref2": new(int),
	}

	tests := []test{
		{
			name: "when ref is nil",
			ref:  nil,
			mark: false,

			wantref:  nil,
			wantmark: false,
		},
		{
			name: "when ref is not nil and mark is false",
			ref:  refs["ref1"],
			mark: false,

			wantref:  refs["ref1"],
			wantmark: false,
		},
		{
			name: "when ref is not nil and mark is true",
			ref:  refs["ref2"],
			mark: true,

			wantref:  refs["ref2"],
			wantmark: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mref := atomicmarkablereference.New[int](nil, false)
			mref.Set(tt.ref, tt.mark)

			if mref.GetReference() != tt.wantref {
				t.Errorf("got %v, want %v", mref.GetReference(), tt.wantref)
			}

			if mref.IsMarked() != tt.wantmark {
				t.Errorf("got %v, want %v", mref.IsMarked(), tt.wantmark)
			}
		})
	}
}

func Test_AtomicMarkableReference_Get(t *testing.T) {
	type test struct {
		name string
		ref  *int
		mark bool

		wantref  *int
		wantmark bool
	}

	refs := map[string]*int{
		"ref1": new(int),
		"ref2": new(int),
	}

	tests := []test{
		{
			name: "when ref is nil",
			ref:  nil,
			mark: false,

			wantref:  nil,
			wantmark: false,
		},
		{
			name: "when ref is not nil and mark is false",
			ref:  refs["ref1"],
			mark: false,

			wantref:  refs["ref1"],
			wantmark: false,
		},
		{
			name: "when ref is not nil and mark is true",
			ref:  refs["ref2"],
			mark: true,

			wantref:  refs["ref2"],
			wantmark: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mref := atomicmarkablereference.New(tt.ref, tt.mark)
			ref, mark := mref.Get()

			if ref != tt.wantref {
				t.Errorf("got %v, want %v", ref, tt.wantref)
			}

			if mark != tt.wantmark {
				t.Errorf("got %v, want %v", mark, tt.wantmark)
			}
		})
	}
}

func Test_AtomicMarkableReference_CompareAndSet_Multi(t *testing.T) {
	one := 1
	two := 2
	three := 3

	t.Run("reference is changed", func(t *testing.T) {
		ai := atomicmarkablereference.New(&one, false)

		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			for !ai.CompareAndSet(&two, &three, false, false) {
				runtime.Gosched()
			}
		}()

		if val := ai.CompareAndSet(&one, &two, false, false); !val {
			t.Errorf("got %v, want %v", val, true)
		}

		wg.Wait()

		if ai.GetReference() != &three {
			t.Errorf("got %v, want %v", ai.GetReference(), &three)
		}

		if ai.IsMarked() {
			t.Errorf("got %v, want %v", ai.IsMarked(), false)
		}
	})

	t.Run("mark is changed", func(t *testing.T) {
		ai := atomicmarkablereference.New(&one, false)

		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			for !ai.CompareAndSet(&one, &one, true, false) {
				runtime.Gosched()
			}
		}()

		if val := ai.CompareAndSet(&one, &one, false, true); !val {
			t.Errorf("got %v, want %v", val, true)
		}

		wg.Wait()

		if ai.GetReference() != &one {
			t.Errorf("got %v, want %v", ai.GetReference(), &one)
		}

		if ai.IsMarked() {
			t.Errorf("got %v, want %v", ai.IsMarked(), false)
		}
	})
}
