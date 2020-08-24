package main

import (
	"endito/compiler"
	"endito/document"
)

func main() {
	d := document.FromDir("./example")
	d.Print()
	c := compiler.NewPiler(d)
	c.Compile()
}
