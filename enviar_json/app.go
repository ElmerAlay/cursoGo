package main

import (
	"encoding/json"
	"net/http"
)

//Los atributos con letra mayúsculas son públicos
type Curso struct {
	Title    string
	NoVideos int
	//Si queremos que el json tenga un nombre en específico
	//hacemos lo siguiente
	//atib type `json:"titulo_a_mostrar"`
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", enviar)
	http.ListenAndServe(":8000", nil)
}

func enviar(w http.ResponseWriter, r *http.Request) {
	//Aquí todos llevan coma, incluso el último elemento
	cursos := Cursos{
		Curso{"Go", 30},
		Curso{"Go", 20},
		Curso{"Python", 10},
		Curso{"Json", 300},
	}

	//Si hay atrib no públicos, json no los muestra
	json.NewEncoder(w).Encode(cursos)
}
