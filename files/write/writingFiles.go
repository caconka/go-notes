package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}

func consume(data chan int, done chan bool) {
	f, err := os.Create("concurrent.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}

func main() {
	// Crea el archivo en el directorio dónde estamos ejecutando el programa
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := f.WriteString("Hello World")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	fmt.Println(l, " bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	// ------ //
	// Escribir línea a línea
	f2, err := os.Create("lines.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	d := []string{
		"Welcome to the world of Go1",
		"Welcome to the world of Go2",
		"Welcome to the world of Go3",
	}

	for _, v := range d {
		fmt.Fprintln(f2, v) // escribe un string en un archivo
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	err = f2.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File written successfully")

	// ------ //
	// Añadir a un archivo ya existente
	// el segundo parámetro es el tipo de modo en el que se abre el archivo
	f3, err := os.OpenFile("lines.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
		return
	}

	newLine := "File handling is easy"
	_, err = fmt.Fprintln(f3, newLine)
	if err != nil {
		fmt.Println(err)
		f3.Close()
		return
	}

	err = f3.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File appended successfully")

	// ------ //
	// Escribir en un mismo archivo con 100 goroutines (concurrencia)
	// Podemos comprobarlo abriendo tras ejecutar el archivo y ver las líneas escritas
	data := make(chan int)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}
	go consume(data, done)
	go func() {
		wg.Wait()
		close(data)
	}()
	d2 := <-done
	if d2 == true {
		fmt.Println("File written successfully")
	} else {
		fmt.Println("File writing failed")
	}

}
