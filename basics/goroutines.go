package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func heavy(wg *sync.WaitGroup) {
	time.Sleep(time.Second * 5) // duerme por 5 segundos
	defer wg.Done()
}
func printers(wg *sync.WaitGroup) {
	for i, ch := range "Viaje" {
		fmt.Printf("%#U starts at byte position %d\n", ch, i)
	}
	// Defer se usa para asegurar que una función es llamada posteriormente durante la ejecución del programa.
	defer wg.Done()
}

func main() {
	// Wait groups
	wg := &sync.WaitGroup{}
	// add, done y wait
	// Add: digo cuantas Done debe esperar para seguir/terminar
	// Done: aviso que la funcion que sigue termino
	// Wait: aviso que espere para terminar. Lo que figura luego de Wait se corre cuando terminen las goroutines
	wg.Add(3)
	// Goroutines: usamos "go" frente a la llamada a una funcion para hacerla correr en el background
	go heavy(wg)
	go printers(wg)
	// Asi son las Lambda functions en golang, en este caso la ejecuto como goroutine pero no es obligatorio
	go func() {
		for _, i := range strings.Split("A Europa", "") {
			fmt.Println(i)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Para vivir")
}
