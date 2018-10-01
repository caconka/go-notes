package main

import (
	"fmt"
	"geometry/rectangle" // importamos un paquete desde src/

	// rect "geometry/rectangle" - Le podemos poner un nombre diferente al paquete
	// _ "geometry/rectangle" - Si queremos que sólo se ejecute su init function
	"log"
)

/*
Orden de inicialización
1. Package level variables
2. Init functions

Si un paquete importa otros paquetes estos se importan antes, con este orden
*/

// 1. package variables
var rectLen, rectWidth float64 = 6, 7

// 2. init function
func init() {
	println("main package initialized")
	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

func main() {
	fmt.Println("Geometrical shape properties")
	// 'rectangle' es el nombre del paquete y 'Area' y 'Diagonal' son exportadas
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	fmt.Printf("diagonal of the rectangle %.2f ", rectangle.Diagonal(rectLen, rectWidth))
}
