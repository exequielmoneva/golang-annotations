package main

import "fmt"

func main() {
	// control structures son: if else, for, switch case, break continue
	// !Variable para negar y decir que es igual a false, && para concatenar sentencias y || (alt+124) para los or
	var lista = []string{"casa", "boca", "raiz"}
	var day = "Viernes"
	var day2 = "Miercoles"
	var caso = 5
	for i := 0; i < 10; i++ {
		if i == 0 {
			continue
		} else if i%2 == 0 {
			fmt.Println(i)
		} else {
			fmt.Println("Not even")
		}
	}
	for idx, value := range lista {
		fmt.Println(idx, value)
	}
	// i recibe la posicion de la letra en el string, ch recibe el valor unicode de esa letra (ej: Alt+86 es V)
	for i, ch := range "Viaje" {
		fmt.Printf("%#U starts at byte position %d\n", ch, i)
	}
	switch day {
	case "Lunes":
		fmt.Println("Que bajon, hoy es Lunes")
	case "Viernes":
		fmt.Println("Por fin es Viernes! se termina esta semana de mierda")
	default:
		fmt.Println("Que larga se hace la semana")
	}
	switch day2 {
	case "Lunes":
		fmt.Println("Que bajon, hoy es Lunes")
	case "Martes", "Miercoles", "Jueves":
		fmt.Println("Que larga se hace la semana")
	case "Viernes":
		fmt.Println("Por fin es Viernes! se termina esta semana de mierda")
	default:
		fmt.Println("Fin de semana")
	}
	switch {
	case caso > 5:
		fmt.Println("Es mayor a 5")
	case caso < 6:
		fmt.Println("Es menor a 5")
	case caso == 5:
		fmt.Println("Viernes")
	default:
		fmt.Println("Es 5")
	}
}
