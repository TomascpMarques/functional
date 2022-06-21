package iterators

import (
	"fmt"

	"github.com/tomascpmarques/functional-ext/lib/extensions"
)

type IterableSlice[T comparable] struct {
	count uint
	s     []T
}

/*
Creates a new iterator for a slice with the received values
*/
func IteratorSlice[T comparable](c ...T) *IterableSlice[T] {
	return &IterableSlice[T]{0, c}
}

/* Default implementation for existing base iterator type */
func (i IterableSlice[T]) Get(pos int) (T, error) {
	if pos >= len(i.s) || pos < 0 {
		// Cant return nil due to generic type T being used
		return *new(T), fmt.Errorf("index out of bounds: %d (must be >= 0 and <= %d)", pos, i.Len())
	}
	return i.s[pos], nil
}

func (i *IterableSlice[T]) Set(pos int, value T) error {
	if pos >= len(i.s) || pos < 0 {
		return fmt.Errorf("cant set a value on an invalid index: %d (must be between 0 and %d)", pos, i.Len())
	}
	(*i).s[pos] = value
	return nil
}

func (i *IterableSlice[T]) Push(value ...T) {
	i.s = append(i.s, value...)
}

func (i *IterableSlice[T]) Pop() T {
	lex := (*i).s[i.Len()-1]
	i.s = i.s[:i.Len()-1]
	return lex
}

func (i *IterableSlice[T]) Next() T {
	temp := i.s[i.count]
	i.count += 1 % uint(len(i.s))
	return temp
}

func (i *IterableSlice[T]) Len() uint {
	return uint(len(i.s))
}

func (i *IterableSlice[T]) ForEachMut(f extensions.FBack[T, T]) {
	for k, v := range i.s {
		i.s[k] = f(v)
	}
}

func (i IterableSlice[T]) ForEach(f extensions.FBack[T, T]) IterableSlice[T] {
	lex := IterableSlice[T]{0, make([]T, i.Len())}
	for k, v := range i.s {
		lex.s[k] = f(v)
	}
	return lex
}
