package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gobuffalo/packr"
)

func main() {
	// Usando el paquete flag para acceder al archivo
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	// data es un byte slice que se transforma a string con la función
	fmt.Println("Contents of file:", string(data))

	// ------ //
	// Empaquetar en el binario el archivo buildingTest.txt
	// Para hacer esto se debe usar el comando `packr install -v readingFiles`
	// sino seguirá comportándose de la misma manera, yendo a por el archivo
	// local
	box := packr.NewBox("../read")
	dataBuilding := box.String("buildingTest.txt")
	fmt.Println("Contents of file:", dataBuilding)

	// ------ //
	// Leer el archivo en trozos, cuando los archivos son muy grandes
	fmt.Printf("Leyendo el archivo en trozos")
	fptr2 := flag.String("fpath2", "test.txt", "file path to read from")
	flag.Parse()

	// Abrimos el archivo con el flag que hemos definido antes
	f, err := os.Open(*fptr2)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	r := bufio.NewReader(f)
	b := make([]byte, 3) // se define la cantidad de bytes a leer cada vez
	for {
		n, err := r.Read(b)
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Println(string(b[0:n]))
	}

	// ------ //
	// Leer un archivo linea a linea
	fmt.Printf("Leyendo el archivo linea a linea")
	fptr3 := flag.String("fpath3", "lineByLine.txt", "file path to read from")
	flag.Parse()

	f2, err := os.Open(*fptr3)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f2.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f2) // abrimos un scanner
	for s.Scan() {            // el método Scan() lee la próxima línea
		fmt.Println(s.Text()) // Text() nos devuelve esa línea
	}

	err = s.Err() // comprobamos si nos da algún error al escanear
	if err != nil {
		log.Fatal(err)
	}
}
