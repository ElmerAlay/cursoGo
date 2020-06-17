package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
		//para obtener cualquier archivo dentro de
		//nuestro directorio
		//solo colocamos /nom_archivo
		//Empieza con 1 porque la pos 0 es /
		//http.ServeFile(w,r,r.URL.Path[1:])
	})
	http.ListenAndServe(":8000", nil)
}
