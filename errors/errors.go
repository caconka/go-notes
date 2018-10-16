package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

// Los errores en go son una interfaz
type error interface {
	Error() string
}

func assertion(p string) {
	f, err := os.Open(p)
	if err, ok := err.(*os.PathError); ok {
		fmt.Println("File at path", err.Path, "failed to open")
		return
	}
	fmt.Println(f, "opened successfully")
}

func methods(h string) {
	addr, err := net.LookupHost(h)
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println("temporary error")
		} else {
			fmt.Println("generic error: ", err)
		}
		return
	}
	fmt.Println(addr)
}

func directComp(p string) {
	// La función Glob devuelve los nombres d elos archivos que coincidan
	// var ErrBadPattern = errors.New("syntax error in pattern")
	files, error := filepath.Glob(p)
	if error != nil && error == filepath.ErrBadPattern {
		fmt.Println(error)
		return
	}
	fmt.Println("matched files", files)
}

func main() {
	f, err := os.Open("/test.txt")
	// 1. Assertion: os.PathError es el tipo del error que da os.Open(), y este
	// tiene (Path string) para conseguir más info del error
	assertion("/test.txt")

	// 2. Assertion y obtener información usando los métodos del tipo
	methods("falsehostexample.com")

	// 3. Direct comparison
	directComp("[")

	// Forma convencional de comprobar si hay error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}
