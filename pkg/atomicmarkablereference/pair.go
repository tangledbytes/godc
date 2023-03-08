package atomicmarkablereference

type Pair[T any] struct {
	Reference *T
	Mark      bool
}

func NewPair[T any](reference *T, mark bool) *Pair[T] {
	return &Pair[T]{Reference: reference, Mark: mark}
}
