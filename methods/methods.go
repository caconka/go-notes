package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	length int
	width  int
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

// Los métodos se pueden llamar igual porque pertenecen a tipos diferentes
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

/* ----------------------------------------- */

type Employee struct {
	name string
	age  int
}

// Method with value receiver
func (e Employee) changeName(newName string) {
	e.name = newName
}

// Method with pointer receiver
func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

/* ----------------------------------------- */
// Como en el tipo int no se pueden agregar, podemos hacer custom types
type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

func main() {
	r := Rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("Area of rectangle %d\n", r.Area())
	c := Circle{
		radius: 12,
	}
	fmt.Printf("Area of circle %f", c.Area())

	/* ----------------------------------------- */
	// Diferencia entre método como puntero o como valor
	e := Employee{
		name: "Mark Andrew",
		age:  50,
	}
	fmt.Printf("̀\nEmployee name before change: %s", e.name)
	e.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	e.changeAge(51) // es lo mismo a (&e).changeAge(51)
	fmt.Printf("\nEmployee age after change: %d", e.age)

	/* ----------------------------------------- */
	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println("Sum is", sum)
}
