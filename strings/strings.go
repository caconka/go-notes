package main

import "fmt"

// Los strings son inmutables en go, son un slice of bytes
// Hay caracteres que ocupan más de un byte, como la ñ. Por ello existe el tipo
// rune, que es un alias de int32

func printCharsAndBytes(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}

func main() {
	name := "Señor" // la ñ ocupa 2 bytes
	printCharsAndBytes(name)
	h := "hello"
	fmt.Println(mutate([]rune(h)))
	// fmt.Println(mutate(h)) // error strings son inmutables
}
