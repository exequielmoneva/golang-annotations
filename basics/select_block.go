package main

import (
	"fmt"
	"os"
	"time"
)

func Select(c chan int, quits chan struct{}) {
	// Select es el switch de los channels
	// es async
	for {
		time.Sleep(time.Second)
		select {
		case <-c:
			fmt.Println("recibido en c")
		case <-quits:
			fmt.Println("Quit")
			os.Exit(0) // para cerrar el programa. Elejimos 0 que significa safe
		}
	}
}

func main() {
	c := make(chan int)
	quits := make(chan struct{})

	go Select(c, quits)
	go func() {
		c <- 22
		quits <- struct{}{} // asis e declara una estructura vacia
	}()
	select {}
}
