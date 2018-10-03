package main

import "oop/employee"

func main() {
	// instanciamos un employee
	e := employee.New("Roberto", "Gómez", 34, 22)
	e.LeavesRemaining()
}
