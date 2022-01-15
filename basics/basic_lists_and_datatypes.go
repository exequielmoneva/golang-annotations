package main

import (
	"fmt"
	"strings"
)

func main() {
	// split y append sobre listas
	test := "Boca la concha de tu madre"
	fmt.Println("Boca la concha de tu madre")
	test2 := strings.Split(test, " ")
	test2 = append(test2, "ganen", "un", "puto", "partido")
	for _, element := range test2 {
		fmt.Println(element)
	}
	fmt.Println(test2)
	// punteros: & se llama reference operator
	// &nombreDeVariable permite guardar la posicion de memoria de la variable nombrada
	test3 := &test
	fmt.Println(test3)
	// para recibir el valor de la posicion de memoria se usa * delante de la variable que tiene el puntero (test3)
	fmt.Println(*test3)
}
