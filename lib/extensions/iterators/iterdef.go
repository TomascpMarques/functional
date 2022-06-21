package iterators

import (
	"reflect"

	"github.com/tomascpmarques/functional/lib/extensions"
)

/* Iterator spec defeiniton */
type IterSlice[T comparable, K comparable] interface {
	Pop() T
	Next() T
	Len() uint
	Push(value ...T)
	Get(pos int) (T, error)
	Set(pos int, value T) error
	Map(f extensions.FBack[T, K])
	ForEach(f extensions.FBack[T, K]) IterSlice[T, K]
}

type IterMap[T comparable, K comparable] interface {
	Pop() T
	Next() K
	Len() uint
	Set(pos T, value K)
	Get(pos T) (K, error)
	Keys() []reflect.Value
	Map(f extensions.FBack[T, K])
	Push(value ...extensions.Tuple[T, K])
	ForEach(f extensions.F[T])
	ForEachNew(f extensions.FBack[T, K]) IterMap[T, K]
}
