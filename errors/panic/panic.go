package main

import "fmt"

// * Debemos evitar en la medida de lo posible utilizar panic en favor de error

/*
Con el uso de panic la ejecución sigue hasta que todas las funciones del
goroutine actual hayan finalizado, momento en el cual el programa imprime el
panic. Por lo tanto cuando tenemos defer functions estas se imprimen antes
*/

/*
recover se utiliza para controlar los panics, sólo es útil dentro de una
función que tenga defer, y dentro del mismo goroutine. De esta forma devolvemos
el valor del panic y no pararía la ejecución, el problema es que perdemos el
rastro (trace). Podría ser como un try catch
Los panics se pueden dar también en tiempo de ejecución
*/
func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func fullName(firstName *string, lastName *string) {
	defer recoverName() // se ejecuta antes del panic, probar quitando esta llamada
	if firstName == nil {
		panic("runtime error: first name cannot be empty")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be empty")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}
