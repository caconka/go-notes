package main

import "fmt"

/*
b := [3]int́{5,10,2} // [5,10,2]
var c [5]int // [0,0,0,0,0]
len(b) // 3
*/

/*
Slices: hacen referencia a los valores de un array, ellos no tienen valores

d := [5]int{76, 77, 78, 79, 80} // esto crea un array y devuelve un slice que
es almacenado en 'd'

var e []int = d[1:4] // [77, 78, 79]
*/

func main() {
	// [...] calcula el length del array automáticamente
	a := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range a { // range devuelve el index y el value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a", sum)
}
