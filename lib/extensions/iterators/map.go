package iterators

import "github.com/tomascpmarques/functional-ext/lib/extensions"

type IterableMap[T comparable, K comparable] struct {
	count uint
	s     map[T]K
}

/*
Create a new iterator dor a map with the received values
*/
func IteratorMap[T comparable, K comparable](k ...extensions.Tuple[T, K]) *IterableMap[T, K] {
	lex := make(map[T]K)
	for _, tuple := range k {
		lex[tuple.K] = tuple.V
	}
	return &IterableMap[T, K]{0, lex}
}

func (im *IterableMap[T, K]) Pop() T {
	panic("not implemented") // TODO: Implement
}

func (im *IterableMap[T, K]) Next() T {
	panic("not implemented") // TODO: Implement
}

func (im *IterableMap[T, K]) Len() uint {
	panic("not implemented") // TODO: Implement
}

func (im *IterableMap[T, K]) Push(value ...T) {
	panic("not implemented") // TODO: Implement
}

func (im *IterableMap[T, K]) Get(pos int) (T, error) {
	panic("not implemented") // TODO: Implement
}

func (im *IterableMap[T, K]) Set(pos int, value T) error {
	panic("not implemented") // TODO: Implement
}

func (im *IterableMap[T, K]) Map(f extensions.FBack[T, K]) {
	panic("not implemented") // TODO: Implement
}

func (im *IterableMap[T, K]) ForEach(f extensions.FBack[T, K]) IterableSlice[K] {
	panic("not implemented") // TODO: Implement
}
