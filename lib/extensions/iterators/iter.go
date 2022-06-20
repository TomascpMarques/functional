package iterators

import "github.com/tomascpmarques/functional-ext/lib/extensions"

/* Iterator spec defeiniton */
type Iter[T comparable, K comparable] interface {
	Pop() T
	Next() T
	Len() uint
	Push(value ...T)
	Get(pos int) (T, error)
	Set(pos int, value T) error
	Map(f extensions.FBack[T, K])
	ForEach(f extensions.FBack[T, K]) IterableSlice[K]
}
