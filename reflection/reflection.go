package main

import (
	"fmt"
	"reflect"
)

type order struct {
	orderId    int
	customerId int
}

func createQuery(q interface{}) {
	// reflect.ValueOf() devuelve el tipo, en este caso sería main.order
	// mientras que la función Kind() devuelve el tipo específico
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field: %d | type: %T | value: %v\n", i, v.Field(i), v.Field(i))
		}
	}

}

func main() {
	o := order{
		orderId:    456,
		customerId: 56,
	}
	createQuery(o)
}
