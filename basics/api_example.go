package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Para crear una Api necesitamos delarar Mux, que maneja las respuestas de los endpoint
	// se declaran mediante HandleFunc("ruta del endpoint", funcion)
	// La funcion requiere declarar ResponseWriter para la respuesta y Request para leer datos del request
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(("This is a " + r.Method + " method")))
	})
	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Boca la concha de tu madre"))
	})

	fmt.Println("Server is running on localhost:3000")
	http.ListenAndServe("localhost:3000", mux)
}
