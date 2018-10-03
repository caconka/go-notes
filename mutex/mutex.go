package main

import (
	"fmt"
	"sync"
)

/*
CRITICAL SECTION:
Cuando un programa se ejecuta simultáneamente, las partes de código que
modifican recursos compartidos no deben ser accedidas por varios Goroutines al
mismo tiempo. A esta sección del código que modifica los recursos compartidos
se denomina critical section
*/

/*
RACE CONDITION:
Cuando la salida/resultado depende de una secuencia de eventos que se ejecutan
en un orden arbitrario y van a trabajar sobre el mismo recurso compartido, se
puede producir un error cuando cuando dichos eventos no llegan en el orden que
esperamos. Dos procesos compiten a la carrera por llegar antes que el otro por
lo que la salida será diferente dependiendo cual llegó antes
*/

var x = 0
var y = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()   // bloquea para que sea accesible sólo por un goroutine a la vez
	x = x + 1  // recurso compartido
	m.Unlock() // se necesita indicar dónde se desbloquea
	wg.Done()
}

/* --------------------------------------------- */
// Solución al race condition vía channels
func incrementChan(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	y = y + 1 // como el channel tiene capacidad de 1, sólo un goroutine podrá
	// acceder cada vez
	<-ch
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m) // se le pasa el mutex como pointer porque si no cada
		// goroutine tendría su propio mutex
	}
	w.Wait()
	fmt.Println("final value of x", x) // sin mutex el resultado de x varía, con
	// mutex nos aseguramos de que llegue a 1000 en este caso, porque sólo un
	// goroutine incrementará la x cada vez

	/* --------------------------------------------- */
	// Solución channels
	var w2 sync.WaitGroup
	ch := make(chan bool, 1) // se crea un channel con capacidad de 1
	for i := 0; i < 900; i++ {
		w2.Add(1)
		go incrementChan(&w2, ch)
	}
	w2.Wait()
	fmt.Println("final value of y", y)
}

/*
MUTEX vs CHANNELS
En general, se utilizan los channels cuando los goroutines necesitan comunicarse
entre sí y los mutexes cuando sólo un goroutine debe acceder al critical section
*/
