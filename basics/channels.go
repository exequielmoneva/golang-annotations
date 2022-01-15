package main

import "fmt"

type Pais struct {
	nombre string
}

func main() {
	// los channels son colas
	// se declaran asi: <name> := make(chan <datatype>)
	// Ejemplo de unbuffered channel (solo puede tener un valor)
	c := make(chan int)
	// send
	go func() {
		c <- 1
	}()

	// snif
	val := <-c
	fmt.Println(val)

	// Ejemplo de Buffered channel (puede tener mas de un valor)
	ch := make(chan int, 3)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		close(ch) // Cerramos el channel luego de que no ingresan mas valores
	}()
	// Cerramos los channel para que el bucle for pueda saber hasta donde iterar
	for i := range ch {
		fmt.Println(i)
	}

	strch := make(chan *Pais, 3)
	go func() {
		strch <- &Pais{"Italia"}
		strch <- &Pais{"España"}
		strch <- &Pais{"Suiza"}
		strch <- &Pais{"Inglaterra"}
		strch <- &Pais{"Alemania"}
		close(strch)
	}()
	for i := range strch {
		fmt.Printf("%s\n", i.nombre)
	}
}
