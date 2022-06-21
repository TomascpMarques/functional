package iterators

import (
	"fmt"
	"reflect"

	"github.com/tomascpmarques/functional-ext/lib/extensions"
)

type IterableMap[T comparable, K comparable] struct {
	count *uint
	keys  []reflect.Value
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
	return &IterableMap[T, K]{new(uint), reflect.ValueOf(lex).MapKeys(), lex}
}

func (im IterableMap[T, K]) Pop() T {
	keys := reflect.ValueOf(im.s).MapKeys()
	wantted := keys[len(keys)-1]
	delete(im.s, wantted.Interface().(T))
	return wantted.Interface().(T)
}

func (im IterableMap[T, K]) Next() K {
	temp := im.s[im.keys[*im.count].Interface().(T)]
	*im.count += 1 % uint(len(im.keys))
	return temp
}

func (im IterableMap[T, K]) Len() uint {
	return uint(len(im.keys))
}

func (im IterableMap[T, K]) Push(value ...extensions.Tuple[T, K]) {
	for _, v := range value {
		im.s[v.K] = v.V
	}
}

func (im IterableMap[T, K]) Get(pos T) (K, error) {
	if v, ok := im.s[pos]; ok {
		return v, nil
	} else {
		return *new(K), fmt.Errorf("index out of bounds: %v (must be >= 0 and <= %d)", pos, im.Len())
	}
}

func (im IterableMap[T, K]) Set(pos T, value K) {
	im.s[pos] = value
}

func (im IterableMap[T, K]) Map(f extensions.FBack[T, K]) {
	for _, v := range im.keys {
		im.s[v.Interface().(T)] = f(v.Interface().(T))
	}
}

func (im IterableMap[T, K]) Keys() []reflect.Value {
	return im.keys
}

func (im IterableMap[T, K]) ForEach(f extensions.FBack[T, K]) IterMap[T, K] {
	lex := IterableMap[T, K]{new(uint), im.keys, make(map[T]K, len(im.keys))}
	for _, v := range im.keys {
		lex.s[v.Interface().(T)] = f(v.Interface().(T))
	}
	return lex
}

func ApplyToNewIterMap[K1 comparable, V1 comparable, V2 comparable](
	target *IterableMap[K1, V1],
	f extensions.FBack[V1, V2],
) IterMap[K1, V2] {
	return IterableMap[K1, V2]{new(uint), target.Keys(), func() map[K1]V2 {
		ley := make(map[K1]V2)
		for _, v := range target.Keys() {
			computedValue := v.Interface().(K1)
			if vl, err := target.Get(computedValue); err == nil {
				ley[computedValue] = f(vl)
			}
		}
		return ley
	}()}
}
