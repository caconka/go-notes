package main

import (
	"fmt"
)

// Devuelve una función que devuelve un string
func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func main() {
	// a y b son closures, son dos funciones con un valor de <t> independiente
	// A la vez son First Class Functions, a una variable se le asigna una función
	a := appendStr()
	b := appendStr()

	fmt.Println(a("World"))    // t = Hello World
	fmt.Println(b("Everyone")) // t = Hello Everyone

	fmt.Println(a("Gopher")) // t = Hello World Gopher
	fmt.Println(b("!"))      // t = Hello Everyone !
}
