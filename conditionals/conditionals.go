package main

import "fmt"

// if statement ; condition {...}

func number() int {
	num := 15 * 5
	return num
}

func main() {
	// num es sólo accesible dentro de este if else
	if num := 10; num%2 == 0 {
		fmt.Println(num, " is even")
	} else {
		fmt.Println(num, " is odd")
	}

	// fallthrough evita que se pare la ejecución, sigue al 'case' siguiente
	switch num := number(); { // num no es una constante
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}
}
