package extensions

type Tuple[k comparable, v comparable] struct {
	K k
	V v
}

func NewTuple[T comparable, K comparable](k T, v K) Tuple[T, K] {
	return Tuple[T, K]{k, v}
}
