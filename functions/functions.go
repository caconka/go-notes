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

func main() {
		// _ Blank identifier: desecha un valor retornado, en este caso es 'perimeter'
	area, _ := rectProps(10.8, 5.6)
	fmt.Printf("Area %f", area)
}
