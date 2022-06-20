package main

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/tomascpmarques/functional-ext/lib/extensions"
	"github.com/tomascpmarques/functional-ext/lib/extensions/iterators"
)

func main() {
	x := struct{ i int }{4}
	y := struct{ i int }{5}
	z := struct{ i int }{6}
	lay := iterators.IteratorSlice(x, y, z)
	lay.Push(struct{ i int }{7})
	fmt.Println(lay)

	lex := iterators.IteratorSlice(1, 2, 3, 4, 5, 6, 7)
	lex.Push(8)
	lex.Push(9)
	lex.Push(10, 11, 12, 13)
	fmt.Println(lex)
	fmt.Println(lex.Get(100))
	fmt.Println(lex.Get(4))
	lex.Set(4, 22)
	fmt.Println(lex)

	ley := []int{1, 2, 3, 4, 5, 6, 7, 8}
	lej := extensions.ApplyToNewSlice(
		&ley,
		func(i int) int {
			return i * 2
		},
	)
	fmt.Println(lej)

	sum := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}
	sumy := extensions.ApplyToNewMap(
		&sum,
		func(i int) string {
			return strconv.Itoa(i)
		},
	)
	fmt.Println(sumy)
	fmt.Println(reflect.TypeOf(sumy))

	sey := iterators.IteratorMap(
		extensions.NewTuple(1, 2),
		extensions.NewTuple(2, 3),
		extensions.NewTuple(4, 5),
		extensions.NewTuple(6, 7),
	)
	fmt.Println(sey)
}
