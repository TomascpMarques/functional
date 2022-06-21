package main

import (
	"fmt"
	"strconv"

	"github.com/tomascpmarques/functional-ext/lib/extensions"
	itrs "github.com/tomascpmarques/functional-ext/lib/extensions/iterators"
)

type IFace interface {
	SetSomeField(newValue string)
	GetSomeField() string
	Yes() IFace
}

type Implementation struct {
	someField string
}

func (i *Implementation) GetSomeField() string {
	return i.someField
}

func (i *Implementation) SetSomeField(newValue string) {
	i.someField = newValue
}

func (i *Implementation) Yes() IFace {
	return i
}

func Create() *Implementation {
	return &Implementation{someField: "Hello"}
}

func main() {
	lex := itrs.IteratorMap(
		extensions.NewTuple(1, 2),
		extensions.NewTuple(2, 6),
		extensions.NewTuple(3, 3),
		extensions.NewTuple(4, 2),
		extensions.NewTuple(5, 7),
		extensions.NewTuple(6, 7),
	)

	fmt.Println(lex)
	fmt.Println()
	lex.Set(2, 8)
	fmt.Println(lex)
	fmt.Println()
	fmt.Println(lex.Get(3))
	lex.ForEach(func(i int) int {
		return i % 2
	})
	lay := itrs.ApplyToNewIterMap(lex, func(i int) string {
		return strconv.Itoa(i*2 - 9)
	})
	fmt.Printf("\n<->%v \n", lay)
	fmt.Println()
	fmt.Println()
	fmt.Printf("->%v \n", lex)
	lex.Map(func(i int) int {
		return i % 2
	})
	fmt.Println()
	fmt.Printf("->%v \n", lex)

	var a IFace = Create()
	a.SetSomeField("World")
	fmt.Println(a.GetSomeField())
}
