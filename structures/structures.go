package main

import "fmt"

type Address struct {
	city, state string // Se pueden agrupar las definiciones del mismo tipo
}

type Employee struct {
	firstName string
	lastName  string
	age       int
	address   Address
}

type Employee2 struct {
	FirstName string
	LastName  string
	age       int // sólo se exportan las que van en mayúsculas
	Address
}

func main() {
	emp1 := Employee{
		lastName:  "Pérez",
		age:       39,
		firstName: "Luis",
		address: Address{
			city:  "Parla",
			state: "Madrid",
		},
	}
	emp2 := Employee{"Ana", "Gómez", 29, Address{"Loredo", "Santander"}}
	// Tiene que seguir el mismo orden de la definición de la estructura
	emp3 := Employee{
		firstName: "Pedro",
		address: Address{
			state: "Barcelona",
		},
	} // los demás campos serán iniciados a 0 en el caso de int y vacios en string
	fmt.Println(emp1, emp2, emp3, emp3.address.state)

	emp4 := Employee2{
		LastName:  "Sánchez",
		FirstName: "Claudia",
		age:       32,
		Address: Address{ // Address tiene que ser en mayúscula al declararse
			// porque no se ha declarado nombre del campo. Ahora city y state estarán
			// al mismo nivel que los anteriores
			city:  "Ponferrada",
			state: "León",
		}}
	fmt.Println(emp4.FirstName, "'s city is", emp4.city)
}
