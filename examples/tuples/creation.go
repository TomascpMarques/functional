package main

import (
	"fmt"

	"github.com/tomascpmarques/functional-ext/lib/extensions"
)

func main() {
	// Creating a new tuple throught a constructor
	// With  type parameters
	_ = extensions.NewTuple[int, string](1, "hello")
	// Without type parameters (more idiomaticly correct)
	_ = extensions.NewTuple(1, "hello")

	// Creating a new tuple throught direct from the type
	_ = extensions.Tuple[int, string]{1, "hello"}
	//  Or (more idiomaticly correct)
	lex := extensions.Tuple[int, string]{K: 1, V: "hello"}

	// Accessing the Key and Value in Tuples
	_ = lex.K // The key
	_ = lex.V // The value

	fmt.Println("This example is only code defenition, no code result is visible, check the example file for code.")
}
