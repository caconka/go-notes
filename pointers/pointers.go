package main

import "fmt"

func change(val *int) {
	*val = 55 // para acceder al valor de un pointer se hace con *
	// val = 55 // error
}

func modify(sls []int) {
	sls[0] = 90
}

func main() {
	a := 58
	fmt.Println("value of a before function call is", a)
	b := &a // la variable b almacena la dirección de a, no su valor
	change(b)
	fmt.Println("value of a after function call is", a)

	c := [3]int{89, 90, 91}
	modify(c[:]) // no le tenemos que pasar la referencia, para trabajar con
	// arrays están los slices
	fmt.Println(c)
}
