package main

import "fmt"

/* func <funcName> (paramName type) returnType ́{...} */

// Como 'area' y 'perimeter' están declarados como valores de retorno, no es
// necesario especificarlos en el return
func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return
}

// Variadic functions: admiten múltiples argumentos
func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func main() {
	// _ Blank identifier: desecha un valor retornado, en este caso es 'perimeter'
	area, _ := rectProps(10.8, 5.6)
	fmt.Printf("Area %f", area)

	fmt.Println("\n-- Variadic Functions --")
	nums := []int{89, 90, 95}
	find(89, nums...) // así le estamos pasando el contenido del slice
}
