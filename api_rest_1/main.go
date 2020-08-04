package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func main() {
	router := mux.NewRouter()

	//Llenando el arreglo
	people = append(people, Person{
		ID:        "1",
		FirstName: "Elmer",
		LastName:  "Alay",
		Address: &Address{
			City:  "Guatemala",
			State: "Guatemala",
		},
	})

	people = append(people, Person{
		ID:        "2",
		FirstName: "Dua",
		LastName:  "Lipa",
	})

	//endpoints
	router.HandleFunc("/people", getPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", getPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", createPersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", deletePersonEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}

func getPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func getPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Person{})
}

func createPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person

	_ = json.NewDecoder(req.Body).Decode(&person)

	person.ID = params["id"]

	people = append(people, person)

	json.NewEncoder(w).Encode(people)
}

func deletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(people)
}
