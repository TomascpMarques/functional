package main

import (
	"fmt"

	e "github.com/tomascpmarques/functional-ext/lib/extensions"
	i "github.com/tomascpmarques/functional-ext/lib/extensions/iterators"
)

func main() {
	/*
		Due to the way the IteratorMap is built
		There is no waning for restated keys, this
		situation will cause value re-writes
	*/
	lex := i.IteratorMap(
		e.NewTuple(1, 2),
		e.NewTuple(2, 6),
		e.NewTuple(3, 3),
		e.NewTuple(4, 2),
		e.NewTuple(5, 7),
		e.NewTuple(6, 7),
	)

	lex.ForEach(func(i int) {
		fmt.Printf("The value is %d\n", i)
	})
}
