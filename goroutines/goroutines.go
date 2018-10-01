package main

import (
	"fmt"
	"time"
)

// goroutine -> para mejorar la concurrencia
// channel -> para comunicarse con los goroutines (sólo admiten un tipo)

/*
data := <- a // read from channel a
a <- data // write to channel a
*/

/*
Los channels tienen capacidad y por defecto es 0. Si se sobrepasa la capacidad
entramos en un 'Deadlock'
ch := make(chan string, 3)
ch <- "Jose" // len(ch) = 1
ch <- "Lola" // len(ch) = 2
<- ch // len(ch) = 1
*/

func hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

// la función main tiene su propio goroutine "main goroutine"
func main() {
	done := make(chan bool) // declaración de un channel que admite booleanos
	fmt.Println("Main going to call hello go goroutine")
	go hello(done)
	<-done // bloquea la ejecución hacia abajo
	fmt.Println("Main received data")

	ch := make(chan int) // un channel bidireccional puede ser convertido a
	// unidireccional pero no al revés
	go producer(ch)
	/*
		for {
			v, ok := <-ch
			if ok == false { // ok será false cuando se cierre el channel close(chnl)
				break
			}
			fmt.Println("Received ", v, ok)
		}
	*/
	for v := range ch {
		fmt.Println("Received ", v)
	}
}
