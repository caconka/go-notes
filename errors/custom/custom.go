package main

import (
	"fmt"
	"math"
)

// Por convención los errores deben terminar con la palabra Error
type areaError struct {
	err    string
	radius float64
}

// Implementa el método Error() para que sea un error
func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"radius is negative", radius}
	}
	return math.Pi * radius * radius, nil
}

func calcAreaCircl() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		// Comprobamos que el error sea de tipo areaError
		if err, ok := err.(*areaError); ok {
			fmt.Printf("Radius %0.2f is less than zero \n", err.radius)
			fmt.Println(err.err)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of rectangle1 %0.2f", area)
}

/* ---------------------------------------------- */
// Obtención de información vía métodos
type areaRectError struct {
	err    string
	length float64
	width  float64
}

func (e *areaRectError) Error() string {
	return e.err
}

func (e *areaRectError) lengthNegative() bool {
	return e.length < 0
}

func (e *areaRectError) widthNegative() bool {
	return e.width < 0
}

func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	if err != "" {
		return 0, &areaRectError{err, length, width}
	}
	return length * width, nil
}

func calcAreaRect() {
	length, width := -5.0, -9.0
	area, err := rectArea(length, width)
	if err != nil {
		if err, ok := err.(*areaRectError); ok {
			// Como el err es de tipo areaRectError tiene estos dos métodos
			if err.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", err.length)
			}
			if err.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", err.width)
			}
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("area of rect", area)
}

func main() {
	calcAreaCircl()
	calcAreaRect()
}
