package main

import (
	"fmt"
	"reflect"
)

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

	// Maps: son igual que los slices pero con key:value
	// personSalary := map[string]int{} // esto crea un map vacio
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}

	value, ok := personSalary["jamie"] // value = 15000, ok = true
	// value, ok := personSalary["john"] // value = 15000, ok = false
	if ok {
		fmt.Println("Salary jamie: ", value)
	}

	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	map2 := map1
	// Para comprar maps utilizamos el paquete reflect
	eq := reflect.DeepEqual(map1, map2)
	fmt.Println("map1 equal map2 =", eq)
}
