package main

import (
	"fmt"
	"sync"
)

// Con defer hacemos que una función se ejecute justo antes del return

type rect struct {
	length int
	width  int
}

func (r rect) area(wg *sync.WaitGroup) {
	defer wg.Done() // con esta linea evitamos poner los otros 3 wg.Done()
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		// wg.Done()
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		// wg.Done()
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
	// wg.Done()
}

func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}

// El orden ejecución cuando tenemos varios defer es LIFO (Last In First Out)
func multipleDefer() {
	name := "Naveen"
	fmt.Printf("Orignal String: %s\n", name) // 'Naveen'
	fmt.Printf("Reversed String: ")          // 'neevaN'
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}
}

func main() {
	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")

	// Orden de ejecución LIFO
	multipleDefer()

	a := 5
	defer printA(a) // Se ejecutará lo último, al final de main(), por lo que
	// primero tendremos que a es 10 y luego se emprimirá el a = 5 que vale aquí
	a = 10
	fmt.Println("value of a before deferred function call", a)
}
