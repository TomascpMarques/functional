package main

import (
	"fmt"

	"github.com/tomascpmarques/frontline/lib/html/elements"
)

func main() {
	lex := elements.NewP("Lorem ipsum dolor sit amet")
	ley := elements.NewP("yesss")
	ley.SetAttributes(map[string]string{"class": "lorem"})

	lex.PushNewElement(ley)

	fmt.Printf("-> %s", lex.MarkItUp())
}
