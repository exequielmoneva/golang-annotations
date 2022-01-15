package main

import (
	"fmt"
)

// Declaro la interfaz

type Persona interface {
	Descripcion()
}

type Jugador struct {
	Nombre   string
	Apellido string
	Altura   float32
	Posicion string
}
type Dt struct {
	Nombre   string
	Apellido string
}

// Para forzar a una estructura a implementar todos los metodos de la interfaz
func NewDt(name string, lastname string) Persona {
	return &Dt{name, lastname}
	// si GetName no esta implementado para Dt ocurrira el siguiente error
	// cannot use &Dt literal (type *Dt) as type Persona in return argument:
	//	*Dt does not implement Persona (missing Descripcion method)
}

func (jugador Jugador) GetPosicion() string {
	return jugador.Posicion
}

func (directorTecnico Dt) GetName() string {
	return directorTecnico.Nombre + " " + directorTecnico.Apellido + ","
}

func (jugador Jugador) Ficha() {
	fmt.Println("Ficha tecnica:")
	fmt.Println("Nombre completo:", jugador.Nombre, jugador.Apellido)
	fmt.Println("Altura:", jugador.Altura)
	fmt.Println("Posición:", jugador.Posicion)
}

// Declaro el uso de la interfaz

func (mp *Jugador) Descripcion() {
	mp.Ficha()
	fmt.Println("El mejor", mp.GetPosicion(), "del futbol argentino")
}

func (cb *Dt) Descripcion() {
	fmt.Println()
	fmt.Println("Nombre:", cb.Nombre, cb.Apellido)
	fmt.Println(cb.GetName(), "el mejor director tecnico del futbol sudamericano")
}

// Cuando no se o no quiero especificar el tipo de los datos, le declaro una interfaz vacia

func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main() {
	// Interfaces y estructuras
	mp := Jugador{"Martin", "Palermo", 1.9, "Delantero centro"}
	cb := Dt{"Carlos", "Bianchi"}
	mp.Descripcion()
	cb.Descripcion()
	Anything(cb)
	Anything(mp)
	Anything(2.4)

}
