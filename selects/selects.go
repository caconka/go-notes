package main

import (
	"fmt"
	"time"
)

/*
Los selects son usados para elegir uno de muchos send/receive channel operation,
son *bloqueantes* hasta que uno de los send/receive operation está listo
Si varios send/receive estan listos a la vez se eligirá uno al azar
*/
func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"

}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
		/*
			Se puede añadir un default que se ejecutará si los de arriba no están listos,
			en este caso siempre se ejecutaría porque tardan 3 y 6 segundos los servidores

			default:
				fmt.Println("default case executed")
		*/
	}
	fmt.Println("next") // no se ejecuta hasta que el select termine
}
