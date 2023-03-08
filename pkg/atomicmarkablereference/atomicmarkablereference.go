package atomicmarkablereference

import "sync/atomic"

// AtomicMarkableReference is a reference to an object that may be marked
// as "in use", in which case attempted atomic operations on it will fail.
//
// This is a port of the Java AtomicMarkableReference class.
type AtomicMarkableReference[T any] struct {
	pair atomic.Pointer[Pair[T]]
}

func New[T any](reference *T, mark bool) *AtomicMarkableReference[T] {
	apair := atomic.Pointer[Pair[T]]{}
	apair.Store(NewPair(reference, mark))

	return &AtomicMarkableReference[T]{pair: apair}
}

// CompareAndSet sets the value to the given updated value if the
// current value == the expected value and the current mark == the
// expected mark. This operation is atomic (or so the Java docs claims).
func (amr *AtomicMarkableReference[T]) CompareAndSet(
	oldref *T,
	newref *T,
	oldmark bool,
	newmark bool,
) bool {
	curr := amr.pair.Load()

	if oldref != curr.Reference && oldmark != curr.Mark {
		return false
	}

	if curr.Reference == newref && curr.Mark == newmark {
		return true
	}

	return amr.pair.CompareAndSwap(curr, NewPair(newref, newmark))
}

// GetReference returns the current reference.
func (amr *AtomicMarkableReference[T]) GetReference() *T {
	return amr.pair.Load().Reference
}

// IsMarked returns the current mark.
func (amr *AtomicMarkableReference[T]) IsMarked() bool {
	return amr.pair.Load().Mark
}

// Set sets the reference and mark atomically.
func (amr *AtomicMarkableReference[T]) Set(newReference *T, newMark bool) {
	amr.pair.Store(NewPair(newReference, newMark))
}

// Get returns the current reference and mark.
func (amr *AtomicMarkableReference[T]) Get() (reference *T, mark bool) {
	pair := amr.pair.Load()
	return pair.Reference, pair.Mark
}
