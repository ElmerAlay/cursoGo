package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	http.HandleFunc("/saludo", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "What´s up")
	})

	http.ListenAndServe(":8000", nil)
}

//Siempre deben llevar esos 2 parámetros para hacer
//una petición web
func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}
